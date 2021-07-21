package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	dataFileNameKey = "data"
	metadataKey     = "metadata"
	dataPath        = "./tmpData"
)

func getDataFileFromRequest(r *http.Request) (io.Reader, []byte, error) {
	if r.MultipartForm == nil {
		return nil, nil, fmt.Errorf("the request must be multipart/form-data")
	}
	dataFileHeaders, ok := r.MultipartForm.File[dataFileNameKey]
	if !ok {
		return nil, nil, fmt.Errorf("a `%v` file is required", dataFileNameKey)
	}

	if len(dataFileHeaders) == 0 {
		return nil, nil, fmt.Errorf("no file found")
	}
	dataFileHeader := dataFileHeaders[0]
	if dataFileHeader.Size > maxFileRequestSize {
		return nil, nil, fmt.Errorf("data file size exceeds maximum")
	}

	dataFile, err := dataFileHeader.Open()
	if err != nil {
		return nil, nil, fmt.Errorf("cannot open the file: %v", err)
	}
	data := make([]byte, dataFileHeader.Size)
	_, err = dataFile.Read(data)
	if err != nil {
		return nil, nil, err
	}
	reader := bytes.NewReader(data)

	return reader, data, nil
}

func buildFilePath(fileName string) string {
	return fmt.Sprintf("%v/%v", dataPath, fileName)
}

func exists(fileName string) bool {
	_, err := os.Stat(buildFilePath(fileName))
	if os.IsNotExist(err) {
		return false
	}

	return err == nil
}

func deleteFile(fileName string) error {
	if exists(fileName) {
		return os.Remove(buildFilePath(fileName))
	}

	return fmt.Errorf("file not exist")
}
