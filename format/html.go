package format

import (
	"bytes"
	"html/template"

	"github.com/will14smith/agb-calendar/model"
)

const pageTemplate = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>AGB Competitions</title>
	</head>
	<body>
		{{range .}}
            <h1>{{ .Name }}</h1>
            <dl>
                <dt>Location</dt><dd>{{ .Location.Name }}</dd>
                <dt>Round</dt><dd><ul>
                    {{range .Rounds}}
                        <li>{{ . }}</li>
                    {{else}}<li><strong>No rows</strong></li>{{end}}
                </ul></dd>
                {{if eq .StartDate.Unix .EndDate.Unix}}
                <dt>When</dt><dd>{{template "date" .StartDate }}</dd>
                {{else}}
                <dt>When</dt><dd>{{template "date" .StartDate }} - {{template "date" .EndDate }}</dd>
                {{end}}
                {{if .DrivingDirections}}
                    <dt>Driving Time</dt><dd>{{ .DrivingDirections.Duration }}</dd>
                {{end}}
                {{if .PublicDirections}}
                    <dt>Transit Time</dt><dd>{{ .PublicDirections.Duration }}</dd>
                {{end}}
            </dl>
        {{else}}<div><strong>No rows</strong></div>{{end}}
	</body>
</html>
`
const dateTemplate = "{{define `date`}}{{ .Format `2006-01-02` }}{{end}}"

func ConvertToHtml(competitions []*model.Competition) string {
	templates, err := template.New("page").Parse(pageTemplate)
	if err != nil {
		panic(err)
	}

	_, err = templates.Parse(dateTemplate)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	templates.ExecuteTemplate(&buf, "page", competitions)

	return buf.String()
}
