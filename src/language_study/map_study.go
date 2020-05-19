package language_study

import "fmt"
func Map_study() {
	type user struct{ name string }

	m := map[int]struct{
		name string
		age int
	} {
		1:{"user1",10},
		2:{"user2", 10},
	}

	fmt.Println(m[1].name)
	//m[1].name = "xxxx" //map中取出的是value的副本，这样修改变量没有任何意义，会报错
	//正确做法如下：
	//方法一：完全替换value
	u := m[1]
	u.name = "xxx"
	m[1] = u
	fmt.Println(m[1].name)
	//方法二：通过指针的方式
	m2 := map[int]*user{
		1: &user{"user1"},
	}
	m2[1].name = "aaa"
	fmt.Println(m2[1].name)

	m3 := map[string][]string{}
	var s3 []string
	s3 = append(s3, "1")
	s3 = append(s3, "2","3","5")
	m3["1"] = s3
	fmt.Println(m3)

}
