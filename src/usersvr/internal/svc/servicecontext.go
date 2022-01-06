package svc

import (
	"github.com/go-things/things/shared/third/weixin"
	"github.com/go-things/things/shared/utils"
	"github.com/tal-tech/go-zero/core/stores/sqlx"

	"github.com/go-things/things/src/usersvr/internal/config"
	//"gitee.com/godLei6/things/shared/third/weixin"
	"github.com/go-things/things/src/usersvr/model"
)

type ServiceContext struct {
	Config        config.Config
	UserInfoModel model.UserInfoModel
	UserCoreModel model.UserCoreModel
	UserModel     model.UserModel
	WxMiniProgram *weixin.MiniProgram
	UserID        *utils.SnowFlake
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	ui := model.NewUserInfoModel(conn, c.CacheRedis)
	uc := model.NewUserCoreModel(conn, c.CacheRedis)
	um := model.NewUserModel(conn, c.CacheRedis)
	WxMiniProgram := weixin.NewWexinMiniProgram(c.WexinMiniprogram, c.CacheRedis)
	UserID := utils.NewSnowFlake(c.NodeID)

	return &ServiceContext{
		Config:        c,
		UserInfoModel: ui,
		UserCoreModel: uc,
		UserModel:     um,
		WxMiniProgram: WxMiniProgram,
		UserID:        UserID,
	}
}
