package main

import (
	"ch4/github"
	"golang.org/x/text/search"
	"html/template"
	"log"
	"os"
	"time"
)

const templ = `{{.TotalCount}} issues:
    {{range .Items}}
    Number: {{.Number}}
    User: {{.User.Login}}
    Title: {{.Title | printF "%.64s"}}
    Age: {{.CreatedAt | daysAgo}} days
    {{end}}`

func daysAgo(t time.Time) int {
	return int (time.Since(t).Hours() / 24)
}


var report = template.Must(template.New("issuelist").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(templ))

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
