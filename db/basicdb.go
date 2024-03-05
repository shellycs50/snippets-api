package db

import (
	"encoding/json"
	"os"
)

func dbSave(filename string, data []byte) error {
	err := os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func dbRead(filename string) ([]byte, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func getDb(location *DB) error {
	file, err := dbRead("./db.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &location)
	if err != nil {
		return err
	}
	return nil
}
