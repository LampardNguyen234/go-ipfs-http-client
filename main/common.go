package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type response struct {
	Result interface{} `json:"Result"`
	Error error `json:"Error"`
}

func (resp *response) MarshalJSON() ([]byte, error) {
	data := make(map[string]interface{})
	data["Result"] = resp.Result
	if resp.Error != nil {
		data["Error"] = resp.Error.Error()
	} else {
		data["Error"] = nil
	}

	return json.Marshal(data)
}

func writeResponse(w http.ResponseWriter, resp *response) {
	w.Header().Set("Content-Type", "application/json")

	respInBytes, err := json.Marshal(resp)
	if err != nil {
		Logger.Errorf("%v\n", err)
		return
	}
	fmt.Println(string(respInBytes))
	_, err = w.Write(respInBytes)
	if err != nil {
		Logger.Errorf("%v\n", err)
	}
}

func writeError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	resp := &response{
		Result: nil,
		Error:  err,
	}

	writeResponse(w, resp)
}

func parseArguments(r *http.Request) (map[string]string, error) {
	res := make(map[string]string)

	if r.URL == nil {
		return nil, fmt.Errorf("request does not contain an URL")
	}

	args := strings.Split(r.URL.String(), urlParamsSeparator)
	if len(args) <= 1 {
		return res, nil
	}

	args = strings.Split(args[1], argSeparator)
	for _, arg := range args {
		values := strings.Split(arg, keyValueSeparator)
		if len(values) != 2 {
			return nil, fmt.Errorf("invalid argument")
		}

		res[values[0]] = values[1]
	}

	return res, nil
}
