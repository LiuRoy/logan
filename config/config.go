package config

import (
	"logan/tools"
)


var (
	MysqlUrl string
	RedisCluster string
	SentryDsn string
	TestConfig tools.WatchedParam
)


func init() {
	MysqlUrl = tools.GetSingle("mysql_config")
	RedisCluster = tools.GetSingle("redis_cluster")
	SentryDsn = tools.GetSingle("sentry_dsn")

	tools.WatchSingle("test_config", &TestConfig)
}
