package main

import "fmt"




func GenerateChannel(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func Sieve(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

func Generate(limit int) []int {
	var primes []int

	ch := make(chan int)
	go GenerateChannel(ch)

	for i := 0; i < limit; i++ {
		primes = append(primes, <-ch)
		ch1 := make(chan int)
		go Sieve(ch, ch1, primes[i])
		ch = ch1
	}

	return primes
}



func main(){
	fmt.Println(Generate(10000))
}
