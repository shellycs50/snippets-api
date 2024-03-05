package controller

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"shellycs50.com/snippetsapi/db"
)

func ErrorResponse(err error, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"error": fmt.Sprintf("%s", err),
	})
}

func PostSnippet(c *gin.Context) {
	var snippet db.Snippet
	err := c.BindJSON(&snippet)
	//validate
	if snippet.Content == "" || snippet.Name == "" || snippet.Language == "" {
		ErrorResponse(errors.New("missing required fields"), c)
		return
	}
	snippet.UploadDateTime = time.Now()

	if err != nil {
		ErrorResponse(err, c)
	}

	saveErr := db.SaveSnippet(snippet)
	if saveErr != nil {
		ErrorResponse(errors.New("could not save (server error)"), c)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Snippet saved successfully",
	})
}
func GetAllSnippets(c *gin.Context) {
	snippets := db.GetAllSnippets()
	c.JSON(http.StatusOK, snippets)
}
func GetAllLangs(c *gin.Context) {
	langs := db.GetAllLangs()
	c.JSON(http.StatusOK, langs)
}
