package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"net/http"
	"time"

	"github.com/ecodeclub/ekit/iox"
)

// 模拟下游业务
type TestHandler struct {
}

func (t *TestHandler) RegisterRouter(server *gin.Engine) {
	server.GET("/test", func(c *gin.Context) {
		time.Sleep(100 * time.Millisecond)
	})

	server.GET("/batch-test", func(c *gin.Context) {
		time.Sleep(150 * time.Millisecond)
	})
}

func BatchTest() {
	req, err := http.NewRequest(http.MethodPost,
		"/test", iox.NewJSONReader())
	req.Header.Set("content-type", "application/json")
	require.NoError(t, err)
}
