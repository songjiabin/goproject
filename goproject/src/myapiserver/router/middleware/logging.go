package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/willf/pad"
	"io/ioutil"
	"myapiserver/handler"
	"myapiserver/pkg/errno"
	"regexp"
	"time"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

//处理每个请求
func Logging(c *gin.Context) {

	logs.Info("-----------------------------------------------------------------------------------")

	start := time.Now().UTC()
	path := c.Request.URL.Path

	reg := regexp.MustCompile("(/v1/user|/login)")
	if !reg.MatchString(path) {
		return
	}

	// Skip for the health check requests.
	if path == "/sd/health" || path == "/sd/ram" || path == "/sd/cpu" || path == "/sd/disk" {
		return
	}

	// Read the Body content
	var bodyBytes []byte
	if c.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}

	// Restore the io.ReadCloser to its original state
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	// The basic informations.
	method := c.Request.Method
	ip := c.ClientIP()

	//log.Debugf("New request come in, path: %s, Method: %s, body `%s`", path, method, string(bodyBytes))
	blw := &bodyLogWriter{
		body:           bytes.NewBufferString(""),
		ResponseWriter: c.Writer,
	}
	c.Writer = blw

	// Continue.
	c.Next()

	// Calculates the latency.
	end := time.Now().UTC()
	latency := end.Sub(start)

	code, message := -1, ""

	// get code and message
	var response handler.Response
	if err := json.Unmarshal(blw.body.Bytes(), &response); err != nil {
		log.Errorf(err, "response body can not unmarshal to model.Response struct, body: `%s`", blw.body.Bytes())
		code = errno.InternalServerError.Code
		message = err.Error()
	} else {
		code = response.Code
		message = response.Message
	}

	log.Infof("%-13s | %-12s | %s %s | {code: %d, message: %s}", latency, ip, pad.Right(method, 5, ""), path, code, message)
}
