package main

import (
	"bufio"
	"fmt"
	"os"
	//"golang.org/x/crypto/bcrypt"
)

func main(){
	fmt.Printf(" password for bcrypt : ")
	b,_ := bufio.NewReader(os.Stdin).ReadString('\n')
	scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan() // use `for scanner.Scan()` to keep reading
    b2 := scanner.Text()
	// var pass string
	var b3 string 
	fmt.Scanln(&b3)

	fmt.Printf("%T \n",&b)
	fmt.Println(b)
	fmt.Printf("%T \n",&b2)
	fmt.Println(b2)
	fmt.Printf("%T \n",&b3)
	fmt.Println(b3)
}