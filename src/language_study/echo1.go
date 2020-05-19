package language_study

import (
	"fmt"
	"os"
)

type celi int64
type beli int64

func Echo1() {

	var i celi = 1
	b := beli(i)

	fmt.Println(b, i)
	var s, sep string
	for i:=1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
