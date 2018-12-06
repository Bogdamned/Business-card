package main

import (
	"github.com/gin-gonic/gin"
	//ttpl "gopkg.in/night-codes/go-gin-ttpl.v1"
)

func main() {
	r := gin.Default()
	//ttpl.Use(r, "template/*", template.FuncMap{"dot": dot})
	r.LoadHTMLGlob("template/*")

	r.GET("/index", Page)

	r.Run(":8080")
}
