package interfaces

import (
	"fmt"
	"html/template"
	"net/http"
	"sort"
)

type Table struct {
	Tracks         []*Track
	SelectedColumn string
}

func (t Table) Len() int      { return len(t.Tracks) }
func (t Table) Swap(i, j int) { t.Tracks[i], t.Tracks[j] = t.Tracks[j], t.Tracks[i] }
func (t Table) Less(i, j int) bool {
	fromRow := t.Tracks[i]
	toRow := t.Tracks[j]

	if t.SelectedColumn == "Title" && fromRow.Title != toRow.Title {
		return fromRow.Title < toRow.Title
	}
	if t.SelectedColumn == "Artist" && fromRow.Artist != toRow.Artist {
		return fromRow.Artist < toRow.Artist
	}
	if t.SelectedColumn == "Album" && fromRow.Album != toRow.Album {
		return fromRow.Album < toRow.Album
	}
	if t.SelectedColumn == "Year" && fromRow.Year != toRow.Year {
		return fromRow.Year < toRow.Year
	}
	if t.SelectedColumn == "Length" && fromRow.Length != toRow.Length {
		return fromRow.Length < toRow.Length
	}
	return false
}

func tableTemplate() *template.Template {
	return template.Must(template.New("playlist").Parse(`
	<!DOCTYPE html>
	<html lang="en">
	  <head>
		<meta charset="UTF-8" />
		<meta name="viewport" content="width=device-width, initial-scale=1.0" />
		<title>Go Playlist</title>
		<link rel="stylesheet" type="text/css" href="/static/playlist.css">
	  </head>
	  <body>
	 	<main>
			<div class='table-container'>
				<table>
					<tr class='table-header'>
						<th id="sort-title" data-sortBy="Title" class="header-btn"> Title </th>
						<th id="sort-artist" data-sortBy="Artist" class="header-btn"> Artist </th>
						<th id='sort-album' data-sortBy="Album" class="header-btn"> Album </th>
						<th id='sort-year' data-sortBy="Year" class="header-btn"> Year </th>
						<th id='sort-length' data-sortBy="Length" class="header-btn"> Length </th>
					</tr>
					{{range .Tracks}}
					<tr>
						<td>{{.Title}}</td>
						<td>{{.Artist}}</td>
						<td>{{.Album}}</a></td>
						<td>{{.Year}}</a></td>
						<td>{{.Length}}</a></td>
					</tr>
					{{end}}
				</table>
			</div>
			<script src="/static/playlist.js"></script>
		</main> 
	  </body>
	</html>
	`))
}

func TableHndler(w http.ResponseWriter, r *http.Request) {
	sortBy := ""
	playlistTable := tableTemplate()
	params := r.URL.Query()
	sortBy = params.Get("sortBy")
	data := Table{Tracks, sortBy}
	sort.Sort(data)
	if err := playlistTable.Execute(w, data); err != nil {
		fmt.Fprintf(w, "404 page not found: %s", err.Error())
		return
	}
}
