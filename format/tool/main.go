package main

import (
	"bytes"
	"io/ioutil"

	"github.com/will14smith/agb-calendar/format"
)

func main() {
	filePath := "../../process/tool/output.json"

	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	competitions := format.ConvertFromJson(bytes.NewBuffer(buf).String())

	err = ioutil.WriteFile("output.ical", []byte(format.ConvertToICal(competitions)), 0)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("output.html", []byte(format.ConvertToHtml(competitions)), 0)
	if err != nil {
		panic(err)
	}
}
