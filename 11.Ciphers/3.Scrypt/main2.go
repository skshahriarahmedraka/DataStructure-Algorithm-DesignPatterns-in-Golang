package main

import (
    "fmt"
	"bufio"
   "os"
)

func main() {
    // var s string

    // fmt.Scanf("%s", &s)
	s,_ := bufio.NewReader(os.Stdin).ReadSlice('\n')
    fmt.Println("you said", int(s))

// 	scanner := bufio.NewScanner(os.Stdin)
//   for scanner.Scan() {
//     fmt.Println(scanner.Text())
//   }
}