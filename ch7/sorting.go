package main

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

type byArtist []*Track

func (b byArtist) Len() int               { return len(b) }
func (b byArtist) Less(i int, j int) bool { return b[i].Artist < b[j].Artist }
func (b byArtist) Swap(i int, j int)      { b[i], b[j] = b[j], b[i] }

func main() {
	tracks := []*Track{
		{Title: "Go", Artist: "Delilah", Album: "From the Roots Up", Year: 2012, Length: length("3m38s")},
		{Title: "Go", Artist: "Moby", Album: "Moby", Year: 1992, Length: length("3m37s")},
		{Title: "Go Ahead", Artist: "Alicia Keys", Album: "As I Am", Year: 2007, Length: length("4m36s")},
		{Title: "Ready 2 Go", Artist: "Martin Solveig", Album: "Smash", Year: 2011, Length: length("4m24s")},
	}

	fmt.Println("Before Sort")
	printTracks(tracks)
	fmt.Println()

	fmt.Println("Sort")
	sort.Sort(byArtist(tracks))
	printTracks(tracks)
	fmt.Println()

	fmt.Println("Reverse Sort")
	sort.Sort(sort.Reverse(byArtist(tracks)))
	printTracks(tracks)
	fmt.Println()
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")

	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}

	tw.Flush()
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}
