package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var router *gin.Engine

func main() {
	router = gin.Default()

	//router.LoadHTMLGlob("templateReactJs/*")
	//router.Handle("/", http.FileServer(http.Dir("./templateReactJs")))
	router.Static("/static","public");
	initializeRoutes()
	//router.GET("/", showIndexPage)
	router.Run()
}

func render(c *gin.Context, data gin.H, templateName string) {
	//loggedInInterface, _ := c.Get("is_logged_in")
	//data["is_logged_in"] = loggedInInterface.(bool)
	data["is_logged_in"] = false
	//fmt.Print("cccc"+c.Request.Header.Get("Accept"))

	switch c.Request.Header.Get("Accept") {
	case "application/json, text/javascript, */*; q=0.01":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, templateName, data)
	}
}
