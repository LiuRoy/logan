package tools

import (
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/getsentry/raven-go"
	"github.com/gin-gonic/gin"
)

func Recovery(client *raven.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			flags := map[string]string{
				"endpoint": c.Request.URL.String(),
			}
			if rval := recover(); rval != nil {
				debug.PrintStack()
				rvalStr := fmt.Sprint(rval)
				packet := raven.NewPacket(rvalStr, raven.NewException(errors.New(rvalStr),
					raven.NewStacktrace(2, 3, nil)))
				client.Capture(packet, flags)
				c.Writer.WriteHeader(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
