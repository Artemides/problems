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

// Create a slice of 10 tracks
var tracks = []*Track{
	{
		Title:  "Closer",
		Artist: "The Chainsmokers",
		Album:  "Collage",
		Year:   2016,
		Length: 244 * time.Second,
	},
	{
		Title:  "Closer",
		Artist: "Marshmellow",
		Album:  "Forever",
		Year:   2018,
		Length: 244 * time.Second,
	},
	{
		Title:  "Don't Let Me Down",
		Artist: "The Chainsmokers",
		Album:  "Collage",
		Year:   2016,
		Length: 208 * time.Second,
	},
	{
		Title:  "Something Just Like This",
		Artist: "The Chainsmokers",
		Album:  "Memories...Do Not Open",
		Year:   2017,
		Length: 247 * time.Second,
	},
	{
		Title:  "Paris",
		Artist: "The Chainsmokers",
		Album:  "Memories...Do Not Open",
		Year:   2017,
		Length: 221 * time.Second,
	},
	{
		Title:  "Paris",
		Artist: "Dua Lipa",
		Album:  "Journey",
		Year:   2017,
		Length: 221 * time.Second,
	},
	{
		Title:  "Roses",
		Artist: "The Chainsmokers",
		Album:  "Bouquet",
		Year:   2015,
		Length: 237 * time.Second,
	},
	{
		Title:  "Sick Boy",
		Artist: "The Chainsmokers",
		Album:  "Sick Boy",
		Year:   2018,
		Length: 221 * time.Second,
	},
	{
		Title:  "All We Know",
		Artist: "The Chainsmokers",
		Album:  "Collage",
		Year:   2016,
		Length: 194 * time.Second,
	},
	{
		Title:  "Selfie",
		Artist: "The Chainsmokers",
		Album:  "Non-album single",
		Year:   2014,
		Length: 183 * time.Second,
	},
	{
		Title:  "This Feeling",
		Artist: "The Chainsmokers",
		Album:  "Sick Boy",
		Year:   2018,
		Length: 197 * time.Second,
	},
	{
		Title:  "Push My Luck",
		Artist: "The Chainsmokers",
		Album:  "World War Joy",
		Year:   2019,
		Length: 212 * time.Second,
	},
}

// func length(duration string) time.Duration {
// 	timeDuration, err := time.ParseDuration(duration)
// 	if err != nil {
// 		panic(duration)
// 	}

// 	return timeDuration
// }

type byDuration []*Track

func (tracks byDuration) Len() int           { return len(tracks) }
func (tracks byDuration) Swap(i, j int)      { tracks[i], tracks[j] = tracks[j], tracks[i] }
func (tracks byDuration) Less(i, j int) bool { return tracks[i].Length < tracks[j].Length }

type customSort struct {
	tracks []*Track
	less   func(a, b *Track) bool
}

func (c customSort) Len() int           { return len(c.tracks) }
func (c customSort) Swap(i, j int)      { c.tracks[i], c.tracks[j] = c.tracks[j], c.tracks[i] }
func (c customSort) Less(i, j int) bool { return c.less(c.tracks[i], c.tracks[j]) }

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

func PlaylistMain() {
	sort.Sort(byDuration(tracks))
	sort.Sort(customSort{tracks, func(tA, tB *Track) bool {
		if tA.Title != tB.Title {
			return tA.Title < tB.Title
		}

		if tA.Artist != tB.Artist {
			return tA.Artist < tB.Artist
		}

		if tA.Length != tB.Length {
			return tA.Length < tB.Length
		}
		return false
	}})
	printTracks(tracks)
}
