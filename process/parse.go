package process

import (
	"fmt"
	"os"
	"strings"
	"time"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"

	"github.com/will14smith/agb-calendar/model"
	"github.com/yhat/scrape"
)

func ParseFile(path string) (*[]*model.Competition, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	node, err := html.Parse(file)
	if err != nil {
		return nil, err
	}

	eventTable, found := scrape.Find(node, scrape.ByClass("eventstable"))
	if !found {
		return nil, fmt.Errorf("Couldn't file .eventstable")
	}

	dateNode, _ := scrape.Find(node, scrape.ByClass("selecteddate"))
	date, _ := time.Parse("2 January 2006", scrape.Text(dateNode))

	// table > tbody > tr
	currentRow := eventTable.FirstChild.FirstChild

	var competitions []*model.Competition
	for currentRow != nil {
		headerRow := currentRow
		infoRow := currentRow.NextSibling

		nameNode, _ := scrape.Find(headerRow, scrape.ByClass("title"))
		name := scrape.Text(nameNode)

		locationNode, _ := scrape.Find(headerRow, scrape.ByClass("where"))
		location := scrape.Text(locationNode)

		descNode, _ := scrape.Find(infoRow, scrape.ByClass("desc"))
		descLines := extractTextFromDescription(descNode)
		descMap := extractMapFromDescriptionLines(descLines)

		round, _ := descMap["round"]
		organiser, _ := descMap["organiser"]
		phone, _ := descMap["phone"]
		email, _ := descMap["email"]
		web, _ := descMap["web"]
		notes, _ := descMap["notes"]

		competitions = append(competitions, &model.Competition{
			Name:     name,
			Location: &model.Location{Name: location},

			// start and end are equal in this step
			StartDate: date,
			EndDate:   date,

			Rounds: []string{round},

			Organiser: organiser,
			Phone:     phone,
			Email:     email,
			Web:       web,
			Notes:     notes,
		})

		currentRow = infoRow.NextSibling
	}

	return &competitions, nil
}

func extractTextFromDescription(root *html.Node) []string {
	var lines []string

	var currentLine string
	for node := root.FirstChild; node != nil; node = node.NextSibling {
		switch node.Type {
		case html.ElementNode:
			switch node.DataAtom {
			case atom.Span:
				currentLine = currentLine + scrape.Text(node)
			case atom.Br:
				lines = append(lines, currentLine)
				currentLine = ""
			default:
				panic(fmt.Errorf("Unexpected element tag: %v", node.DataAtom))
			}
		case html.TextNode:
			currentLine = currentLine + scrape.Text(node)
		default:
			panic(fmt.Errorf("Unexpected node type: %v", node.Type))
		}
	}

	if currentLine != "" {
		lines = append(lines, currentLine)
	}

	return lines
}

func extractMapFromDescriptionLines(lines []string) map[string]string {
	result := make(map[string]string, 0)

	for _, line := range lines {
		segments := strings.SplitN(line, ":", 2)

		key := strings.ToLower(segments[0])
		value := strings.TrimSpace(segments[1])

		result[key] = value
	}

	return result
}
