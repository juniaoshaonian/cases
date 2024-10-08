package testwebserver

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"
)

// 模拟下游业务
type TestHandler struct {
	count int64
}

func (t *TestHandler) RegisterRouter(server *gin.Engine) {
	server.GET("/test", func(c *gin.Context) {
		time.Sleep(100 * time.Millisecond)
		atomic.AddInt64(&t.count, 1)
	})

	server.GET("/batch-test", func(c *gin.Context) {
		time.Sleep(250 * time.Millisecond)
		atomic.AddInt64(&t.count, 5)
	})
	server.GET("/count", func(c *gin.Context) {
		c.JSON(http.StatusOK, t.count)
	})

}

func BatchTest() {
	getResp("http://localhost:8000/batch-test")

}

func Test() {
	getResp("http://localhost:8000/test")
}

func GetCount() int64 {
	resp, err := http.Get("http://localhost:8080/number")
	if err != nil {

		return 0
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0
	}
	count,_ := strconv.Atoi(string(body))
	return int64(count)
}

func getResp(url string) {
	req, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(nil))
	if err != nil {
		return
	}
	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	resp.Body.Close()
}
