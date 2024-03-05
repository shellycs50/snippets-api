package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"shellycs50.com/snippetsapi/controller"
)

func main() {

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
	}))
	r.GET("/snippets/:id", func(c *gin.Context) {
		userId := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Requested Snippet is Snippet number %s", userId),
		})
	})

	r.POST("/snippets", controller.PostSnippet)

	r.GET("/snippets", controller.GetAllSnippets)

	r.GET("/languages", controller.GetAllLangs)

	r.Run() // 8080 by default
}
