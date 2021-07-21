package client

import (
	"context"
	"github.com/ipfs/go-cid"
	files "github.com/ipfs/go-ipfs-files"
	"github.com/ipfs/interface-go-ipfs-core/path"
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
