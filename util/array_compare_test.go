package util

import (
	"testing"
)

func newTestExample(args...interface{})[]interface{}{
	var a []interface{}
	for _,value:=range args{
		a= append(a, value)
	}
	return a
}

func TestSliceEqual(t *testing.T) {
	createTestSliceEqual(t,newTestExample(1,2,3),newTestExample(1,2,3),true)
	createTestSliceEqual(t,newTestExample('a','b','c'),newTestExample('a','b','c'),true)
	createTestSliceEqual(t,newTestExample(1.2,2.2,3.3),newTestExample(1.2,2.2,3.3),true)
	createTestSliceEqual(t,newTestExample("nihao","hello","world"),newTestExample("nihao","hello","world"),true)
	createTestSliceEqual(t,nil,nil,true)
	createTestSliceEqual(t,nil,newTestExample(1,2,3),false)
	createTestSliceEqual(t,newTestExample('a','b','c'),newTestExample(97,98,99),false)
	createTestSliceEqual(t,newTestExample(1,2,3),newTestExample(4,2,1),false)
	createTestSliceEqual(t,newTestExample(1,2,3),newTestExample(1,2,3,4),false)
	createTestSliceEqual(t,newTestExample("a","b","c"),newTestExample("A","B","C"),false)
}

func createTestSliceEqual(t *testing.T,a,b []interface{},equal bool){
	t.Helper()
	if equal{
		if !SliceEqual(a,b){
			t.Errorf("expect %v==%v,but not",a,b)
		}
	}else {
		if SliceEqual(a,b) {
			t.Errorf("expect %v!=%v,but not",a,b)
		}
	}
}