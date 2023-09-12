package templates

import (
	"html/template"
	"log"
	"os"

	"github.com/Artemides/problems/fundamentals/structs"
)

func RunHtml() {
	var issuesList = template.Must(template.New("issues").Parse(`
		<html>
			<body>
				<h1>{{.TotalCount}} issues:</h1>
				<div>
					<table>
						<tr style='text-align: left'>
							<th>#</th>
							<th>State</th>
							<th>User</th>
							<th>Title</th>
						</tr>
						{{range .Items}}
						<tr>
							<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
							<td>{{.State}}</td>
							<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
							<td><a href='.HTMLURL'>{{.Title}}</a></td>
						</tr>
						{{end}}
					</table>
				</div>
			</body>
		</html>
	`))
	response, err := structs.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if err := issuesList.Execute(os.Stdout, response); err != nil {
		log.Fatal(err)
	}
}

func TrustedHTML() {
	var data struct {
		A string
		B template.HTML
	}
	data.A = "<b>Hello A!</b>"
	data.B = "<b>Hello B!</b>"
	const tmpl = `
		<html>
		 <body>
		 <div>
		 	<p>A: {{.A}}</p>
			<p>B: {{.B}}</p>
		 </div>
		 </body>
		</html>
	`
	template := template.Must(template.New("trustedHtml").Parse(tmpl))

	if err := template.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}
