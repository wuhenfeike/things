package server

import (
	"context"
	"gitee.com/i-Things/share/conf"
	"gitee.com/i-Things/share/errors"
	"github.com/i-Things/things/service/dmsvr/internal/domain/deviceMsg"
)

type (
	Server interface {
		Subscribe(handle Handle) error
	}
	Handle       func(ctx context.Context) ServerHandle
	ServerHandle interface {
		ActionCheck(req *deviceMsg.PublishMsg) error
	}
)

func NewServer(c conf.EventConf) (Server, error) {
	switch c.Mode {
	case conf.EventModeNats, conf.EventModeNatsJs:
		return newNatsClient(c)
	}
	return nil, errors.Parameter.AddMsgf("mode:%v not support", c.Mode)
}
