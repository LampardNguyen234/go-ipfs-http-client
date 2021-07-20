package client

import (
	"crypto/rand"
	"log"
	"os"
	"testing"
)

var ipfsClient *Client

const (
	testDataInputDirectory = "./testdata/inputs/"
	testDataOutputDirectory = "./testdata/outputs/"
)

//func init() {
//	log.Printf("This runs before test\n")
//	var err error
//	ipfsClient, err = NewLocalClient()
//	if err != nil {
//		panic(err)
//	}
//}

func init() {
	log.Printf("This runs before test\n")
	infuraAccount, err := readInfuraKey(".secret")
	if err != nil {
		panic(err)
	}

	ipfsClient, err = NewInfuraClient(infuraAccount.ProjectId, infuraAccount.ProjectSecret)
	if err != nil {
		panic(err)
	}
}

func createRandomFile(fileName string, length int) error {
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer func() {
		err = f.Close()
		if err != nil {
			log.Printf("%v\n", err)
		}
	}()

	err = f.Truncate(0)
	if err != nil {
		return err
	}

	randData := make([]byte, length)
	_, err = rand.Read(randData)
	if err != nil {
		return err
	}
	_, err = f.Write(randData)
	if err != nil {
		return err
	}

	return nil
}

func TestAdd(t *testing.T) {
	fileName := "readme"
	err := createRandomFile(testDataInputDirectory+ fileName, 100000)
	if err != nil {
		panic(err)
	}

	f, err := os.Open(testDataInputDirectory + fileName)
	if err != nil {
		panic(err)
	}

	cId, err := ipfsClient.Add(f, fileName,true)
	if err != nil {
		panic(err)
	}

	log.Printf("Adding file successfully! CID: %v\n", cId)
}

func TestClient_AddWithMetadata(t *testing.T) {
	fileName := "readme"
	err := createRandomFile(testDataInputDirectory+ fileName, 100000)
	if err != nil {
		panic(err)
	}

	f, err := os.Open(testDataInputDirectory + fileName)
	if err != nil {
		panic(err)
	}

	metadata := map[string]string {
		"fileName": "readme",
		"author": "Lampard",
	}

	cId, err := ipfsClient.AddWithMetadata(f, metadata)
	if err != nil {
		panic(err)
	}

	log.Printf("Adding file successfully! CID: %v\n", cId)
}

func TestAddFileFromPath(t *testing.T) {
	fileName := "readme"
	err := createRandomFile(testDataInputDirectory+ fileName, 100000)
	if err != nil {
		panic(err)
	}

	cId, err := ipfsClient.AddFileFromPath(testDataInputDirectory + fileName, true)
	if err != nil {
		panic(err)
	}

	log.Printf("Adding file successfully! CID: %v\n", cId)
}

func TestClient_AddFileFromPathWithMetadata(t *testing.T) {
	fileName := "sample.jpg" // put the file inside the `testDataInputDirectory`

	metadata := map[string]string {
		"fileName": "readme",
		"author": "Lampard",
	}

	cId, err := ipfsClient.AddFileFromPathWithMetadata(testDataInputDirectory + fileName, metadata)
	if err != nil {
		panic(err)
	}

	log.Printf("Adding file successfully! CID: %v\n", cId)
}