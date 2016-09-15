package main

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"

	"github.com/will14smith/agb-calendar/format"
	"github.com/will14smith/agb-calendar/model"
)

var dataBase = path.Join("..", "data")
var dayPathRegex = regexp.MustCompile("[0-9]{4}-[0-9]{2}[\\\\/][0-9]{2}.html$")

func main() {
	files := getAllFiles()
	log.Println("Processing", len(files), "files")

	competitions := parseFiles(files)
	log.Println("Parsed", len(competitions), "competitions")
	writeToFile("competitions.ical", competitions)

	competitions = mergeCompetitions(competitions)
	log.Println("Merged into", len(competitions), "competitions")
	writeToFile("merged.ical", competitions)

	// resolve locations (as best as possible)
}

func getAllFiles() []string {
	files, err := filepath.Glob(path.Join(dataBase, "*", "*.html"))
	if err != nil {
		panic(err)
	}

	return files
}

func parseFiles(files []string) []*model.Competition {
	var competitions []*model.Competition

	for _, file := range files {
		if !dayPathRegex.MatchString(file) {
			continue
		}

		// parse each entry
		fileCompetitions, err := parseFile(file)
		if err != nil {
			panic(err)
		}

		competitions = append(competitions, (*fileCompetitions)...)
	}

	return competitions
}

func mergeCompetitions(competitions []*model.Competition) []*model.Competition {
	var merged []*model.Competition
	lookup := make(map[string]*model.Competition, 0)

	for _, competition := range competitions {
		key, found := findCandidate(lookup, competition)

		if found {
			mergedCompetition := mergeCompetition(lookup[key], competition)
			lookup[key] = mergedCompetition

			continue
		}

		if existing, found := lookup[key]; found {
			// flush
			merged = append(merged, existing)
		}

		lookup[key] = competition
	}

	for _, competition := range lookup {
		merged = append(merged, competition)
	}

	return merged
}

func findCandidate(lookup map[string]*model.Competition, competition *model.Competition) (key string, found bool) {
	key = competition.Name

	existing, found := lookup[key]

	if !found {
		return
	}

	// TODO more in depth?
	if competition.StartDate.Sub(existing.EndDate).Hours() <= 24 {
		return
	}

	found = false
	return
}

func mergeCompetition(a, b *model.Competition) *model.Competition {
	return &model.Competition{
		Name: a.Name,

		Location: a.Location,

		StartDate: a.StartDate,
		EndDate:   b.EndDate,

		Rounds: append(a.Rounds, b.Rounds...),

		Organiser: a.Organiser,
		Phone:     a.Phone,
		Email:     a.Email,
		Web:       a.Web,
		Notes:     a.Notes,
	}
}

func writeToFile(fileName string, competitions []*model.Competition) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	_, err = file.WriteString(format.ConvertToICal(competitions))
	if err != nil {
		panic(err)
	}
	err = file.Close()
	if err != nil {
		panic(err)
	}
}
