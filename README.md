go-api-discogs
===========

Go api wrapper for discogs.com

Example:
```go
package main

import (
	"github.com/AlexanderMartinyuk/go-api-discogs/api"
	"strconv"
)

func main() {
	client := api.NewDiscogsClient(nil, "your_access_key")

	// Search for artists and releases by text.
	search, err := client.Search("Океан Ельзи")
	if err != nil {
		panic(err)
	}

	println("Search results: ")
	for i := 0; i < len(search.Results); i++ {
		result := search.Results[i]
		println("Title: ", result.Title, " type: ", result.Type, " id: ", strconv.Itoa(result.ID))
	}

	// Get specified artist by ID.
	// This is immutable ID of Okean Elzy group in Discogs database. You can get this ID as result of Search.
	const okeanElzyGroupID = 125748
	artist, err := client.GetArtistByID(okeanElzyGroupID)
	if err != nil {
		panic(err)
	}

	println("Name: ", artist.Name, " Profile: ", artist.Profile)

	// Get all releases of specified artist by artist ID.
	releases, err := client.GetReleasesByArtistID(okeanElzyGroupID)
	if err != nil {
		panic(err)
	}

	println("Releases: ")
	for i := 0; i < len(releases.Releases); i++ {
		release := releases.Releases[i]
		println("Title: ", release.Title, " year: ", release.Year, " id: ", strconv.Itoa(release.ID))
	}

	// Get release details by release ID.
	const yananebibuvReleaseID = 111886
	details, err := client.GetReleaseDetailsByID(yananebibuvReleaseID)
	if err != nil {
		panic(err)
	}

	println("Tracks: ")
	for i := 0; i < len(details.Tracklist); i++ {
		track := details.Tracklist[i]
		println("Title: ", track.Title, " position: ", track.Position, " duration: ", track.Duration)
	}
}
```
