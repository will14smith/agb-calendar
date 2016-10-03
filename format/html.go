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
        <style>
        /* I'm so sorry... */
        * { box-sizing: border-box; font-family: sans-serif; }
        dd { margin: 0; }
        .title { width: 100px; font-weight: bold; }
        .body { margin-bottom: 10px; }
        .location, .when, .driving, .transit { float: left; margin-right: 50px; }
        h1, .rounds.title, .driving.title { clear: both; }
        ul { margin: 0; padding: 0; margin-left: 165px; }
        </style>
	</head>
	<body>
		{{range .}}
            <h1>{{ .Name }}</h1>
            <dl>
                <dt class="location title">Location</dt><dd class="location data">{{ .Location.Name }}</dd>
                <dt class="rounds title">Round</dt><dd class="rounds body"><ul>
                    {{range .Rounds}}
                        <li>{{ . }}</li>
                    {{else}}<li><strong>No rows</strong></li>{{end}}
                </ul></dd>
                {{if eq .StartDate.Unix .EndDate.Unix}}
                <dt class="when title">When</dt><dd class="when body">{{template "date" .StartDate }}</dd>
                {{else}}
                <dt class="when title">When</dt><dd class="when body">{{template "date" .StartDate }} - {{template "date" .EndDate }}</dd>
                {{end}}
                {{if .DrivingDirections}}
                    <dt class="driving title">Driving Time</dt><dd class="driving body">{{ .DrivingDirections.Duration }}</dd>
                {{end}}
                {{if .PublicDirections}}
                    <dt class="transit title">Transit Time</dt><dd class="transit body">{{ .PublicDirections.Duration }}</dd>
                {{end}}
            </dl>@
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
