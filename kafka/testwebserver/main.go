package testwebserver

import "github.com/gin-gonic/gin"

func InitServer()  {
	r := gin.Default()
	hdl := &TestHandler{}
	hdl.RegisterRouter(r)
	r.Run(":8080")
}
