package file_study

import (
	"fmt"
	"os"
)

func FileTree(){
	paths := os.Args[1:]
	if len(paths) > 1 {
		return
	}

	fmt.Println(os.UserHomeDir())

	//if err != nil {
	//	fmt.Println(err)
	//}

}
