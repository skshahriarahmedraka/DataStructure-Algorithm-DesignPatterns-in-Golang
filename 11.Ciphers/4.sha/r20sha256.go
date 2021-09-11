package main

// import (
// 	"crypto/sha256"
// 	"fmt"
// 	"io"
// 	"log"
// 	"os"
// )

// func main (){
// 	f,err := os.Open("r20.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer f.Close()

// 	h := sha256.New()

// 	fmt.Println(h)

// 	_,err = io.Copy(h,f)

// 	if err != nil {
// 		log.Panic(err)
// 	}
	


// }