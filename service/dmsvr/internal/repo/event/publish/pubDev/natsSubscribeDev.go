package pubDev

import (
	"gitee.com/unitedrhino/share/ctxs"
	"gitee.com/unitedrhino/share/errors"
	"gitee.com/unitedrhino/share/events"
	"gitee.com/unitedrhino/share/utils"
	"gitee.com/unitedrhino/things/share/domain/deviceMsg"
	"github.com/nats-io/nats.go"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type (
	natsSubDev struct {
		subscription *nats.Subscription
	}
)

func newNatsSubDev(subscription *nats.Subscription) *natsSubDev {
	return &natsSubDev{subscription: subscription}
}

func (s *natsSubDev) UnSubscribe() error {
	return s.subscription.Unsubscribe()
}

func (s *natsSubDev) GetMsg(timeout time.Duration) (ele *deviceMsg.PublishMsg, err error) {
	msg, err := s.subscription.NextMsg(timeout)
	if err != nil {
		logx.Errorf("%s.NextMsg failure err:%v", utils.FuncName(), err)
		if err == nats.ErrTimeout {
			return nil, errors.TimeOut.AddMsg("设备回复超时")
		}
		return nil, err
	}
	msg.Ack()
	emsg := events.GetEventMsg(msg.Data)
	if emsg == nil {
		logx.Error(msg.Subject, string(msg.Data))
		return
	}
	ctx := emsg.GetCtx()
	//向jaeger推送当前节点信息，路径名为主题名
	ctx, span := ctxs.StartSpan(ctx, msg.Subject, "")
	logx.Debugf("%s trace:%s  spanID:%s topic:%s", utils.FuncName(),
		span.SpanContext().TraceID(), span.SpanContext().SpanID(), msg.Subject)
	defer span.End()
	ele, err = deviceMsg.GetDevPublish(ctx, emsg.GetData())
	if err != nil {
		logx.WithContext(ctx).Error(msg.Subject, string(msg.Data), err)
		return
	}
	return
}
