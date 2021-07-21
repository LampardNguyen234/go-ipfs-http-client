package client

import (
	"fmt"
	files "github.com/ipfs/go-ipfs-files"
	"testing"
)

func TestClient_Get(t *testing.T) {
	cIdStr := "QmXwAzHmBLkSGLAnfqWXXshLQiJt6D2uHczoU6ea1EhmwQ"

	fileNode, err := ipfsClient.Get(cIdStr, true)
	if err != nil {
		panic(err)
	}
	data := make([]byte, 1000)
	fmt.Println(fileNode.Size())
	file, ok := fileNode.(files.File)
	if !ok {
		panic("not ok")
	}
	n, err := file.Read(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(n, data)
}
