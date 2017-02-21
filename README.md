# logan

logan是基于gin的web项目，集成了sentry、gorm、premetheus等工具，旨在高效快速的web开发。

## 编译部署

1. clone本项目至GOPATH下，执行 ```git clone https://github.com/LiuRoy/logan.git $GOPATH/src```

2. 安装依赖库，此项目依赖包列写在requirements文件中，执行命令```go get -u <依赖包>```进行安装

3. 执行```go run main.go```，启动server

## gin

gin是基于go语言轻量级web开发框架（项目地址： [https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)），不仅运行效率很高，而且提供了参数校验、模板渲染、定制化插件等十分有用的功能，极大简化了接口开发难度。为了提供一款快速上手的后台开发模板，logan就用到了gin。

### 添加接口

定义好函数类型为```func(*gin.Context)```的函数，然后添加到对应的路径下即可。

```go
router := gin.New()
router.Use(gin.Logger(), tools.Recovery(sentryClient), tools.Prometheus())

// add endpoint
router.GET(tools.DefaultMetricPath, tools.LatestMetrics)
router.GET("/ping", func(c *gin.Context) {c.String(http.StatusOK, "pong")})
router.POST("/message", apis.AddMessage)
router.GET("/message", apis.GetMessage)
router.Run(address)
```
### 参数获取

+ 针对参数放在url中的请求，可以使用```gin.Context.GetQuery```或者```gin.Context.DefaultQuery```，获取需要的参数。

```go
func GetMessage(c *gin.Context) {
	messageId, exist := c.GetQuery("message_id")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"result": "bad params"})
	} else {
		msgId, _ := strconv.ParseUint(messageId, 10, 0)
		message := model.GetMessage(uint(msgId))
		c.JSON(http.StatusOK, gin.H{"result": "ok", "message": *message})
	}
}
```

+ 针对参数放在body的请求，通过```gin.Context.Bind```解析验证参数(可以制定哪些参数是否必需)。

```go
type message struct {
	Type string `json:"type" binding:"required"`
	InitiatorId uint `json:"initiator" binding:"required"`
	ConsumerId uint `json:"consumer" binding:"required"`
	ResourceId string `json:"resource_id"`
	IsFollow bool `json:"isfollow"`
	Gcid string `json:"gcid"`
	Cid uint `json:"cid"`
	Response string `json:"response"`
	reply string `json:"reply"`
}

func AddMessage(c *gin.Context) {
	var param message
	if c.BindJSON(&param) == nil {
		switch param.Type {
		case "follow":
			model.AddMessage(param.Type, param.InitiatorId,  "aaa", "bbb",  param.ConsumerId, "", "", "", "", "")
		default:
			panic("unknown type")
		}
		c.JSON(http.StatusOK, gin.H{"result": "ok"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"result": "bad params"})
	}
}
```

### 模板渲染

后台常用的返回格式是json，```gin.Context.JSON```很方便的可以把结果序列化为json返回给调用放。

```go
c.JSON(http.StatusOK, gin.H{"result": "ok", "message": *message})
```

## 配置获取

logan提供了两种从consul获取的方式：启动时获取和运行监听。

### 启动时获取

针对大部分不改变的配置，logan只会在启动进程的时候访问consul读取配置。

```go
var SentryDsn string = tools.GetSingle("sentry_dsn")
```

### 运行监听

针对某些经常改变的配置参数，希望consul每次改动在不重启进程的情况下立即生效的场景。每次使用的时候访问consul肯定会增加接口耗时。logan使用了watch方式，每一个参数都有一个对应的goroutine阻塞地监听，一旦改动，马上更新(建议不要设置太多)。

```go
var TestParam tools.WatchedParam
tools.WatchSingle("test_param", &TestParam)
var currentValue string = TestParam.Get()
```

## 监控

监控了各个接口的调用次数、出错次数、接口耗时。如果需要定制一些特定的监控，可以参考prometheus.go文件在业务代码中打点。监控数据通过/logan/metrics获取。

## sentry

在sentry.go中，如果发现接口运行报错，会将错误栈信息发送给sentry，并更新相关接口的出错次数。

## 缓存

使用redis作为缓存，只需要将redis集群消息以```["10.33.111.4:6379", "10.33.22.4:6380", "10.33.44.4:6381", "10.33.1.55:6382", "10.33.1.54:6383", "10.33.1.74:6384"]```格式配置在consul中即可

## 数据库

为了简化业务代码，logan使用了gorm（文档链接： [http://jinzhu.me/gorm/](http://jinzhu.me/gorm/)）。

```go
// 对应的表结构见https://github.com/LiuRoy/logan/blob/master/model/mysql.sql
type Message struct {
	MsgId uint `gorm:"column:msgid;primary_key"`
	Type string `gorm:"column:type;type:varchar(32)"`
	InitiatorId uint `gorm:"column:initiatorid"`
	InitiatorName string `gorm:"column:initiatorname;type:varchar(255)"`
	InitiatorPortrait string `gorm:"column:initiatorportrait;type:varchar(255)"`
	ConsumerId uint `gorm:"column:consumerid"`
	ResourceId string `gorm:"column:resource_id;type:varchar(255)"`
	ExtraInfo1 string `gorm:"column:extra_info1;type:varchar(512)"`
	ExtraInfo2 string `gorm:"column:extra_info2;type:varchar(512)"`
	ExtraInfo3 string `gorm:"column:extra_info3;type:varchar(512)"`
	ExtraInfo4 string `gorm:"column:extra_info4;type:varchar(512)"`
	InsertTime time.Time `gorm:"column:insert_time" sql:"DEFAULT:current_timestamp"`
}

func (Message) TableName() string {
	return "msgcenter_innodb"
}

func AddMessage(msgType string, initiatorId uint, initiatorName string,
	initiatorPortrait string, consumerId uint, resourceId string,
	extraInfo1, extraInfo2, extraInfo3, extraInfo4 string) *Message {
	message := Message{
		Type: msgType,
		InitiatorId: initiatorId,
		InitiatorName: initiatorName,
		InitiatorPortrait: initiatorPortrait,
		ConsumerId: consumerId,
		ResourceId: resourceId,
		ExtraInfo1: extraInfo1,
		ExtraInfo2: extraInfo2,
		ExtraInfo3: extraInfo3,
		ExtraInfo4: extraInfo4,
	}
	DbConnection.Create(&message)
	DbConnection.NewRecord(message)

	RedisConnection.Set("aaaaaaa", message.MsgId, 0)
	return &message
}

func GetMessage(messageId uint) *Message {
	message := Message{}
	DbConnection.Where("msgid = ?", messageId).First(&message)
	return &message
}
```
