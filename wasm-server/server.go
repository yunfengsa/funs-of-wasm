package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func init() {
}

func main() {

	router := gin.Default()
	router.Static("/dist", "dist")
	router.LoadHTMLFiles("dist/index.html")

	router.GET("/wasm", func(c *gin.Context) {
		Cors(c)
		c.Header("Content-Type", "application/wasm")
		c.File("./main.wasm")
	})
	router.Run(":3000")
}

// Cors 设置跨域
func Cors(c *gin.Context) {
	// c.Header("Access-Control-Allow-Origin", "http://localhost:8000")
	// c.Header("Vary", "Origin")
	// c.Header("Access-Control-Allow-Credentials", "true")
	method := c.Request.Method
	origin := c.Request.Header.Get("Origin")
	var headerKeys []string
	for k := range c.Request.Header {
		headerKeys = append(headerKeys, k)
	}
	headerStr := strings.Join(headerKeys, ", ")
	if headerStr != "" {
		headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
	} else {
		headerStr = "access-control-allow-origin, access-control-allow-headers"
	}
	if origin != "" {
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Vary", "Origin")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", headerStr)
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		// c.Set("content-type", "application/json")
	}

	//放行所有OPTIONS方法
	if method == "OPTIONS" {
		c.JSON(http.StatusOK, "Options Request!")
	}
	c.Next()
}
