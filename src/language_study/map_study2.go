package language_study

import "fmt"

func MapStudy2() {
	m := map[string][2]int{
		"a":{1,2},
	}
	//s:= m["a"][:] //数组必须是addressable
	fmt.Println(m)
	/*
	go 中对addressable做了详细的规定：
	对于一个对象x，如果它的类型为T，那么&x会产生一个*T的指针，这个指针只想x
	x必须为可寻址的，也就是一下几种方式：
	1.一个变量 &x
	2.指针引用 &*x
	3.slice索引操作 &s[1]：不论slice是否可寻址，其元素值是可以寻址的，因为slice底层是数组，是可以寻址的
	4.可寻址struct的字段 &point.X
	5.可寻址数组的索引操作 &a[0]
	6.composite literal类型 &struct{x int }{1}

	以下情况是不可寻址的，不能使用&x获取指针：
	1.字符串中的字节：字符串不可改变
	2.map对象中的元素：如果对象不存在，则返回零值，零值是不可变对象，不能寻址；如果对象存在，因为go中map实现中元素的地址是变化的，这意味着寻址的结果没有意义
	3.接口对象的动态值(通过type assertions获得)
	4.常数：如果常数可以寻址，那就可以通过指针修改常数值，破坏常数的定义
	5.Literal值
	6.package级别的函数
	7.方法method
	8.中间值（函数调用、显示类型转换、各种类型的操作(出了指针引用*x)）

	规范中还有几处提到了addressable
	1.调用一个receiver为指针类型的方法时，使用一个addressable的值将会自动获取这个值的指针
	2.++，--语句的操作对象必须是addressable或者是map的index操作
	3.赋值语句=的左边对象必须是addressable，或者是map的index操作，或者是_

	*/
}
