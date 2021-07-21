package main

import (
	"encoding/json"
	"fmt"
	client "github.com/LampardNguyen234/go-ipfs-http-client"
	"net/http"
	"strconv"
)

func (s *server) pin(w http.ResponseWriter, r *http.Request) {
	params, err := parseArguments(r)
	if err != nil {
		Logger.Errorf("invalid arguments: %v\n", err)
		writeError(w, fmt.Errorf("invalid arguments"))
		return
	}
	Logger.Println(params)

	// get the `cid` parameter
	cIdStr, ok := params["cid"]
	if !ok {
		Logger.Error("cid not found")
		writeError(w, fmt.Errorf("cid not found"))
		return
	}

	// get the `pin` parameter
	var pin bool
	if _, ok := params["pin"]; !ok {
		pin = true
	} else {
		var err error
		pin, err = strconv.ParseBool(params["pin"])
		if err != nil {
			Logger.Error(err)
			writeError(w, fmt.Errorf("`pin` is invalid"))
			return
		}
	}

	// get the metadata file from the IPFS node
	Logger.Println("Retrieving the metadata file...")
	metadataFile, err := s.ipfsClient.Get(cIdStr, true)
	if err != nil {
		Logger.Error(err)
		writeError(w, fmt.Errorf("cannot retrieve cid file"))
		return
	}
	rawData, err := client.ParseNodeFile(metadataFile)
	if err != nil {
		Logger.Error(err)
		writeError(w, fmt.Errorf("cannot retrieve cid file"))
		return
	}
	Logger.Printf("metadata rawData size: %v\n", len(rawData))
	Logger.Printf("rawData: %v\n", string(rawData))
	metadata := make(map[string]interface{})
	err = json.Unmarshal(rawData, &metadata)
	if err != nil {
		Logger.Error(err)
		writeError(w, fmt.Errorf("cannot retrieve cid file"))
		return
	}

	// get the cID of the data file
	dataCID, ok := metadata["uri"]
	if !ok {
		Logger.Errorf("cannot find `uri` in metadata %v\n", metadata)
		writeError(w, fmt.Errorf("metadata file is invalid"))
		return
	}
	dataCIDStr, ok := dataCID.(string)
	if !ok {
		Logger.Errorf("%v is not a string\n", dataCID)
		writeError(w, fmt.Errorf("metadata file is invalid"))
		return
	}
	Logger.Printf("dataCID: %v\n", dataCIDStr)

	if pin {
		Logger.Printf("Pinning %v and %v...\n", cIdStr, dataCIDStr)
		err = s.ipfsClient.PinCID(cIdStr)
		if err != nil {
			writeError(w, err)
			return
		}
		err = s.ipfsClient.PinCID(dataCIDStr)
		if err != nil {
			writeError(w, err)
			return
		}
		writeResponse(w, &response{
			Result: map[string]interface{}{"status": "success"},
			Error:  nil,
		})
		Logger.Printf("Pinning succeeded!!\n")
	} else {
		Logger.Printf("Un-pinning %v and %v\n", cIdStr, dataCIDStr)
		err = s.ipfsClient.UnPinCID(cIdStr)
		if err != nil {
			Logger.Error(err)
			writeError(w, err)
			return
		}
		err = s.ipfsClient.UnPinCID(dataCIDStr)
		if err != nil {
			Logger.Error(err)
			writeError(w, err)
			return
		}
		writeResponse(w, &response{
			Result: map[string]interface{}{"status": "success"},
			Error:  nil,
		})
		Logger.Printf("Un-pinning succeeded!!\n")
	}
}