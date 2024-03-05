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
	snippet.Deleted = 0

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

func DeleteSnippet(c *gin.Context) {
	id := c.Param("id")
	err := db.DeleteSnippet(id)
	if err != nil {
		ErrorResponse(err, c)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Snippet deleted successfully",
	})
}

func GetSnippet(c *gin.Context) {
	id := c.Param("id")
	snippet, err := db.GetSnippet(id)
	if err != nil {
		ErrorResponse(err, c)
	}
	c.JSON(http.StatusOK, snippet)
}

func EditSnippet(c *gin.Context) {
	id := c.Param("id")
	var snippet db.Snippet
	err := c.BindJSON(&snippet)
	if err != nil {
		ErrorResponse(err, c)
	}
	err = db.EditSnippet(id, snippet)
	if err != nil {
		ErrorResponse(err, c)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Snippet edited successfully",
	})
}
