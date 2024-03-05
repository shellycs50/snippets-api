package db

import (
	"time"
)

type Snippet struct {
	Name           string    `json:"name"`
	Content        string    `json:"content"`
	Language       string    `json:"language"`
	UploadDateTime time.Time `json:"uploadDateTime"`
}

type DB struct {
	Snippets  []Snippet `json:"snippets"`
	Languages []string  `json:"languages"`
}
