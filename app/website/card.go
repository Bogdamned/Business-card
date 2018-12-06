package main

import "github.com/gin-gonic/gin"

type (
	obj map[string]interface{}
	C   struct {
		ID       uint64 `form:"id"  json:"id"  bson:"_id"`
		Root     bool   `form:"-"  json:"root"  bson:"root"`
		Name     string `form:"name" json:"name" bson:"name" binding:"required,min=3"`
		Notice   string `form:"notice" json:"notice" bson:"notice"`
		Login    string `form:"login" json:"login" bson:"login" binding:"required"`
		Password string `form:"password" json:"-" bson:"password"`
		Updated  uint   `form:"-" json:"-" bson:"updated"`
		Deleted  bool   `form:"-" json:"-" bson:"deleted"`
	}
)

func Page(c *gin.Context) {
	c.HTML(200, "admins.html", gin.H{
		"title":     "Администраторы",
		"panelPath": "panelPath",
		//"admin":     c.MustGet("user").(A),
		"active": obj{"admins": true},
	})
}
