package language_study

func update_value(data[3]int)  {
	//println(&data[0])
	println(data[0], &data[0], data[1], &data[1],data[2],&data[2])
	for i, x := range data { //从data复制品中取值
		if i == 0 {
			data[0] += 100
			data[1] += 100
			data[2] += 100
		}
		println(x, data[i], &data[i])
		//println(x, data[i])
	}
}

func For_study() {
	data := [3]int{200,300,400}
	update_value(data)
	println(data[0], &data[0], data[1], &data[1],data[2],&data[2])
	//for i, x := range data { //从data复制品中取值
	//	if i == 0 {
	//		data[0] += 100
	//		data[1] += 100
	//		data[2] += 100
	//	}
	//	println(x, &data[i])
	//	//println(x, data[i])
	//}
	////println(data[0], data[1], data[2])
	//
	//println()
	//println(&data1[0])
	//for i, x := range data1[:]{ //复制slice，不包括底层array
	//	if i == 0 { //i==0仅修改data1时，x已经取值，所以时10
	//		data1[0] += 100
	//		data1[1] += 100
	//		data1[2] += 100
	//	}
	//	println(x, &data1[i])
	//}
	//println(data1[0], data1[1], data1[2])
}