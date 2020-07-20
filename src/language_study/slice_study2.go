package language_study

import "fmt"

func SliceStudy2() {
	arr := []int{22,33,44,55,66,88,77}
	s := []*int{}
	for _, v := range arr {
		s = append(s, &v)
	}

	for _, v := range s{
		fmt.Println(*v)
	}
}
