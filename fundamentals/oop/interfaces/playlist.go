package interfaces

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	duration, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return duration

}

func printTracks(tracks []*Track) {
	format := "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "------", "------", "------", "------", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}

	tw.Flush()
}

type byArtist []*Track

func (a byArtist) Len() int           { return len(a) }
func (a byArtist) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byArtist) Less(i, j int) bool { return a[i].Year < a[j].Year }

func Playlist() {
	// sort.Sort(byArtist(tracks))
	sort.Sort(sort.Reverse(byArtist(tracks)))
	printTracks(tracks)
}
