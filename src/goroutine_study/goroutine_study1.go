package goroutine_study

import "sync"

func sum(id int) {
	var x int64
	for i := 0; i < 10000; i++ {
		x += int64(i)
	}
	println(id, x)
}

func Goroutine_Study1() {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func(id int) {
			defer wg.Done()
			sum(i)
		}(i)
	}
	wg.Wait()
}