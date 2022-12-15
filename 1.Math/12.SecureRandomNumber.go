package main

import (
    "fmt"
    "crypto/rand"
    "math/big"
)

func main() {
    nBig, err := rand.Int(rand.Reader, big.NewInt(27))
    if err != nil {
        panic(err)
    }
    n := nBig.Int64()
    fmt.Printf("Here is a random %T in [0,27) : %d\n", n, n)
}