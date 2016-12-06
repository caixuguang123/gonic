package handlers

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"math/rand"
	"net/http"
	. "github.com/gonic/render"
	. "github.com/gonic/models"
)

func geneSessionToken() string {
	return strconv.FormatInt(rand.Int63(), 16)
}

func ShowRegistrationPage(c *gin.Context) {
	Render(c, gin.H{
		"title":"Register"}, "register.html")
}

func Register(c *gin.Context) {
	userName := c.PostForm("userName")
	password := c.PostForm("password")

	if _, err := RegisterUser(userName, password); err == nil {
		token := geneSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)
		Render(c, gin.H{"title":"successful login"}, "loginSuccess.html")
	} else {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"errorTitle":"reigster failed",
			"errorMessage":err.Error()})
	}
}

func ShowLoginPage(c *gin.Context) {
	Render(c, gin.H{"title":"login"}, "login.html")

}

func PerformLogin(c *gin.Context) {
	userName := c.PostForm("account")
	password := c.PostForm("psw")
	if valid := IsUserValid(userName, password); valid {
		token := geneSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)
		Render(c, gin.H{"title":"successful login"}, "loginSuccess.html")
	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"errTitle":"login failed",
			"errMessage":"user Name or Password Wrong"})
	}
}

func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "", "", false, true)
	c.Redirect(http.StatusTemporaryRedirect, "/")
}
