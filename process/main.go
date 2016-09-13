package main

import (
	"fmt"
	"path"
	"path/filepath"
	"regexp"
)

var dataBase = path.Join("..", "data")
var dayPathRegex = regexp.MustCompile("[0-9]{4}-[0-9]{2}[\\\\/][0-9]{2}.html$")

func main() {
	// iterate days in data folder
	files, err := filepath.Glob(path.Join(dataBase, "*", "*.html"))
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if !dayPathRegex.MatchString(file) {
			continue
		}

		// parse each entry
		fileCompetitions, err := parseFile(file)
		if err != nil {
			panic(err)
		}

		fmt.Println(fileCompetitions)
	}

	// merge entries
	// resolve locations (as best as possible)
}
