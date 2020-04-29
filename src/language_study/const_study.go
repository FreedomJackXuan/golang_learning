package language_study

import (
	"strconv"
	"unsafe"
)

func Const_test() {
	const xxx = 123
	println(xxx)
	const y = 1.23
	{
		const xxx = "aaa"
		println(xxx)

	}

	const ptrSize = unsafe.Sizeof(uintptr(0))
	const strSize = len("hello world")
	println(ptrSize)
	println(strSize)

	//const (
	//	a = iota
	//	b
	//	c
	//)
	//println(b)

	//const (
	//	a = iota
	//	b
	//	c = 100
	//	d = 100
	//	e = iota //恢复iota自增，包括c,d
	//	f
	//)
	//println(a,b,c,d,e,f)

	a, _ := strconv.ParseInt("1100100", 2, 64)
	println(a)
}
