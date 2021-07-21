package client

import (
	"encoding/json"
	"io/ioutil"
)

type InfuraAccount struct {
	ProjectId     string `json:"ProjectId"`
	ProjectSecret string `json:"ProjectSecret"`
}

func ReadInfuraKey(fileName string) (*InfuraAccount, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var res InfuraAccount
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
