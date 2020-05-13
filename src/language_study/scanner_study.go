package language_study

import (
	"bufio"
	"fmt"
	"os"
)

func Scanner_Study(){
	counts := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() != "exit" {
			counts[input.Text()] ++
		} else {
			break
		}
	}

	for key, value := range counts {
		fmt.Println("key:",key, "value:", value)
	}
}
