package file_study

import (
	"fmt"
	"os"
)

func File_os(){
	fmt.Println(os.UserHomeDir()) //获取host path

	//res := os.Mkdir("/Users/muyou/ttt/ff", 777) // 创建文件
	res := os.Chdir("/Users/muyou/gostudy/golang_learning/src")
	fmt.Println(res)

}
