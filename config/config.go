package main

import (
	"fmt"
	"sync"
	"net/http"


	"github.com/gin-gonic/gin"
	consulApi "github.com/hashicorp/consul/api"
	consulWatch "github.com/hashicorp/consul/watch"
)

const (
	address string = "10.33.1.132:8500"
	prefix string = "preview/service/kingkong"
)

var (
	kv *consulApi.KV
	MysqlUrl string
	RedisCluster string
	TestConfig WatchedParam
)


// 监听consul配置
type WatchedParam struct {
	value string
	lock sync.RWMutex
}

func (v *WatchedParam) Get() string {
	v.lock.RLock()
	defer v.lock.RUnlock()
	return v.value
}

func (v *WatchedParam) Set(value string) {
	v.lock.Lock()
	defer v.lock.Unlock()
	v.value = value
}


func getSingle(key string) string {
	pair, _, err := kv.Get(fmt.Sprintf("%s/%s", prefix, key), nil)
	if err != nil {
		panic(err)
	}
	return string(pair.Value)
}

func watchSingle(key string, param *WatchedParam) {

	params := map[string]interface{}{
		"type": "key",
		"key": fmt.Sprintf("%s/%s", prefix, key),
	}
	plan, _ := consulWatch.Parse(params)
	plan.Handler = func(idx uint64, raw interface{}) {
		if raw == nil {
			return
		}

		v, ok := raw.(*consulApi.KVPair)
		if ok && v != nil {
			newValue := string(v.Value)
			fmt.Println(newValue)
			param.Set(newValue)
		}
	}

	go plan.Run(address)
}

func init() {
	consulConfig := consulApi.Config{Address: address}
	client, err := consulApi.NewClient(&consulConfig)
	if err != nil {
		panic(err)
	}

	kv = client.KV()
	MysqlUrl = getSingle("mysql_config")
	RedisCluster = getSingle("redis_cluster")

	watchSingle("test_config", &TestConfig)
}

func main() {
	fmt.Println(MysqlUrl)
	fmt.Println(RedisCluster)

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "%s\n", TestConfig.Get())
	})
	router.Run("127.0.0.1:8999")
}
