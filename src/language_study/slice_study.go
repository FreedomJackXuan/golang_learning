package language_study

import "fmt"

/*
Slice 底层结构
struct Slice{
	btye* array;
	uintgo len;
	uintgo cap;
}
*/


func Slice_study() {
	data := [...]int{0,1,2,3,4,5,6,7,8,9}
	s1 := data[1:5:8] // [low:high:max]  len = high - low;   cap = max - low
	fmt.Println(s1,len(s1),cap(s1))
}
