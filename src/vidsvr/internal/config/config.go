package config

import (
	"github.com/i-Things/things/shared/conf"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Database    conf.Database
	CacheRedis  cache.ClusterConf
	Event       conf.EventConf     //和things内部交互的设置
	TimedJobRpc conf.RpcClientConf `json:",optional"`
}