package leetcode

import "testing"

func TestConvert(t *testing.T) {
	if Convert("PAYPALISHIRING",3)!="PAHNAPLSIIGYIR"{
		t.Error("error:",Convert("PAYPALISHIRING",3))
	}
}
