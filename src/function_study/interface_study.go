package function_study

import "fmt"

type Tester struct {
	s interface {
		String() string
	}
}

type UserIn struct {
	id int
	name string
}

func (self *UserIn) String() string {
	return fmt.Sprintf("user %d, %s", self.id, self.name)
}

func Interfase_Study () {
	t := Tester{&UserIn{1, "Tom"}}
	fmt.Println(t.s.String())
}
//15387741582
