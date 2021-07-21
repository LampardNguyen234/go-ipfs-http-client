package client

import (
	"context"
	"fmt"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/interface-go-ipfs-core/path"
)

// PinCID pins the file with the given cIdStr on the server.
func (c *Client) PinCID(cIdStr string) error {
	cId, err := cid.Parse(cIdStr)
	if err != nil {
		return err
	}

	err = c.Pin().Add(context.Background(), path.IpfsPath(cId))
	if err != nil {
		return err
	}

	_, isPinned, err := c.IsPinned(cIdStr)
	if err != nil {
		return err
	}

	if !isPinned {
		return fmt.Errorf("something is wrong with the IPFS server")
	}
	return nil
}

// UnPinCID un-pins the file with the given cIdStr on the server.
func (c *Client) UnPinCID(cIdStr string) error {
	cId, err := cid.Parse(cIdStr)
	if err != nil {
		return err
	}

	err = c.Pin().Rm(context.Background(), path.IpfsPath(cId))
	if err != nil {
		return err
	}

	_, isPinned, err := c.IsPinned(cIdStr)
	if err != nil {
		return err
	}

	if isPinned {
		return fmt.Errorf("something is wrong with the IPFS server")
	}
	return nil
}

// IsPinned checks whether the content of the cIdStr is pinned on the server or not, as well as the reason why it's pinned.
func (c *Client) IsPinned(cIdStr string) (string, bool, error) {
	cId, err := cid.Parse(cIdStr)
	if err != nil {
		return "", false, err
	}

	return c.Pin().IsPinned(context.Background(), path.IpfsPath(cId))
}
