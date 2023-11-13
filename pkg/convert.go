package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"io"
	"os"
)

func ConvertFile(filename string) (interface{}, error) {
	// read the bytes of the file in filename
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file %s: %w", filename, err)

	}
	defer file.Close()

	jsonData, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading file %s: %v", filename, err)
	}

	return ConvertString(jsonData)
}

func ConvertString(jsonString []byte) (interface{}, error) {
	var doc interface{}
	vr, err := bsonrw.NewExtJSONValueReader(bytes.NewReader(jsonString), true)
	if err != nil {
		return nil, fmt.Errorf("error creating bsonrw.NewExtJSONValueReader: %w", err)
	}
	decoder, err := bson.NewDecoder(vr)
	if err != nil {
		return nil, fmt.Errorf("error creating bson.NewDecoder: %w", err)
	}

	err = decoder.Decode(&doc)
	if err != nil {
		return nil, fmt.Errorf("error decoding bson: %w", err)
	}

	return doc, nil
}

func ConvertMapStringInterface(m map[string]interface{}) (interface{}, error) {
	j, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("error marshaling map[string]interface{}: %w", err)
	}
	return ConvertString(j)
}
