package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
)

type (
	obj map[string]interface{}
	A   struct {
		ID       uint64 `form:"id"  json:"id"  bson:"_id"`
		Root     bool   `form:"-"  json:"root"  bson:"root"`
		Name     string `form:"name" json:"name" bson:"name" binding:"required,min=3"`
		Notice   string `form:"notice" json:"notice" bson:"notice"`
		Login    string `form:"login" json:"login" bson:"login" binding:"required"`
		Password string `form:"password" json:"-" bson:"password"`
		Updated  uint   `form:"-" json:"-" bson:"updated"`
		Deleted  bool   `form:"-" json:"-" bson:"deleted"`
	}

	Application struct {
		collection *mgo.Collection
	}
)

func Page(c *gin.Context) {
	c.HTML(http.StatusOK, "card.html", gin.H{
		"title": "Сервис ремонта бытовой техники",
	})
}
