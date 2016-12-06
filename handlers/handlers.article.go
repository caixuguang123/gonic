package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	. "github.com/gonic/models"
	. "github.com/gonic/render"
)

func ShowIndexPage(c *gin.Context) {
	articles := GetAllArticles()
	Render(c, gin.H{
		"title":"home page",
		"payload":articles,
	}, "index.html")
}

func GetArticle(c *gin.Context) {
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		if article, err := GetArticleById(articleID); err == nil {
			c.HTML(
				http.StatusOK,
				"article.html",
				gin.H{
					"title": article.Title,
					"payload": article,
				},
			)
		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
	//test line
}

func ShowArticleCreatePage(c *gin.Context) {
	Render(c, gin.H{"title":"createArticle"}, "createArticle.html")
}

func CreateArticle(c *gin.Context) {
	title := c.PostForm("articleTitle")
	content := c.PostForm("articleContent")
	if a, err := CreateNewArticle(title, content); err == nil {
		Render(c, gin.H{
			"title":"create successfully",
			"payload":a}, "submissionSuccess.html")
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}