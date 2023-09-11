package templates

import (
	"html/template"
	"log"
	"os"
	"time"

	"github.com/Artemides/problems/fundamentals/structs"
)

func Run() {
	const templ = `{{.TotalCount}} issues:
	 {{range .Items}}
	 -------------------
	 Number:	{{.Number}}
	 User:		{{.User.Login}}
	 Title:		{{.Title | printf "%.64s"}}
	 Age: 		{{.CreatedAt | daysAgo}} days
	 {{end}}`

	report := template.Must(template.New("report").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(templ))
	response, err := structs.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if err := report.Execute(os.Stdout, response); err != nil {
		log.Fatal(err)
	}
}

func daysAgo(date time.Time) int {
	return int(time.Since(date).Hours() / 24)
}
