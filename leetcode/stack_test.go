package leetcode

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	res:=IsValid("()[]{}")
	if !res{
		t.Error("与预期不符")
	}
}

func TestCalculate(t *testing.T) {
	res:=Calculate2("()+-1234567890")
	fmt.Println(res)
}
