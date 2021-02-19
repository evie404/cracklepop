package main

import (
	"fmt"
	"strconv"
)

func cracklePop(n int) string {
	if n%15 == 0 {
		return "CracklePop"
	}

	if n%5 == 0 {
		return "Pop"
	}

	if n%3 == 0 {
		return "Crackle"
	}

	return strconv.Itoa(n)
}

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(cracklePop(i))
	}
}
