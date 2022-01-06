package logic

import (
	"context"
	"encoding/json"
	"github.com/go-things/things/shared/def"
	"github.com/go-things/things/shared/errors"
	"github.com/go-things/things/shared/utils"
	"github.com/go-things/things/src/dmsvr/device"
	"github.com/go-things/things/src/dmsvr/dm"
	"github.com/go-things/things/src/dmsvr/internal/exchange/types"
	"github.com/go-things/things/src/dmsvr/internal/repo"
	"github.com/go-things/things/src/dmsvr/internal/repo/model/mysql"
	"github.com/go-things/things/src/dmsvr/internal/svc"
	"github.com/tal-tech/go-zero/core/logx"
	"strings"
	"time"
)

type PublishLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	ld       *dm.LoginDevice
	pi       *mysql.ProductInfo
	template *device.Template
	topics   []string
	dreq     device.DeviceReq
	dd       repo.DeviceDataRepo
}

func NewPublishLogic(ctx context.Context, svcCtx *svc.ServiceContext) LogicHandle {
	return LogicHandle(&PublishLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	})
}

func (l *PublishLogic) initMsg(msg *types.Elements) error {
	var err error
	l.ld, err = dm.GetClientIDInfo(msg.ClientID)
	if err != nil {
		return err
	}
	l.pi, err = l.svcCtx.ProductInfo.FindOneByProductID(l.ld.ProductID)
	if err != nil {
		return err
	}
	l.template, err = device.NewTemplate([]byte(l.pi.Template))
	if err != nil {
		return err
	}
	err = utils.Unmarshal([]byte(msg.Payload), &l.dreq)
	if err != nil {
		return errors.Parameter.AddDetail("things topic is err:" + msg.Topic)
	}
	l.dd = l.svcCtx.DeviceData(l.ctx)
	l.topics = strings.Split(msg.Topic, "/")
	return nil
}

func (l *PublishLogic) DeviceResp(err error, data map[string]interface{}) {
	l.svcCtx.DevClient.DeviceResp(l.dreq.Method, l.dreq.ClientToken, l.topics, err, data)
}

func (l *PublishLogic) HandlePropertyReport(msg *types.Elements) error {
	tp, err := l.template.VerifyReqParam(l.dreq, device.PROPERTY)
	if err != nil {
		l.DeviceResp(err, nil)
		return err
	} else if len(tp) == 0 {
		err := errors.Parameter.AddDetail("need right param")
		l.DeviceResp(err, nil)
		return err
	}
	params := device.ToVal(tp)
	timeStamp := time.Now()
	if l.dreq.Timestamp != 0 {
		timeStamp = time.UnixMilli(l.dreq.Timestamp)
	}
	err = l.dd.InsertPropertiesData(l.ld.ProductID, l.ld.DeviceName, params, timeStamp)
	if err != nil {
		l.DeviceResp(errors.Database, nil)
		l.Errorf("HandlePropertyReport|InsertPropertyData|err=%+v", err)
		return err
	}
	l.DeviceResp(errors.OK, nil)
	return nil
}

func (l *PublishLogic) HandlePropertyGetStatus(msg *types.Elements) error {
	respData := make(map[string]interface{}, len(l.template.Properties))
	switch l.dreq.Type {
	case device.REPORT:
		for id, _ := range l.template.Property {
			data, err := l.dd.GetPropertyDataWithID(l.ld.ProductID, l.ld.DeviceName, id, 0, 0, 1)
			if err != nil {
				l.Errorf("HandlePropertyGetStatus|GetPropertyDataWithID|get id:%s|err:%s",
					id, err.Error())
				return err
			}
			if len(data) == 0 {
				l.Slowf("HandlePropertyGetStatus|GetPropertyDataWithID|not find id:%s", id)
				continue
			}
			respData[id] = data[0].Param
		}
	default:
		err := errors.Parameter.AddDetailf("not suppot type :%s", l.dreq.Type)
		l.DeviceResp(err, nil)
		return err
	}
	l.DeviceResp(errors.OK, respData)
	return nil
}

func (l *PublishLogic) HandleProperty(msg *types.Elements) error {
	l.Slowf("PublishLogic|HandleProperty")
	switch l.dreq.Method {
	case device.REPORT, device.REPORT_INFO:
		return l.HandlePropertyReport(msg)
	case device.GET_STATUS:
		return l.HandlePropertyGetStatus(msg)
	case device.CONTROL_REPLY:
		return l.HandleResp(msg)
	default:
		return errors.Method
	}
	return nil
}

func (l *PublishLogic) HandleEvent(msg *types.Elements) error {
	l.Slowf("PublishLogic|HandleEvent")
	dbData := repo.Event{}
	dbData.ID = l.dreq.EventID
	dbData.Type = l.dreq.Type
	if l.dreq.Method != device.EVENT_POST {
		return errors.Method
	}
	tp, err := l.template.VerifyReqParam(l.dreq, device.EVENT)
	if err != nil {
		l.DeviceResp(err, nil)
		return err
	}
	dbData.Params = device.ToVal(tp)
	if l.dreq.Timestamp != 0 {
		dbData.TimeStamp = time.UnixMilli(l.dreq.Timestamp)
	} else {
		dbData.TimeStamp = time.Now()
	}
	err = l.dd.InsertEventData(l.ld.ProductID, l.ld.DeviceName, &dbData)
	if err != nil {
		l.DeviceResp(errors.Database, nil)
		l.Errorf("InsertEventData|err=%+v", err)
		return errors.Database.AddDetail(err)
	}
	l.DeviceResp(errors.OK, nil)

	return nil
}
func (l *PublishLogic) HandleResp(msg *types.Elements) error {
	l.Slowf("PublishLogic|HandleResp")
	resp := device.DeviceResp{}
	err := json.Unmarshal([]byte(msg.Payload), &resp)
	if err != nil {
		return errors.Parameter.AddDetail(err)
	}
	l.svcCtx.DevClient.DeviceReqSendResp(&resp, msg.Topic)
	return nil
}

func (l *PublishLogic) HandleThing(msg *types.Elements) error {
	l.Slowf("PublishLogic|HandleThing")
	if len(l.topics) < 5 || l.topics[1] != "up" {
		return errors.Parameter.AddDetail("things topic is err:" + msg.Topic)
	}
	switch l.topics[2] {
	case def.PROPERTY_METHOD: //属性上报
		return l.HandleProperty(msg)
	case def.EVENT_METHOD: //事件上报
		return l.HandleEvent(msg)
	case def.ACTION_METHOD: //设备响应行为执行结果
		return l.HandleResp(msg)
	default:
		return errors.Parameter.AddDetail("things topic is err:" + msg.Topic)
	}
	return nil
}
func (l *PublishLogic) HandleOta(msg *types.Elements) error {
	l.Slowf("PublishLogic|HandleOta")
	return nil
}

func (l *PublishLogic) HandleDefault(msg *types.Elements) error {
	l.Slowf("PublishLogic|HandleDefault")
	return nil

}
func (l *PublishLogic) Handle(msg *types.Elements) error {
	l.Infof("PublishLogic|req=%+v", msg)
	err := l.initMsg(msg)
	if err != nil {
		return err
	}
	if len(l.topics) > 1 {
		switch l.topics[0] {
		case "$thing":
			return l.HandleThing(msg)
		case "$ota":
			return l.HandleOta(msg)
		case l.pi.ProductID:
			return l.HandleDefault(msg)
		default:
			return errors.Parameter.AddDetailf("not suppot topic :%s", msg.Topic)
		}
	}
	return nil
}
