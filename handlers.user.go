package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"math/rand"
	"net/http"
)

func geneSessionToken() string {
	return strconv.FormatInt(rand.Int63(), 16)
}

func showRegistrationPage(c *gin.Context) {
	render(c, gin.H{
		"title":"Register"}, "register.html")
}

func register(c *gin.Context) {
	userName := c.PostForm("userName")
	password := c.PostForm("password")

	if _, err := registerUser(userName, password); err == nil {
		token := geneSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)
		render(c, gin.H{"title":"successful login"}, "loginSuccess.html")
	} else {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"errorTitle":"reigster failed",
			"errorMessage":err.Error()})
	}
}

func showLoginPage(c *gin.Context) {
	render(c, gin.H{"title":"login"}, "login.html")

}

func performLogin(c *gin.Context) {
	userName := c.PostForm("account")
	password := c.PostForm("psw")
	if valid := isUserValid(userName, password); valid {
		token := geneSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)
		render(c, gin.H{"title":"successful login"}, "loginSuccess.html")
	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"errTitle":"login failed",
			"errMessage":"user Name or Password Wrong"})
	}
}

func logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "", "", false, true)
	c.Redirect(http.StatusTemporaryRedirect, "/")
}