package function_study

func test(a *int){
	*a = 100
	println(&a, a)
}

func GoFuc1() {
	var a int = 10
	p := &a
	println(&p, p)
	test(p)
	println(a)
}
