package client

import (
	"fmt"
	"testing"
)

func TestClient_PinFile(t *testing.T) {
	cIdStr := "QmS2gWrAzp42swfjweNQgzHL4dC1UVR6a2rokPJxMfgPfM"
	why, isPinned, err := ipfsClient.IsPinned(cIdStr)
	if err != nil {
		panic(err)
	}
	if isPinned {
		panic("already pinned")
	}

	err = ipfsClient.PinCID(cIdStr)
	if err != nil {
		panic(err)
	}

	why, isPinned, err = ipfsClient.IsPinned(cIdStr)
	if err != nil {
		panic(err)
	}
	if !isPinned {
		panic("should pin file")
	}

	fmt.Println("File has been pinned with reason:", why)
}

func TestClient_UnPinFile(t *testing.T) {
	cIdStr := "QmS2gWrAzp42swfjweNQgzHL4dC1UVR6a2rokPJxMfgPfM"
	_, isPinned, err := ipfsClient.IsPinned(cIdStr)
	if err != nil {
		panic(err)
	}
	if !isPinned {
		err = ipfsClient.PinCID(cIdStr)
		if err != nil {
			panic(err)
		}
	}

	_, isPinned, err = ipfsClient.IsPinned(cIdStr)
	if err != nil {
		panic(err)
	}
	if !isPinned {
		panic("should pin file")
	}

	err = ipfsClient.UnPinCID(cIdStr)
	if err != nil {
		panic(err)
	}

	_, isPinned, err = ipfsClient.IsPinned(cIdStr)
	if err != nil {
		panic(err)
	}
	if isPinned {
		panic("should un-pin file")
	}

	fmt.Println("Success!!")
}

func TestClient_IsPinned(t *testing.T) {
	cIdStr := "QmTcwAhqcX1GhJpBj7SaFHaYFLc5JjtZU2b2J2qe66n6Ua"

	why, isPinned, err := ipfsClient.IsPinned(cIdStr)
	if err != nil {
		panic(err)
	}

	fmt.Println(why, isPinned, err)
}
