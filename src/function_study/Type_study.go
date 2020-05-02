package function_study

import "fmt"

type Attr struct {
	pre int
	sec int
}

type File struct {
	name string
	size int
	Attr
}

type Anim struct {
	id int
	name string
}

type Dog struct {
	Anim
	eat string
}

func (self *Anim) ToString() string {
	return fmt.Sprintf("Anim ToString %p  , %v", self, self)
}

func (self *Dog) ToString() string {
	return fmt.Sprintf("Dog ToString %p  , %v", self, self)
}


func Type_study(){
	//f := File{
	//	name: "xxx",
	//	size: 1024,
	//	Attr: Attr{1,2},
	//}
	//fmt.Println(f)

	/*通过匿名字段，获得了和继承蕾丝的复用能力*/
	dog := Dog{Anim {1,"xiaohuang"}, "gotou"}
	fmt.Println(dog.ToString())
	fmt.Println(dog.Anim.ToString())

}
