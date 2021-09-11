package main

import "fmt"

func RomanToInteger(roman string)uint64{
	var ro =map[string]uint64{"I":1,"V":5,"X":10,"L":50,
		"C":100,"D":500,"M":1000} //,"_V":5000,"_X":10000,"_L":50000, "_C":100000,"_D":500000,"_M":1000000
	var (
		t uint64
		slen uint64
		h uint64


	)
	slen =uint64(len(roman))

	for  ;h<slen;h++ {
		if (h+1 <slen)  && ro[string(roman[h])] < ro[string(roman[h+1])] {
			t -= ro[string(roman[h])]
		}else {
			t += ro[string(roman[h])]
		}
	}
	return t


}

func main(){
	fmt.Printf("give a roman number (like VIII) : ")
	var (
		s string
		n uint64
	)

	fmt.Scanln(&s)
	n=RomanToInteger(s)
	fmt.Println("integer value : ",n)
}