package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"math/big"
	"strings"
)

func main() {
	id, _ := uuid.NewV4()

	var i big.Int
	i.SetString(strings.Replace(id.String(), "-", "", 4), 16)
	// or if your uuid is [16]byte
	// i.SetBytes(uuid[:])
	fmt.Println(i.String())
}
