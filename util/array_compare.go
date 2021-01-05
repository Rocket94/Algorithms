package util

import (
	"reflect"
)

type slice interface {
	len   int            //长度有多大
	cap   int            //容量有多大
}


func SliceEqual(a, b slice) bool {
	reflect.SliceOf(a)
	if !CheckSliceType(a, b) {
		return false
	}

	if len(va) != a.len {
		return false
	}
	for i, x := range va {
		if x != vb[i] {
			return false
		}
	}
	return true
}

func CheckSliceType(a, b []interface{}) bool {
	switch a.(type) {
	case nil:
		if b == nil {
			return true
		} else {
			return false
		}
	case []int:
		if _, ok := b.([]int); ok {
			return true
		} else {
			return false
		}
	case []string:
		if _, ok := b.([]string); ok {
			return true
		} else {
			return false
		}
	case []float64:
		if _, ok := b.([]float64); ok {
			return true
		} else {
			return false
		}
	case []byte:
		if _, ok := b.([]byte); ok {
			return true
		} else {
			return false
		}
	case []uint:
		if _, ok := b.([]uint); ok {
			return true
		} else {
			return false
		}
	default:
		return false
	}
}
