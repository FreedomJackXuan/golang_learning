package function_study

import "log"

func ca(){
	log.Println("ca", recover())
}

func Panic_study() {
	//defer ca()
	//defer log.Println(recover())
	defer func() {
		if err:=recover(); err!=nil {
			log.Println(err)
		} else {
			log.Println("xxx")
		}
	}()
	panic("111")
}
