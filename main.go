package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"shellycs50.com/snippetsapi/controller"
)

func main() {

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
	}))

	r.GET("/snippets/:id", controller.GetSnippet)

	r.GET("/snippets", controller.GetAllSnippets)

	r.GET("/languages", controller.GetAllLangs)

	r.POST("/snippets", controller.PostSnippet)

	r.PUT("/snippets/:id", controller.EditSnippet)

	r.DELETE(("snippets/:id"), controller.DeleteSnippet)

	r.Run() // 8080 by default
}
