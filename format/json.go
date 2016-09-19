package format

import (
	"bytes"
	"encoding/json"

	"github.com/will14smith/agb-calendar/model"
)

func ConvertToJson(competitions []*model.Competition) string {
	buf, err := json.MarshalIndent(competitions, "", "  ")
	if err != nil {
		panic(err)
	}

	return bytes.NewBuffer(buf).String()
}

func ConvertFromJson(str string) []*model.Competition {
	var competitions []*model.Competition
	err := json.Unmarshal([]byte(str), &competitions)
	if err != nil {
		panic(err)
	}

	return competitions
}
