package main

import (
	"fmt"
	"os"
	"servd/generator"
	"strconv"
)

func main() {
	args := os.Args[1:]
	pcount := 50
	wcount := 50

	if len(args) >= 1 {
		pcount, _ = strconv.Atoi(args[0])
	}

	if len(args) >= 2 {
		wcount, _ = strconv.Atoi(args[1])
	}

	fmt.Println("pcount = ", pcount)
	fmt.Println("wcount = ", wcount)
	generator.Generate("/usr/share/dict/words", pcount, wcount)
}
