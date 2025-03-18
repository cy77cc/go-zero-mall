package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql      MySQL
	CacheRedis cache.CacheConf
	Salt       string
}

type MySQL struct {
	DataSource string
}
