package function_study

import "fmt"

type Tester struct {
	s interface {
		String() string
	}
}

type Stringer interface {
	String() string
}

type Printer interface {
	Stringer
	Print()
}

type UserIn struct {
	id int
	name string
}

func (self *UserIn) String() string {
	return fmt.Sprintf("user string %d, %s", self.id, self.name)
}

func (self *UserIn) Print() {
	fmt.Println("user print ", self.id, self.name)
}

func Interfase_Study () {
	//t := Tester{&UserIn{1, "Tom"}}
	//fmt.Println(t.s.String())
	var a Printer = &UserIn{1,"aaa"}
	a.Print()
	fmt.Println(a.String())
}
