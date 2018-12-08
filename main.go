package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//ttpl.Use(r, "template/*", template.FuncMap{"dot": dot})
	r.LoadHTMLGlob("template/*")
	r.Static("files", "./files")
	//r.Handle("css", http.SetPrefix("/css", http.FileServer(http.Dir("style.css"))))
	r.GET("/index", Page)

	r.Run(":8080")
}
