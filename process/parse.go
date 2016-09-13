package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"

	"github.com/will14smith/agb-calendar/model"
	"github.com/yhat/scrape"
)

func parseFile(path string) (*[]model.Competition, error) {
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

	// table > tbody > tr
	currentRow := eventTable.FirstChild.FirstChild

	var competitions []model.Competition
	for currentRow != nil {
		headerRow := currentRow
		infoRow := currentRow.NextSibling

		nameNode, _ := scrape.Find(headerRow, scrape.ByClass("title"))
		name := scrape.Text(nameNode)

		competitions = append(competitions, model.Competition{
			Name: name,
		})

		currentRow = infoRow.NextSibling
	}

	return &competitions, nil
}
