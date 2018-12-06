package main

import (
	"html/template"

	"github.com/gin-gonic/gin"
	ttpl "gopkg.in/night-codes/go-gin-ttpl.v1"
)

func main() {
	r := gin.Default()
	ttpl.Use(r, "template/*", template.FuncMap{"dot": dot})

	r.Run(":8080")
}
