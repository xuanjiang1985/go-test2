package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// os.Getenv(EnvGinMode)
	// gin.SetMode(gin.ReleaseMode)
	// gin.SetMode(gin.DebugMode)

	router := gin.Default()
	router.LoadHTMLFiles("html/index.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", "心跳检测")
	})

	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/gin/report/:uuid", func(c *gin.Context) {
		name := c.Param("uuid")
		c.String(http.StatusOK, name)
	})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	// For each matched request Context will hold the route definition
	router.POST("/user/:name/*action", func(c *gin.Context) {
		//c.FullPath() == "/user/:name/*action" // true
		c.String(http.StatusOK, c.FullPath())
	})

	router.Run(":8080")
}
