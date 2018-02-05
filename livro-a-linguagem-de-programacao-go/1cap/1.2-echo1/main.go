//Echo 1 exibe seus argumentos de linha de comando
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	itemSlice()
}

func itemSlice() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
