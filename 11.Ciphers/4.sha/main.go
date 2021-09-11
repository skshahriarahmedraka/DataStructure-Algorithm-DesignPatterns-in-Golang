package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

func main(){
	h1 := sha256.New()
	h1.Write([]byte("sk shahriar"))
	fmt.Printf("var type h1.Sum(nil) : %T\n",h1.Sum(nil))

	fmt.Printf("sha256 hash(hex) : %x\n",h1.Sum(nil))
	s:= hex.EncodeToString(h1.Sum(nil))
	fmt.Println("sha256 hash(string) : ",s)

	h2 := sha512.New()
	h2.Write([]byte("ahmed raka"))
	fmt.Printf("\n\nsha512 hash(hex) : %x\n",h2.Sum(nil))
	s2:= hex.EncodeToString(h1.Sum(nil))
	fmt.Println("sha512 hash(string) : ",s2)
	fmt.Printf("var type : %T\n",s2)
}