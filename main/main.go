package main
import (
	"github.com/gin-gonic/gin"
	. "github.com/gonic/route"
)


func main() {
	Router = gin.Default()

	Router.LoadHTMLGlob("templates/*")

	InitializeRoutes()
	//router.GET("/", showIndexPage)
	Router.Run()
}
