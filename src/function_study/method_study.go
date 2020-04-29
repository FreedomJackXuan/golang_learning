package function_study

import "fmt"

type User struct {
	name string
	age int
}

func (self *User) PrintUser() {
	fmt.Printf("%p, %v\n",self,self)
}

type User1 struct {
	name string
	age int
}

func (self User1) PrintUser1() {
	fmt.Println(self)
}

func Method_study() {
	u:=User{"111",2}
	u.PrintUser()
	mTest := u.PrintUser
	mDisply:=(*User).PrintUser
	u.name="xxx"
	u.age=40
	mTest()  //隐式传递
	mDisply(&u)  //显示传递

	fmt.Println("================================")

	u1:=User1{"111",2}
	u1.PrintUser1()
	mTest1 := u1.PrintUser1
	mDisply1:=(*User1).PrintUser1
	u1.name="xxx"
	u1.age=40
	mTest1()  //隐式传递 由于不是指针类型，对修改的值不会修改，而是copy之前的
	mDisply1(&u1)  //显示传递
}
