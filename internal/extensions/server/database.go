package server

import (
	json "ToDoApp/Internal/services/JSON"
	"ToDoApp/internal/services/files"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func ConnectToDB() (*sqlx.DB, error) {
	var filePath string = "../../configs/conf.json"

	configFileData, err := files.ReadFile(filePath)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer configFileData.Close()

	jsonParser := json.DecodeToJSON(configFileData)
	jsonParser.Decode(&config)

	return sqlx.Connect("postgres", "user=foo dbname=bar sslmode=disable")
}
