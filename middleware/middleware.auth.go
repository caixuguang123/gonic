package middleware
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func EnsureLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("------------ensure logged in------------")
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if !loggedIn {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func EnsureNotLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("------------ensure not logged in------------")
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if loggedIn {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func SetUserStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("------------set user status------------")
		if token, err := c.Cookie("token"); err == nil || token != "" {
			fmt.Println("------------already logged in---------with token " + token)
			c.Set("is_logged_in", true)

		} else {
			fmt.Println("------------not logged in------------")
			c.Set("is_logged_in", false)
		}

	}
}
