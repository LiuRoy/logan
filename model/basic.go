package model

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/redis.v5"

	"logan/config"
)

var (
	DbConnection *gorm.DB
	RedisConnection *redis.ClusterClient
)

func init() {
	db, err := gorm.Open("mysql", config.MysqlUrl)
	if err != nil {
		panic(err)
	}
	DbConnection = db

	var clusterList []string
	if json.Unmarshal([]byte(config.RedisCluster), &clusterList) != nil{
		panic("redis cluster configure decode error")
	}
	clusterOpt := redis.ClusterOptions{Addrs: clusterList}
	RedisConnection = redis.NewClusterClient(&clusterOpt)
}
