package string_study

import (
	"strings"
	"testing"
)

func test1() string {
	var s string
	for i:=0; i < 1000; i++{
		s +="a"
	}
	return s
}

func test2() string {
	s := make([]string, 1000)
	for i :=0; i < 1000; i++{
		s[i] = "a"
	}
	return strings.Join(s, "")
}

func BenchmarkString_Study(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//test1() //BenchmarkString_Study-4   	    5182	    255573 ns/op
		//test2() //BenchmarkString_Study-4   	   86475	     12054 ns/op
	}
}