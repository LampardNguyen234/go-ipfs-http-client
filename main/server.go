package main

import (
	client "github.com/LampardNguyen234/go-ipfs-http-client"
	"github.com/gorilla/mux"
	"net/http"
)

type server struct {
	ipfsClient *client.Client
	localIPFSClient *client.Client
	errChan chan error
	resultChan chan interface{}
}

func newServer() (*server, error) {
	infuraAccount, err := client.ReadInfuraKey("../.secret")
	if err != nil {
		return nil, err
	}

	ipfsClient, err := client.NewInfuraClient(infuraAccount.ProjectId, infuraAccount.ProjectSecret)
	if err != nil {
		return nil, err
	}

	localClient, err := client.NewLocalClient()
	if err != nil {
		Logger.Errorf("cannot create local IPFS: %v\n", err)
	}

	return &server{ipfsClient: ipfsClient, localIPFSClient: localClient}, nil
}

func main() {
	Logger.Println("Starting server...")

	s, err := newServer()
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/add", s.add).Methods(http.MethodPost)
	r.HandleFunc("/pin", s.pin).Methods(http.MethodPost)
	Logger.Println(http.ListenAndServe(":8888", r))
}
