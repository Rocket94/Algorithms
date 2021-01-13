package main

import (
	"bytes"
	"io"
)

const debug = false

func main() {
	var buf *bytes.Buffer
	if debug {
		buf = new(bytes.Buffer)
	}
	f(buf)
}
//调用f的时候给out赋了一个nil的bytes.Buffer，虽然out的动态值是nil，但是其动态类型是bytes.Buffer这样判断为true
func f(out io.Writer) {
	if out != nil {
		out.Write([]byte("hello"))
	}
}