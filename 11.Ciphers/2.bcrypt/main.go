package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	//////////////////////////
	"golang.org/x/crypto/bcrypt"
)
var start time.Time
func Panicking(err error ,msg string){
	if err != nil {
		log.Panicln(msg," ERROR : ",err)
	}
}

func main(){


	fmt.Printf(" password for bcrypt : ")
	s,_ := bufio.NewReader(os.Stdin).ReadString('\n')

	start = time.Now()

	hashedPass :=makeHash(s)

	fmt.Println("time of execution : ",time.Since(start))

	fmt.Println("hashed []byte :",hashedPass)
	fmt.Println("hashed string :",string(hashedPass))
	fmt.Printf("\n\n")
	b := validateHash(s,hashedPass)

	fmt.Println(" validateHash : ",b)


}

func makeHash(pass string) []byte {
	hPass,err := bcrypt.GenerateFromPassword([]byte(pass),bcrypt.DefaultCost) 
	// bcrypt.DefaultCost // password for bcrypt : sk shahriar ahmed raka  // time of execution :  85.867127ms
	Panicking(err,"GenerateFromPassword ")
	
	return hPass
}

func validateHash(pass string , hashPass []byte ) bool {
	err := bcrypt.CompareHashAndPassword(hashPass,[]byte(pass))
	
	return err ==  nil

}