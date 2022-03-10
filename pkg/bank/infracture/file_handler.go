// This package handles reading and writing to json files
package infracture

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strings"

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

// Write to json based on empty interface, so it can write transactions as well summary outputs
func WriteJson(filename string, data interface{}, isIndent bool) error {
	file, _ := os.OpenFile(filename, os.O_CREATE, os.ModePerm)
	defer file.Close()
	encoder := json.NewEncoder(file)
	if isIndent {
		encoder.SetIndent("", " ")
	} else {
		encoder.SetIndent("", "")
	}

	if err := encoder.Encode(data); err != nil {
		return err
	}
	return nil
}

// Deleate a file name starts with, arguments path and filename
func DeleteFileNameStartsWith(path string, fnStarts string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for _, file := range files {
		if strings.HasPrefix(file.Name(), fnStarts) {
			err := os.Remove(path + "/" + file.Name())
			if err != nil {
				return err
			}
		}
	}
	return nil
}
