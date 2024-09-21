package model

import (
	"database/sql"
	"log"
	"reflect"
	"yomikyasu/internal/database"
)

const (
	CREATE_CONFIG_QUERY string = "INSERT INTO configs (use_natural_voice, speech_speed, full_text_service_url) VALUES (?, ?, ?);"
	GET_ALL_CONFIG      string = "SELECT * FROM configs"
    DELETE_CONFIG string = "DELETE FROM configs wehre id = ?"
)

func CreateConfig(db *database.Service, config *Config) (*sql.Result, error) {
	stmt, _ := (*db).Prepare(CREATE_CONFIG_QUERY)
	defer stmt.Close()

	result, err := stmt.Exec(config.UseNaturalVoice, config.SpeechSpeed, config.FullTextServiceUrl)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func GetConfigs(db *database.Service) ([]Config, error) {
	stmt, _ := (*db).Prepare(GET_ALL_CONFIG)

	result, err := (*stmt).Query()

	if err != nil {
		return nil, err
	}

	configs := make([]Config, 0)
	for result.Next() {
		c := Config{}

		s := reflect.ValueOf(&c).Elem()
		numCols := s.NumField()
		columns := make([]interface{}, numCols)
		for i := 0; i < numCols; i++ {
			field := s.Field(i)
			columns[i] = field.Addr().Interface()
		}

		err := result.Scan(columns...)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(c)
		configs = append(configs, c)
	}

	return configs, nil
}
