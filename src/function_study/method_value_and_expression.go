package function_study

import "fmt"

/*
方法绑定对象实例时，会隐式将实例作为第一实参(receiver)， 可以是 T 或 *T
*/

type AA struct {
	name string
	age int
}

func (self AA) Test1(){
	fmt.Printf("Test1 %p，%v\n",&self,self)
}

func (self *AA) Test2(){
	fmt.Printf("Test2 %p，%v\n", self, self)
}

func Method_1(){
	a := AA{name: "xxx", age:123}
	fmt.Printf("AA %p，%v\n",&a,a)
	mValue1 := a.Test1 //立即复制 receiver，因为不是指针类型，不受后续修改影响。
	mExpression1 := (*AA).Test1

	mValue2 := a.Test2
	mExpression2 := (*AA).Test2
	a.name, a.age = "sss", 2

	// 可以将Test1设想为  func Test1(self AA){}  可以看出传入的是值类型，会copy一份作为参数传入
	a.Test1()
	mValue1()
	mExpression1(&a)

	// 可以将Test2设想为  func Test2(self *AA){}  可以看出传入的是一个指针
	a.Test2()
	mValue2()
	mExpression2(&a)

}