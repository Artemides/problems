package interfaces

type Table struct {
	Rows           []*Track
	SelectedColumn string
	Ascending      bool
}

func (t Table) Len() int      { return len(t.Rows) }
func (t Table) Swap(i, j int) { t.Rows[i], t.Rows[j] = t.Rows[j], t.Rows[i] }
func (t Table) Less(i, j int) bool {
	fromRow := t.Rows[i]
	toRow := t.Rows[j]

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
