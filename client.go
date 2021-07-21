package client

import (
	"encoding/base64"
	"fmt"
	"github.com/LampardNguyen234/go-ipfs-http-client/httpapi"
	"net/http"
)

// Client implements an HTTP client for interacting with IPFS nodes.
type Client struct {
	*httpapi.HttpApi
}

// NewLocalClient creates a client pointing to the local IPFS node.
func NewLocalClient() (*Client, error) {
	c := &http.Client{
		Transport: &http.Transport{
			Proxy:             http.ProxyFromEnvironment,
			DisableKeepAlives: true,
		},
	}

	client, err := httpapi.NewURLApiWithClient(LocalIPFS, c)
	if err != nil {
		return nil, err
	}

	return &Client{client}, nil
}

// NewClient returns a new Client pointing to the given url.
func NewClient(url string) (*Client, error) {
	c := &http.Client{
		Transport: &http.Transport{
			Proxy:             http.ProxyFromEnvironment,
			DisableKeepAlives: true,
		},
	}

	client, err := httpapi.NewURLApiWithClient(url, c)
	if err != nil {
		return nil, err
	}

	return &Client{client}, nil
}

// NewInfuraClient creates a new Client pointing to the Infura service.
func NewInfuraClient(projectId, projectSecret string) (*Client, error) {
	client, err := NewClient(InfuraEndPoint)
	if err != nil {
		return nil, err
	}
	client.Headers.Add("Authorization", "Basic " + infuraBasicAuth(projectId, projectSecret))

	return client, nil
}

// GetUnixFs returns the httpapi.UnixfsAPI of a Client.
func (c *Client) GetUnixFs() (*httpapi.UnixfsAPI, error) {
	tmpUnixFs := c.Unixfs()
	unixFs, ok := tmpUnixFs.(*httpapi.UnixfsAPI)
	if !ok {
		return nil, fmt.Errorf("cannot parse the UnixFs to an httpapi.UnixfsAPI")
	}

	return unixFs, nil
}

func infuraBasicAuth(projectId, projectSecret string) string {
	auth := projectId + ":" + projectSecret
	return base64.StdEncoding.EncodeToString([]byte(auth))
}