package db

import (
	"encoding/json"
)

func GetAllSnippets() []Snippet {
	db := DB{}
	getDb(&db)
	return db.Snippets
}

func GetAllLangs() []string {
	db := DB{}
	getDb(&db)
	return db.Languages
}

func SaveSnippet(snippet Snippet) error {
	db := DB{}
	getDb(&db)

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
	db := DB{}
	getDb(&db)

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
