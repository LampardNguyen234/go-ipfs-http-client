package main

import (
	"fmt"
	"golang.org/x/crypto/sha3"
	"io"
	"net/http"
	"os"
)

func (s *server) upload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(maxFileRequestSize)
	if err != nil {
		writeResponse(w, &response{
			Result: nil,
			Error:  fmt.Errorf("cannot retrieve file from request: %v", err),
		})
		return
	}

	dataFile, rawData, err := getDataFileFromRequest(r)
	if err != nil {
		writeResponse(w, &response{
			Result: nil,
			Error:  err,
		})
		return
	}

	fileName := fmt.Sprintf("%x", sha3.Sum256(rawData))
	if exists(fileName) {
		Logger.Errorf("file %v already exists\n", fileName)
		writeResponse(w, &response{
			Result: nil,
			Error:  fmt.Errorf("file already exists"),
		})
		return
	}

	f, err := os.OpenFile(buildFilePath(fileName), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		Logger.Errorf("cannot open file to store: %v\n", err)
		writeResponse(w, &response{
			Result: nil,
			Error:  err,
		})
		return
	}
	defer func() {
		err = f.Close()
		if err != nil {
			Logger.Error(err)
		}
	}()
	_, err = io.Copy(f, dataFile)
	if err != nil {
		Logger.Errorf("cannot read data of file: %v\n", err)
		writeResponse(w, &response{
			Result: nil,
			Error:  err,
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	writeResponse(w, &response{
		Result: map[string]interface{}{"fileName" : fileName},
		Error:  nil,
	})
}

