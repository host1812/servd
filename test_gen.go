package main

import (
	"servd/generator"
)

func main() {
	generator.Generate("/usr/share/dict/words", 50, 0)
}
