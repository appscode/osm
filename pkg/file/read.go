package file

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

func readFileAs(path string, obj interface{}) bool {
	d, err := ioutil.ReadFile(path)
	if err != nil {
		return false
	}
	err = json.Unmarshal(d, obj)
	if err != nil {
		return false
	}
	return true
}

func GetGCSCred(path string) (map[string]string, error) {
	data := make(map[string]string)
	if !readFileAs(path, &data) {
		return nil, errors.New("Credential not found")
	}
	return data, nil
}
