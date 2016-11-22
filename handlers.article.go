package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()
	render(c, gin.H{
		"title":"home page",
		"payload":articles,
	}, "index.html")
}

func RgetAllArticles(c *gin.Context)  {

}

func getArticle(c *gin.Context) {
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		if article, err := getArticleById(articleID); err == nil {
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

func showArticleCreatePage(c *gin.Context) {
	render(c, gin.H{"title":"createArticle"}, "createArticle.html")
}

func createArticle(c *gin.Context) {
	title := c.PostForm("articleTitle")
	content := c.PostForm("articleContent")
	if a, err := createNewArticle(title, content); err == nil {
		render(c, gin.H{
			"title":"create successfully",
			"payload":a}, "submissionSuccess.html")
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}