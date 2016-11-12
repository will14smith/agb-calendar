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
	sort.Sort(ByDateThenName(competitions))

	err = ioutil.WriteFile("output.ical", []byte(format.ConvertToICal(competitions)), 0)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("output.html", []byte(format.ConvertToHtml(competitions)), 0)
	if err != nil {
		panic(err)
	}
}

type ByDateThenName []*model.Competition

func (a ByDateThenName) Len() int      { return len(a) }
func (a ByDateThenName) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByDateThenName) Less(i, j int) bool {
	if a[i].StartDate.Before(a[j].StartDate) {
		return true
	}

	if !a[i].StartDate.Equal(a[j].StartDate) {
		return false
	}

	return a[i].Name < a[j].Name
}
