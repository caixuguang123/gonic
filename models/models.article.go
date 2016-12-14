package models

import (
	"errors"
	"gopkg.in/mgo.v2"
	"github.com/gonic/db"
	"gopkg.in/mgo.v2/bson"
)

type Article struct {
	ID      string `json:"id" bson:"id"`
	Title   string `json:"title" bson:"title"`
	Content string `json:"content" bson:"content"`
}

func GetAllArticles() []Article {
	var articleList []Article
	query := func(c *mgo.Collection) error {
		return c.Find(nil).All(&articleList)
	}

	error := db.WitchCollection("article", query)
	if error != nil {
		return nil
	}
	return articleList
}

func GetArticleById(id string) (*Article, error) {
	var article Article
	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{"id": id}).One(&article)
	}

	error := db.WitchCollection("article", query)
	if error != nil {
		return nil, error
	} else {
		return &article, nil
	}
	return nil, errors.New("Article not found")
}

func CreateNewArticle(title, content string) (*Article, error) {
	bid := bson.NewObjectId()
	id := bid.Hex()
	articleNew := Article{id, title, content}
	query := func(c *mgo.Collection) error {
		return c.Insert(articleNew)
	}
	error := db.WitchCollection("article", query)
	if error != nil {
		return nil, error
	} else {
		return &articleNew, nil
	}
}
