package algorithms

func BF(main, model string) []int {
	res := []int{}
	if model == "" || main == "" || len(main) < len(model) {
		return res
	}
	b_main := []byte(main)
	b_model := []byte(model)

	for i := 0; i <= len(b_main)-len(b_model); i++ {
		for j := 0; j < len(b_model); j++ {
			if b_main[i+j] == b_model[j] {
				if j == len(b_model)-1 {
					res = append(res, i)
				}
			} else {
				break
			}
		}
	}
	return res
}

var mk = []int64{1, 26, 676, 17576, 456976, 11881376, 308915776, 8031810176, 208827064576, 5429503678976, 141167095653376}

//利用hash值传递，可以避免一次次的计算整个长度为m的hash值，效率高于BF
func RK(main, model string) []int {
	res := []int{}
	if model == "" || main == "" || len(main) < len(model) {
		return res
	}

	m := len(model)
	n := len(main)

	model_hash := StringHash([]byte(model))
	main_b:=[]byte(main)

	//只用计算这一次
	main_hash := StringHash([]byte(main)[0:m])
	if main_hash == model_hash {
		res = append(res, 0)
	}
	for i := 1; i <= n-m; i++ {
		//计算hash值，利用前一个结果可以简单计算，不需要调用StringHash函数
		main_hash = (main_hash - mk[m]*int64(main_b[i-1])) + int64(main_b[i+m-1])
		if main_hash == model_hash {
			res = append(res, i)
		}
	}
	return res
}

func StringHash(b []byte) int64 {
	if len(b) > 10 {
		return -1
	}
	var res int64 = 0

	for i := 0; i < len(b); i++ {
		res += int64(b[i]) * mk[len(b)-1-i]
	}
	return res
}
