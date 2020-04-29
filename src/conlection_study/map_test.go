package conlection_study

import "fmt"

func map_test() {
	i := map[int]string{0: "a", 1: "a", 2: "a", 3: "a", 4: "a", 5: "a", 6: "a", 7: "a", 8: "a", 9: "a"}

	for k:= range i{
		i[k+k] = "x"
		delete(i,k)
	}
	/*
		由于map是散列，在map迭代时是按hash顺序取值的，而非添加顺序，所以在迭代时插入数据又删除数据，可能会引发意想不到的问题
	*/
	fmt.Println(i)
}
