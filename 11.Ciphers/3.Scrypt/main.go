package main

// Function scrypt
//    Inputs: This algorithm includes the following parameters:
//       Passphrase:                Bytes    string of characters to be hashed
//       Salt:                      Bytes    string of random characters that modifies the hash to protect against Rainbow table attacks
//       CostFactor (N):            Integer  CPU/memory cost parameter - Must be a power of 2 (e.g. 1024)
//       BlockSizeFactor (r):       Integer  blocksize parameter, which fine-tunes sequential memory read size and performance. (8 is commonly used)
//       ParallelizationFactor (p): Integer  Parallelization parameter. (1 .. 232-1 * hLen/MFlen)
//       DesiredKeyLen (dkLen):     Integer  Desired key length in bytes (Intended output length in octets of the derived key; a positive integer satisfying dkLen ≤ (232− 1) * hLen.)
//       hLen:                      Integer  The length in octets of the hash function (32 for SHA256).
//       MFlen:                     Integer  The length in octets of the output of the mixing function (SMix below). Defined as r * 128 in RFC7914.
//    Output:
//       DerivedKey:                Bytes    array of bytes, DesiredKeyLen long

import (
	"fmt"
	"os"
	"bufio"
	"encoding/base64"

	"golang.org/x/crypto/scrypt"
)

func main (){
	var CostFactor,BlockSizeFactor,ParallelizationFactor,DesiredKeyLen int
	// N(CostFactor) power of 2 and greater than 1
	// r(BlockSizeFactor)
	// P(paralization factor) // r and p must satisfy r * p < 2^30

	fmt.Printf("password for Scrypt : ")
	pass,err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err !=nil {
		fmt.Errorf("error : %w\n",&err)
	}

	fmt.Printf("salt for Scrypt : ")
	salt,err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err !=nil {
		fmt.Errorf("error : %w\n",err)
	}

	fmt.Printf("CostFactor for Scrypt : ")
	fmt.Scanln(&CostFactor)
	
	fmt.Printf("BlockSizeFactor for Scrypt : ")
	fmt.Scanln(&BlockSizeFactor)
	
	fmt.Printf("ParallelizationFactor for Scrypt : ")
	fmt.Scanln(&ParallelizationFactor)
	
	fmt.Printf("DesiredKeyLen for Scrypt : ")
	fmt.Scanln(&DesiredKeyLen)

	

	dk,err := scrypt.Key([]byte(pass),[]byte(salt),CostFactor,BlockSizeFactor,ParallelizationFactor,DesiredKeyLen)

	if err != nil {
		fmt.Errorf("error : %w \n",err)
	}

	fmt.Println("scrypt hash : ",dk)
	fmt.Println("scrypt hash (string) : ",string(dk))
	fmt.Println(base64.StdEncoding.EncodeToString(dk))


}



