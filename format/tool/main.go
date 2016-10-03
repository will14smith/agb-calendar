package main

import (
	"bytes"
	"io/ioutil"
	"sort"

	"github.com/will14smith/agb-calendar/format"
	"github.com/will14smith/agb-calendar/model"
)

func main() {
	filePath := "../../process/tool/output.json"

	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	competitions := format.ConvertFromJson(bytes.NewBuffer(buf).String())
	sort.Sort(ByDate(competitions))

	err = ioutil.WriteFile("output.ical", []byte(format.ConvertToICal(competitions)), 0)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("output.html", []byte(format.ConvertToHtml(competitions)), 0)
	if err != nil {
		panic(err)
	}
}

type ByDate []*model.Competition

func (a ByDate) Len() int           { return len(a) }
func (a ByDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDate) Less(i, j int) bool { return a[i].StartDate.Before(a[j].StartDate) }
