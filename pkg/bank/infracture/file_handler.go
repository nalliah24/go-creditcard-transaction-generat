// This package handles reading and writing to json files
package infracture

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	m "github.com/nalliah24/go-creditcard-transaction-generator/pkg/bank/model"
)

// Reads config file based on path. default to cmd folder
func ReadConfig(filepath string) ([]m.Config, error) {
	file, err := ioutil.ReadFile(filepath)

	if err != nil {
		return []m.Config{}, errors.New("error reading config file " + err.Error())
	}

	var c []m.Config
	json.Unmarshal(file, &c)

	return c, nil
}

// Write to json
func WriteJson(filename string, trans interface{}) error {
	file, _ := os.OpenFile(filename, os.O_CREATE, os.ModePerm)
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	if err := encoder.Encode(trans); err != nil {
		return err
	}
	return nil
}
