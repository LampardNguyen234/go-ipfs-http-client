package client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/LampardNguyen234/go-ipfs-http-client/httpapi"
	files "github.com/ipfs/go-ipfs-files"
	"github.com/ipfs/interface-go-ipfs-core/options"
	"github.com/ipfs/interface-go-ipfs-core/path"
	"io"
	"os"
	"strings"
)

// Add adds the content of `r` to the IPFS node with the given `name`.
// If you don't want to set the name, leave it empty.
func (c *Client) Add(r io.Reader, name string, isWrappedWithDirectory bool, addOptions ...options.UnixfsAddOption) (path.Resolved, error){
	unixFs, err := c.GetUnixFs()
	if err != nil {
		return nil, err
	}

	fr := files.NewReaderFile(r)
	nodeFile := httpapi.NewNode(name, fr)

	if isWrappedWithDirectory {
		return unixFs.AddWithWrapDirectory(context.Background(), nodeFile, addOptions...)
	}
	return c.Unixfs().Add(context.Background(), nodeFile, addOptions...)
}

// AddWithMetadata adds the content of `r` to the IPFS node with the given metadata.
//
// A metadata file with default name (MetadataFileName) will be created and stored on the IPFS network. The metadata file contains the IPFS URI of the added file.
//
// This function also pins the files to the IPFS node.
func (c *Client) AddWithMetadata(r io.Reader, metadata map[string]interface{}, addOptions ...options.UnixfsAddOption) (path.Resolved, error){
	unixFs, err := c.GetUnixFs()
	if err != nil {
		return nil, err
	}

	// First, we add the true file
	fr := files.NewReaderFile(r)
	nodeFile := httpapi.NewNode("", fr)
	cId, err := c.Unixfs().Add(context.Background(), nodeFile, addOptions...)
	if err != nil {
		return nil, err
	}

	// Store the metadata file
	metadata["uri"] = cId.String()
	metadataBytes, err := json.Marshal(metadata)
	if err != nil {
		return nil, err
	}
	metadataFile := files.NewBytesFile(metadataBytes)
	metadataNodeFile := httpapi.NewNode(MetadataFileName, metadataFile)
	return unixFs.Add(context.Background(), metadataNodeFile, addOptions...)
}

// AddFileFromPath adds the `file` from the given path with its stat to the IPFS node.
func (c *Client) AddFileFromPath(filePath string, isWrappedWithDirectory bool, addOptions ...options.UnixfsAddOption) (path.Resolved, error) {
	unixFs, err := c.GetUnixFs()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	stringsArgs := strings.Split(filePath, "/")
	if len(stringsArgs) == 0 {
		return nil, fmt.Errorf("filePath is not valid")
	}

	fr := files.NewReaderFile(file)
	nodeFile := httpapi.NewNode(stringsArgs[len(stringsArgs) - 1], fr)

	if isWrappedWithDirectory {
		return unixFs.AddWithWrapDirectory(context.Background(), nodeFile, addOptions...)
	}
	return unixFs.Add(context.Background(), nodeFile, addOptions...)
}

// AddFileFromPathWithMetadata adds the content of `r` to the IPFS node with the given metadata.
//
// A metadata file with default name (MetadataFileName) will be created and stored on the IPFS network. The metadata file contains the IPFS URI of the added file.
//
// This function also pins the files to the IPFS node.
func (c *Client) AddFileFromPathWithMetadata(filePath string, metadata map[string]interface{}, addOptions ...options.UnixfsAddOption) (path.Resolved, error){
	unixFs, err := c.GetUnixFs()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	stringsArgs := strings.Split(filePath, "/")
	if len(stringsArgs) == 0 {
		return nil, fmt.Errorf("filePath is not valid")
	}

	// First, we add the true file
	fr := files.NewReaderFile(file)
	nodeFile := httpapi.NewNode(stringsArgs[len(stringsArgs) - 1], fr)
	cId, err := c.Unixfs().Add(context.Background(), nodeFile, addOptions...)
	if err != nil {
		return nil, err
	}

	// Store the metadata file
	metadata["uri"] = cId.String()
	metadataBytes, err := json.Marshal(metadata)
	if err != nil {
		return nil, err
	}
	metadataFile := files.NewBytesFile(metadataBytes)
	metadataNodeFile := httpapi.NewNode(MetadataFileName, metadataFile)
	return unixFs.Add(context.Background(), metadataNodeFile, addOptions...)
}