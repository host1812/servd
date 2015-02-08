package generator

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Generate(wfile string, pcount int, depth int) {
	fmt.Println("start generate content")

	c, err := ioutil.ReadFile(wfile)
	if err != nil {
		fmt.Println("error reading dictionary file")
		return
	}
	words := strings.Split(string(c), " ")

	for _, w := range words {
		fmt.Println(w)
	}

	fmt.Println("finish generate content")
}
