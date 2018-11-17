package main

import (
	"fmt"

	nonce "github.com/LarryBattle/nonce-golang"
)

func main() {
	// var err error
	// id := uuid.New()
	// if err != nil {
	// 	fmt.Println("uuid.New()")
	// }

	// hex, _ := id.MarshalBinary()
	// idd, _ := uuid.FromBytes(hex)

	// fmt.Println(id)
	// fmt.Println(hex)
	// fmt.Println(idd)

	nonce := nonce.NewToken()
	fmt.Print(len(nonce))
}
