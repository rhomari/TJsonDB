package TJsonDB

import (
	"encoding/json"
	"errors"
	"os"
	"reflect"
)

type Returns interface {
	string | map[string]interface{}
}
type Record map[string]interface{}

type Document struct {
	Records      []Record
	RecordsCount int64
}

func (d *Document) GetRecord(Index int64) Record {

	return d.Records[Index]
}
func (R *Record) Get(Key string) ([]Record, error) {
	results := []Record{}
	if reflect.TypeOf((*R)[Key]).Kind() == reflect.Slice {
		v := (*R)[Key].([]interface{})

		for i := range v {

			results = append(results, v[i].(map[string]interface{}))

		}
		return results, nil

	}
	if reflect.TypeOf((*R)[Key]).Kind() == reflect.Map {
		results = append(results, (*R)[Key].(map[string]interface{}))
		return results, nil
	}
	return nil, errors.New("Use [\"key\"] to access the value")
}

type TJsonDB struct {
}

func (T *TJsonDB) OpenDocument(DocumentName string) (Document, error) {
	document := Document{}
	bytes, err := os.ReadFile(DocumentName)
	if err != nil {
		return document, err

	}
	err = json.Unmarshal(bytes, &document.Records)
	if err != nil {
		return document, err
	}
	document.RecordsCount = int64(len(document.Records))
	return document, nil
}
