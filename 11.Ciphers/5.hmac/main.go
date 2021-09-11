package main

import (
	//"bufio"
	"crypto/hmac"
	"crypto/sha256"
	//"crypto/sha512"
	"encoding/hex"
	"fmt"
	//"os"
)
func CheckError(err error ){
	if err != nil {
		panic(err)
	}
}
const secret_key = "my key"

func main () {
	

	fmt.Println("hmac ...\n ")

	hh1:= hmac.New(sha256.New,[]byte(secret_key))
	// hh2:= hmac.New(sha512.New,[]byte(secret_key))

	hh1.Write([]byte("sk shah"))
	// hh2.Write([]byte("sk shahriar ahmed raka"))

	fmt.Println("hmac sha256 : ",hh1.Sum(nil))
	fmt.Println("hmac sha256 : ",hex.EncodeToString(hh1.Sum(nil)))
	// fmt.Println("hmac sha512 : ",hh2.Sum(nil))
	// fmt.Println("hmac sha512 : ",hex.EncodeToString(hh2.Sum(nil)))

	// fmt.Println("bool ",hmac.Equal(hh1.Sum(nil),hh2.Sum(nil)))
	hhx:= hmac.New(sha256.New,[]byte(secret_key))
	hhx.Write([]byte("sk shah"))

	fmt.Println("hmac sha256 : ",hhx.Sum(nil))
	fmt.Println("hmac sha256 : ",hex.EncodeToString(hhx.Sum(nil)))

	fmt.Println("compare : ",hmac.Equal(hh1.Sum(nil),hhx.Sum(nil)))




	




}

// // fmt.Printf("give the first password : ")
// 	// pass ,err := bufio.NewReader(os.Stdin).ReadString('\n')
// 	// CheckError(err)

// 	// h1 := sha256.New()

// 	// h1.Write([]byte(pass))

// 	// fmt.Println("sha256 : ",h1.Sum(nil))
// 	// fmt.Println("sha256 (string): ",hex.EncodeToString(h1.Sum(nil)))

// 	// fmt.Printf("\ngive the second password : ")
// 	// pass2 ,err := bufio.NewReader(os.Stdin).ReadString('\n')
// 	// CheckError(err)

// 	// h2 := sha256.New()

// 	// h2.Write([]byte(pass2))
// 	// fmt.Println("sha256 : ",h2.Sum(nil))
// 	// fmt.Println("sha256 (string): ",hex.EncodeToString(h2.Sum(nil)))