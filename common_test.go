package client

import (
	"fmt"
	"testing"
)

func TestReadInfuraKey(t *testing.T) {
	fileName := ".secret"
	infuraKey, err := ReadInfuraKey(fileName)
	if err != nil {
		panic(err)
	}

	fmt.Println(infuraKey)
}
