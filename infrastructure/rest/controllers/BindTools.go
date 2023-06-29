package controllers

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
)

// BindJSON is a function that binds the request body to the given struct and rewrite the request body on the context
func BindJSON(c *gin.Context, request interface{}) (err error) {
	buf := make([]byte, 5120)
	num, _ := c.Request.Body.Read(buf)
	reqBody := string(buf[0:num])
	c.Request.Body = io.NopCloser(bytes.NewBuffer([]byte(reqBody)))
	err = c.ShouldBindJSON(request)
	c.Request.Body = io.NopCloser(bytes.NewBuffer([]byte(reqBody)))
	return
}
