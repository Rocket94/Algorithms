package util

const INT_MAX =int(^uint32(0)>>1)
const INT_MIN =^INT_MAX

func CompareMax(a,b int)int{
	if a<b{
		return b
	}else {
		return a
	}
}