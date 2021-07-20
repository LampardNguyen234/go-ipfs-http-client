package client

import (
	"context"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/interface-go-ipfs-core/path"
)

// PinFile pins the file with the given cIdStr on the server.
func (c *Client) PinFile(cIdStr string) error {
	cId, err := cid.Parse(cIdStr)
	if err != nil {
		return err
	}

	return c.Pin().Add(context.Background(), path.IpfsPath(cId))
}

// IsPinned checks whether the content of the cIdStr is pinned on the server or not, as well as the reason why it's pinned.
func (c *Client) IsPinned(cIdStr string) (string, bool, error) {
	cId, err := cid.Parse(cIdStr)
	if err != nil {
		return "", false, err
	}

	return c.Pin().IsPinned(context.Background(), path.IpfsPath(cId))
}
