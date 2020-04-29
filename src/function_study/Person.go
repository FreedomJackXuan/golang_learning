package function_study

import "fmt"

type Person struct {
	name string
	age int
}

func (self *Person) SetName (name string) {
	self.name = name
}

func (self *Person) GetName() string {
	return self.name
}

func (self Person) SetAge (age int) {
	self.age = age
}

func (self Person) GetAge() int {
	return self.age
}

func (self *Person) SetPointerAge (age int) {
	self.age = age
}

func (self *Person) GetPointerAge() int {
	return self.age
}

func PersonTest() {
	p := new(Person)
	p.SetName("xxx")
	p.SetAge(10) //由于这个函数是值类型，会copy一份参数传到setage函数中，赋值完之后栈就会释放，在getAge的时候获取到的值为0
	fmt.Println(p.GetName())
	fmt.Println(p.GetAge())
	p.SetPointerAge(100)
	fmt.Println(p.GetPointerAge())
}
