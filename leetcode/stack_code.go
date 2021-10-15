package leetcode

import (
	"Algorithms/datastruct"
	"Algorithms/util"
	"fmt"
)

//利用栈实现表达式求值（计算器）
//
// 1.不带括号，符号有优先级
//方法分两个栈，一个存储数字，一个存储符号,遇见数字直接压入
//遇见符号，去和栈顶元素比较，如果栈顶元素优先级高或者相同，则取2个数字和栈顶计算符号计算，计算结果压入数字栈
//如果优先级比栈顶元素高那么就直接入符号栈
//
// 2. 带括号，符号没有优先级，如下，实现和上面差不多，括号判断比较复杂，考虑的情况比较多
//func Calculate(s string) int {
//	var bs = []byte(s)
//	var num_stack = new(datastruct.Stack)
//	var symbol_stack = new(datastruct.Stack)
//
//	for i := 0; i < len(bs); i++ {
//		btostring := string(bs[i])
//		res, ok := util.IsDigital(btostring)
//		if ok {
//			for j := i + 1; j < len(bs); j++ {
//				if res2, ok := util.IsDigital(string(bs[j])); ok {
//					res = res*10 + res2
//					i++
//				} else {
//					break
//				}
//			}
//			num_stack.Push(res)
//		} else if btostring == "+" || btostring == "-" {
//			sout, ok := symbol_stack.Pop()
//			if ok {
//				if sout == "+" {
//					n1, _ := num_stack.Pop()
//					n2, _ := num_stack.Pop()
//					n1int, _ := n1.(int)
//					n2int, _ := n2.(int)
//					num_stack.Push(n1int + n2int)
//					symbol_stack.Push(btostring)
//				} else if sout == "-" {
//					n1, _ := num_stack.Pop()
//					n2, _ := num_stack.Pop()
//					n1int, _ := n1.(int)
//					n2int, _ := n2.(int)
//					num_stack.Push(n2int - n1int)
//					symbol_stack.Push(btostring)
//				} else if sout == "(" {
//					symbol_stack.Push(sout)
//					symbol_stack.Push(btostring)
//				}
//			} else {
//				symbol_stack.Push(btostring)
//			}
//		} else if btostring == "(" {
//			symbol_stack.Push(btostring)
//			if string(bs[i+1]) == "-" {
//				num_stack.Push(0)
//			}
//		} else if btostring == ")" {
//			symbol, _ := symbol_stack.Pop()
//			if symbol != "(" {
//				n1, _ := num_stack.Pop()
//				n2, _ := num_stack.Pop()
//				n1int, _ := n1.(int)
//				n2int, _ := n2.(int)
//				if symbol == "+" {
//					num_stack.Push(n1int + n2int)
//				} else if symbol == "-" {
//					num_stack.Push(n2int - n1int)
//				}
//				symbol, _ = symbol_stack.Pop()
//				if symbol != "(" {
//					return -1
//				}
//			}
//		}
//	}
//	symbol_last, _ := symbol_stack.Pop()
//	if symbol_last == "+" {
//		n1, _ := num_stack.Pop()
//		n2, _ := num_stack.Pop()
//		n1int, _ := n1.(int)
//		n2int, _ := n2.(int)
//		num_stack.Push(n1int + n2int)
//	}
//	if symbol_last == "-" {
//		n1, _ := num_stack.Pop()
//		n1int, _ := n1.(int)
//		n2, ok := num_stack.Pop()
//		if !ok {
//			return n1int * (-1)
//		}
//		n2int, _ := n2.(int)
//		num_stack.Push(n2int - n1int)
//	}
//	n, _ := num_stack.Pop()
//	nint, _ := n.(int)
//	return nint
//}
// 3.该种方法比较简单，因为没有运算符号优先级，那么从左到右计算即可
//遇见括号需要把括号外面的+/-号记录，如果加号则括号里面符号不变号，如果减号则编号，因为有多个括号包含的情况，则需要用到栈来解决
//最里面的括号最先求值，所以最里面的括号的符号应该是后进先出，遇见'('则将前面符号压入栈;遇见')'则将括号弹出栈，以弃用
func Calculate2(s string) int {
	var bs = []byte(s)
	fmt.Println(bs)
	var opt_stack=new(datastruct.Stack)
	sign:=1
	opt_stack.Push(sign)
	i:=0
	res:=0

	for i<len(bs){
		if string(bs[i])=="("{
			opt_stack.Push(sign)
		}
		if string(bs[i])=="+"{
			sign_temp,_:=opt_stack.Head.Data.(int)
			sign=sign_temp
		}
		if string(bs[i])=="-"{
			sign_temp,_:=opt_stack.Head.Data.(int)
			sign=sign_temp*(-1)
		}
		if string(bs[i])==")"{
			opt_stack.Pop()
		}
		if num,ok:=util.IsDigital(string(bs[i]));ok{
			for j:=i+1;j<len(bs);j++{
				if num_followed,ok:=util.IsDigital(string(bs[j]));ok{
					num=num*10+num_followed
					i++
				}else {
					break
				}
			}
			res+=num*sign
		}
		i++
	}

	return res
}

//利用栈实现括号匹配,利用栈的特性来处理
func IsValid(s string) bool {
	var bs = []byte(s)
	var stack = new(datastruct.Stack)
	for _, b := range bs {
		temp := string(b)
		if temp == "(" || temp == "{" || temp == "[" {
			stack.Push(temp)
		} else if temp == ")" {
			other, ok := stack.Pop()
			if !ok || other != "(" {
				return false
			}
		} else if temp == "}" {
			other, ok := stack.Pop()
			if !ok || other != "{" {
				return false
			}
		} else if temp == "]" {
			other, ok := stack.Pop()
			if !ok || other != "[" {
				return false
			}
		} else {
			return false
		}
	}
	if _, ok := stack.Pop(); !ok {
		return true
	}
	return false
}
