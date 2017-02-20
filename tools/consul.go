package tools

import (
	"fmt"
	"sync"

	consulApi "github.com/hashicorp/consul/api"
	consulWatch "github.com/hashicorp/consul/watch"
)

const (
	address string = "10.33.1.132:8500"
	prefix string = "preview/service/kingkong"
)

var KV *consulApi.KV


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


func GetSingle(key string) string {
	pair, _, err := KV.Get(fmt.Sprintf("%s/%s", prefix, key), nil)
	if err != nil {
		panic(err)
	}
	return string(pair.Value)
}

func WatchSingle(key string, param *WatchedParam) {

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
	conConfig := consulApi.Config{Address: address}
	client, err := consulApi.NewClient(&conConfig)
	if err != nil {
		panic(err)
	}

	KV = client.KV()
}
