package language_study

func Switch_stude() {
	switch x := 5; x{
	case 5:
		x+=10
		println(x)
		//fallthrough //执行下一个case，但不再匹配表达式
	case 10:
		println(10)
		fallthrough
	case 15:
		println(x)
	}
}
