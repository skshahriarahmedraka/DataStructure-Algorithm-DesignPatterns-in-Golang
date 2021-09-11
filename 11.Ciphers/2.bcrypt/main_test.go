package main 

import(
	"testing"
)

func TestmakeHash(t *testing.T){
	if len(string(makeHash("sk"))) <2 {
		t.Error("make hash error")
	}
}