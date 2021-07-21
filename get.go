package client

import (
	"context"
	"errors"
	"fmt"
	"github.com/ipfs/go-cid"
	files "github.com/ipfs/go-ipfs-files"
	"github.com/ipfs/interface-go-ipfs-core/path"
	"io"
)

// Get returns the content with the given cId.
func (c *Client) Get(cIdStr string, isFile bool) (files.Node, error) {
	unixFs, err := c.GetUnixFs()
	if err != nil {
		return nil, err
	}

	cId, err := cid.Parse(cIdStr)
	if err != nil {
		return nil, err
	}

	return unixFs.GetContent(context.Background(), path.IpfsPath(cId), isFile)
}

// ParseNodeFile parses a files.Node (usually returned by a `get` command) to a raw byte array.
// The caller must manage the buffer size allow
func ParseNodeFile(nodeFile files.Node) ([]byte, error) {
	data := make([]byte, 0)
	f, ok := nodeFile.(files.File)
	if !ok {
		return nil, fmt.Errorf("cannot parse file")
	}

	buffer := make([]byte, 1000)
	for {
		n, err := f.Read(buffer)
		if err != nil {
			if errors.Is(err, io.EOF) {
				data = append(data, buffer[:n]...)
				break
			}
			return nil, err
		}
		fmt.Printf("%v, %v\n", len(data), n)
		if n < len(buffer) {
			buffer = buffer[:n]
			data = append(data, buffer...)
			break
		}
		data = append(data, buffer...)
	}

	return data, nil
}
