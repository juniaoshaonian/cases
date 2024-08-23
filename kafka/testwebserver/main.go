package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	hdl := &TestHandler{}
	hdl.RegisterRouter(r)
	r.Run(":8080")
}
