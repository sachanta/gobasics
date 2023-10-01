package main

import (
	"fmt"

	"github.com/sachanta/learngo/cryptit/decrypt"
	"github.com/sachanta/learngo/cryptit/encrypt"
)

func main() {
	fmt.Println(decrypt.Nimbus(encrypt.Nimbus("Srikar")))
}
