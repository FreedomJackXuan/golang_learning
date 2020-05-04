package function_study

//func ca(){
//	log.Println("ca", recover())
//}

// recover必须在延迟调用函数中执行才能正常工作
func ca(x, y int) {
	z:=0

	func() {
		defer func () {
			if recover() != nil {
				z = 0
			}
		}()

		// defer recover()  捕获painc失败
		z = x/y
	}()

	println("x / y =",z)
}

func Panic_study() {
	//defer ca()
	//defer log.Println(recover())
	//defer func() {
	//	if err:=recover(); err!=nil {
	//		log.Println(err)
	//	} else {
	//		log.Println("xxx")
	//	}
	//}()
	//panic("111")
	ca(5,0)
}
