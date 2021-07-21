package client

import (
	"fmt"
	files "github.com/ipfs/go-ipfs-files"
	"testing"
)

func TestClient_Get(t *testing.T) {
	cIdStr := "QmP13aRuCpiJoHzPjJC9hbV6iZWc7tVutvoGKHZHss9bF5"

	fileNode, err := ipfsClient.Get(cIdStr, true)
	if err != nil {
		panic(err)
	}
	data := make([]byte, 1000)
	file, ok := fileNode.(files.File)
	if !ok {
		panic("not ok")
	}
	n, err := file.Read(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(n, data)
	fmt.Println(data[:n])
	fmt.Println(string(data[:n]))
}
