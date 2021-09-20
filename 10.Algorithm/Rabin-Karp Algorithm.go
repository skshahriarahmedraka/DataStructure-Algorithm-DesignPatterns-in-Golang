package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	//"strconv"
)

func main(){
	var s,p string
	fmt.Println("********** Rabin-Karp Algorithm *************")
	fmt.Printf("\ngive the text : ")
	a,_,_:= bufio.NewReader(os.Stdin).ReadLine()
	s=string(a)
	fmt.Printf("\nGive pattern : ")
	a,_,_= bufio.NewReader(os.Stdin).ReadLine()
	p=string(a)

	li:=RabinKarp(s,p)
	fmt.Println(li)
}



func RabinKarp(s string,p string)[]uint64{
	var mm map[string]uint64
	mm = map[string]uint64{
		"a":1,"b":2,"c":3,"d":4,"e":5,"f":6,"g":7,"h":8,"i":9,"j":10,"k":11,"l":12,"m":13,"n":14,"o":15,"p":16,"q":17,"r":18,"s":19,"t":20,"u":21,"v":22,"w":23,"x":24,"y":25,"z":26,
		"A":27,"B":28,"C":29,"D":30,"E":31,"F":32,"G":33,"H":34,"I":35,"J":36,"K":37,"L":38,"M":39,"N":40,"O":41,"P":42,"Q":43,"R":44,"S":45,"T":46,"U":47,"V":48,"W":49,"X":50,"Y":51,"Z":52,
	}

	slen:= uint64(len(s))
	plen := uint64(len(p))
	var num,num2 uint64
	num=0
	num2=0
	for i:= uint64(0);i<plen ;i++{
		s:=mm[string(s[i])]
		num+= s*uint64((math.Pow(100,float64(plen-1-i))))
	//fmt.Println(num)
	}

	for i:= uint64(0);i<plen ;i++{
		s:=mm[string(p[i])]
		num2+= s*uint64(math.Pow(100,float64(plen-1-i)) )
	//fmt.Println(num2)
	}
	var li []uint64
	for i:= uint64(0);i<slen-plen;i++{
		//fmt.Println("nums : ",num," ",num2)
		if num == num2 {
			li=append(li,i+1)
			num-= mm[string(s[i])]*uint64(math.Pow(100,float64(plen-1)))
			num*=100
			//num+=uint64(s[i+plen])
			num+=mm[string(s[i+plen])]

		}else {
			num-= mm[string(s[i])]*uint64(math.Pow(100,float64(plen-1)))
			num*=100
			//num+=uint64(s[i+plen])
			num+=mm[string(s[i+plen])]
		}

	}
	return li

}