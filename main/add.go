package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ipfs/interface-go-ipfs-core/options"
	"github.com/ipfs/interface-go-ipfs-core/path"
	"io"
	"net/http"
	"time"
)

func (s *server) add(w http.ResponseWriter, r *http.Request) {
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

	metadata, err := getMetadataFromRequest(r)
	if err != nil {
		writeResponse(w, &response{
			Result: nil,
			Error:  err,
		})
		return
	}

	var cId path.Resolved
	if s.localIPFSClient != nil {
		Logger.Println("Calculating local cID...")
		localDataFile := bytes.NewReader(rawData)
		cId, err = s.localIPFSClient.AddWithMetadata(localDataFile, metadata, options.Unixfs.HashOnly(true))
		if err != nil {
			writeResponse(w, &response{
				Result: nil,
				Error:  err,
			})
			return
		}
		Logger.Printf("cId from local: %v\n", cId.String())

		// call go-routine to upload file to the IPFS node
		go s.addFileToIPFS(dataFile, metadata)
	} else {
		Logger.Println("Adding file to the IPFS node...")
		localDataFile := bytes.NewReader(rawData)
		cId, err = s.ipfsClient.AddWithMetadata(localDataFile, metadata, options.Unixfs.HashOnly(true))
		if err != nil {
			writeResponse(w, &response{
				Result: nil,
				Error:  err,
			})
			return
		}
		Logger.Printf("cId: %v\n", cId.String())
	}

	writeResponse(w, &response{
		Result: map[string]interface{}{"ipfsHash" : cId.String()},
		Error:  err,
	})
}

func (s *server) addFileToIPFS(dataFile io.Reader, metadata map[string]interface{}) {
	start := time.Now()
	Logger.Println("Start uploading files...")
	cId, err := s.ipfsClient.AddWithMetadata(dataFile, metadata)
	if err != nil {
		Logger.Error(err)
		return
	}
	Logger.Printf("Uploading finished, timeElapsed: %v\n", time.Since(start).Seconds())
	Logger.Printf("cId from infura: %v\n", cId.String())
}

func getMetadataFromRequest(r *http.Request) (map[string]interface{}, error) {
	if r.MultipartForm == nil {
		return nil, fmt.Errorf("the request must be multipart/form-data")
	}

	metadataFields, ok := r.MultipartForm.Value[metadataKey]
	if !ok {
		return nil, fmt.Errorf("a `%v` field is required", metadataKey)
	}
	metadata := make(map[string]interface{})
	if len(metadataFields) != 0 {
		err := json.Unmarshal([]byte(metadataFields[0]), &metadata)
		if err != nil {
			return nil, fmt.Errorf("cannot parse metadata: %v", err)
		}
	}
	return metadata, nil
}
