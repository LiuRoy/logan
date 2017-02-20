package config

import (
	"logan/tools"
)


var (
	MysqlUrl string
	RedisCluster string
	SentryDsn string
)


func init() {
	MysqlUrl = tools.GetSingle("mysql_url")
	RedisCluster = tools.GetSingle("redis_cluster")
	SentryDsn = tools.GetSingle("sentry_dsn")
}
