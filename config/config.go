package config

import (
	"logan/toolkits/consul"
)


var (
	MysqlUrl string
	RedisCluster string
	SentryDsn string
	TestConfig consul.WatchedParam
)


func init() {
	MysqlUrl = consul.GetSingle("mysql_config")
	RedisCluster = consul.GetSingle("redis_cluster")
	SentryDsn = consul.GetSingle("sentry_dsn")

	consul.WatchSingle("test_config", &TestConfig)
}
