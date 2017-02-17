package main

import (
	"flag"
	"fmt"
	"runtime"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	workerNum int
	releaseMode bool
	host string
	port int
)

func init() {
	cpuNum := runtime.NumCPU()
	flag.IntVar(&workerNum, "worker", cpuNum, "runtime MAXPROCS value")
	flag.BoolVar(&releaseMode, "release", false, "gin mode")
	flag.StringVar(&host, "host", "127.0.0.1", "server host")
	flag.IntVar(&port, "port", 8090, "server port")
}

func main() {
	flag.Parse()
	address := fmt.Sprintf("%s:%d", host, port)
	fmt.Printf("%s serving... MAXPROCS:%d release:%t\n", address, workerNum, releaseMode)

	runtime.GOMAXPROCS(workerNum)

	if releaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	// add endpoint
	router.GET("/ping", func(c *gin.Context) {c.String(http.StatusOK, "pong")})

	router.Run(address)
}
