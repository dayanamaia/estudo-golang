// echo2 exibe seus argumentos de linha de comando
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println("comando executado:", s)
}
