package common

import (
	"gitee.com/i-Things/share/def"
	"gitee.com/i-Things/share/utils"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/i-Things/things/service/vidsvr/internal/repo/relationDB"
	"github.com/i-Things/things/service/vidsvr/pb/vid"
	"time"
)

/*
根据用户的输入生成对应的数据库数据
*/
func ToVidmgrInfoDB(in *vid.VidmgrInfo) (*relationDB.VidmgrInfo, error) {
	pi := &relationDB.VidmgrInfo{
		VidmgrID:     in.VidmgrID,
		VidmgrName:   in.VidmgrName,
		VidmgrIpV4:   utils.InetAtoN(in.VidmgrIpV4),
		VidmgrPort:   in.VidmgrPort,
		VidmgrStatus: in.VidmgrStatus,
		VidmgrSecret: in.VidmgrSecret,
		VidmgrType:   in.VidmgrType,
		IsOpenGbSip:  true,  //默认开启
		RtpPort:      10000, //默认RTP端口为10000
		MediasvrType: 2,     //设置为独立主机
		Desc:         in.Desc.GetValue(),
	}
	if in.Tags == nil {
		in.Tags = map[string]string{}
	}
	pi.Tags = in.Tags
	return pi, nil
}

func ToVidmgrInfoRPC(pi *relationDB.VidmgrInfo) *vid.VidmgrInfo {

	if pi.VidmgrType == def.Unknown {
		pi.VidmgrType = def.VidmgrTypeZLMedia //当前默认仅支持zlmediakit
	}
	dpi := &vid.VidmgrInfo{
		VidmgrID:     pi.VidmgrID,   //服务id
		VidmgrName:   pi.VidmgrName, //服务名
		VidmgrIpV4:   utils.InetNtoA(pi.VidmgrIpV4),
		VidmgrPort:   pi.VidmgrPort,
		VidmgrType:   pi.VidmgrType,   //流服务器类型:1:zlmediakit,2:srs,3:monibuca
		VidmgrStatus: pi.VidmgrStatus, //服务状态: 1：离线 2：在线  0：未激活
		VidmgrSecret: pi.VidmgrSecret, //流服务器注秘钥 只读
		RtpPort:      pi.RtpPort,
		IsOpenSip:    pi.IsOpenGbSip,
		Desc:         &wrappers.StringValue{Value: pi.Desc}, //描述
		CreatedTime:  pi.CreatedTime.Unix(),                 //创建时间
		LastLogin:    pi.LastLogin.Unix(),                   //最后登录时间
		FirstLogin:   pi.FirstLogin.Unix(),                  //首次登录时间
		Tags:         pi.Tags,                               //产品tags
	}

	return dpi
}

func UpdatVidmgrInfoDB(old *relationDB.VidmgrInfo, data *vid.VidmgrInfo) error {
	if data.VidmgrName != "" {
		old.VidmgrName = data.VidmgrName
	}
	if data.VidmgrIpV4 != "" {
		old.VidmgrIpV4 = utils.InetAtoN(data.VidmgrIpV4)
	}
	if data.VidmgrPort != 0 {
		old.VidmgrPort = data.VidmgrPort
	}
	if data.VidmgrType != 0 {
		old.VidmgrType = data.VidmgrType
	}
	if data.VidmgrSecret != "" {
		old.VidmgrSecret = data.VidmgrSecret
	}
	if data.FirstLogin != 0 {
		old.FirstLogin = time.Unix(data.FirstLogin, 0)
	}
	if data.LastLogin != 0 {
		old.LastLogin = time.Unix(data.LastLogin, 0)
	}
	if data.MediasvrType != 0 {
		old.MediasvrType = data.MediasvrType
	}
	if data.VidmgrStatus != 0 {
		old.VidmgrStatus = data.VidmgrStatus
	}
	if data.Desc != nil {
		old.Desc = data.Desc.GetValue()
	}
	if data.Tags != nil {
		old.Tags = data.Tags
	}
	return nil
}
