package language_study

import (
	"fmt"
)

/*
Slice 底层结构
struct Slice{
	btye* array;
	uintgo len;
	uintgo cap;
}
*/

func appendIntTest(x []int, y int) []int {
	var tmp []int
	tlen := len(x) + 1
	if tlen <= cap(x) {
		tmp = x[:tlen]
	} else {
		tcap := tlen
		if tcap < 2 * len(x) {
			tcap = 2 * len(x)
		}
		tmp = make([]int, tlen, tcap)
		copy(tmp, x)
	}
	tmp[len(x)] = y
	return tmp
}


func Slice_study() {
	data := [...]int{0,1,2,3,4,5,6,7,8,9}
	s1 := data[1:5:8] // [low:high:max]  len = high - low;   cap = max - low
	fmt.Println(s1,len(s1),cap(s1))

	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}

	var s3 []rune
	for _, r1 := range "123"{
		s3 = append(s3, r1)
	}

	fmt.Printf("%q\n", runes)
	fmt.Printf("%q\n",s3)

	s5 := []int{1,2,3,4,5}
	s5 = appendIntTest(s5,10)
	fmt.Println(s5[:])


	s := make([]int,0, 5)
	n := 0
	s = s[:n+1]
	s[n] = 1
	fmt.Println(s)
	n = 2
	s = s[:n+1]
	s[n] = 2
	fmt.Println(s)
	l := len(s)
	p := s[l - 1]
	s = s[: n-1]
	fmt.Println(l, p, s)
}
