package function_study

import "strings"

func tt() string {
	s := make([]string, 100)
	for i := 0; i < 100; i++{
		s[i] = "a"
	}
	return strings.Join(s,"xxxx")

}

func String_study() {
	a := tt()
	println(a)
}
