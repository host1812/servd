package generator

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	// "strings"
	// "text/scanner"
	"bufio"
)

func Generate(wfile string, pcount int, depth int) {
	fmt.Println("start generate content")

	os.RemoveAll("pages")

	os.Mkdir("pages", 0764)

	words, _ := readWords(wfile)

	for i := 0; i < pcount; i++ {
		title := words[rand.Intn(len(words)-1)]
		fmt.Println(title)
		filename := "pages/" + title + ".html"
		ioutil.WriteFile(filename, []byte(title), 0600)
	}

	fmt.Println("finish generate content")
}

// readWords reads a whole file into memory
// and returns a slice of its lines.
func readWords(wfile string) ([]string, error) {
	file, err := os.Open(wfile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
