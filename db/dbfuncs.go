package db

import (
	"encoding/json"
	"errors"
	"strconv"
)

func GetAllSnippets() []Snippet {
	db, err := getDb()
	if err != nil {
		return []Snippet{}
	}
	return db.Snippets
}

func GetAllLangs() []string {
	db, err := getDb()
	if err != nil {
		return []string{}
	}
	return db.Languages
}

func SaveSnippet(snippet Snippet) error {
	db, err := getDb()
	if err != nil {
		return err
	}
	lengthOfSnippetsList := len(db.Snippets)
	snippet.Id = lengthOfSnippetsList
	db.Snippets = append(db.Snippets, snippet)

	updatedData, err := json.Marshal(db)
	if err != nil {
		return err
	}

	err = dbSave("./db.json", updatedData)
	if err != nil {
		return err
	}
	updateLangs()
	return nil
}

func updateLangs() error {
	db, err := getDb()
	if err != nil {
		return err
	}
	newLangMap := make(map[string]int)
	newLangsList := make([]string, 0)
	for _, val := range db.Snippets {
		if _, ok := newLangMap[val.Language]; !ok {
			newLangMap[val.Language] = 1
			newLangsList = append(newLangsList, val.Language)
		}
	}
	db.Languages = newLangsList
	updatedData, err := json.Marshal(db)
	if err != nil {
		return err
	}

	err = dbSave("./db.json", updatedData)
	if err != nil {
		return err
	}

	return nil
}

func DeleteSnippet(strid string) error {
	db, err := getDb()
	if err != nil {
		return err
	}
	id, err := strconv.Atoi(strid)
	if err != nil {
		return err
	}
	if id < 0 || id >= len(db.Snippets) {
		return errors.New("invalid id")
	}
	db.Snippets[id].Deleted = 1

	updatedData, err := json.Marshal(db)
	if err != nil {
		return err
	}

	err = dbSave("./db.json", updatedData)
	if err != nil {
		return err
	}

	return nil
}

func GetSnippet(id string) (Snippet, error) {
	db, err := getDb()
	if err != nil {
		return Snippet{}, err
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return Snippet{}, err
	}
	if idInt < 0 || idInt >= len(db.Snippets) {
		return Snippet{}, errors.New("invalid id")
	}
	return db.Snippets[idInt], nil
}

func EditSnippet(id string, snippet Snippet) error {
	db, err := getDb()
	if err != nil {
		return err
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	if idInt < 0 || idInt >= len(db.Snippets) {
		return errors.New("invalid id")
	}
	db.Snippets[idInt] = snippet

	updatedData, err := json.Marshal(db)
	if err != nil {
		return err
	}

	err = dbSave("./db.json", updatedData)
	if err != nil {
		return err
	}
	return nil
}
