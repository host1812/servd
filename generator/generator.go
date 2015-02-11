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

func Generate(wfile string, pcount int, wcount int) {
	fmt.Println("start generate content")

	os.RemoveAll("pages")

	os.Mkdir("pages", 0764)

	words, _ := readWords(wfile)
	indexBody := ""

	for p := 0; p < pcount; p++ {
		title := words[rand.Intn(len(words)-1)]
		fmt.Println(title)
		// filename := "pages/" + title + ".html"
		filename := "pages/" + title

		// add generated page to 'index' page
		indexBody += "<a href='/view/" + filename + "'>" + filename + "</a><br />\n"

		pContent := "surovo <br/> "
		for w := 0; w < wcount; w++ {
			pContent += words[rand.Intn(len(words)-1)] + " "
		}

		ioutil.WriteFile(filename, []byte(pContent), 0600)
	}

	ioutil.WriteFile("./index", []byte(indexBody), 0600)

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
