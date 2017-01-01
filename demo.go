package main

import (
    "bitbucket.org/Martinyuk/discogs/api"
)

func main() {
    client := api.NewDiscogsClient(nil, "HNmJpKxApdkwljeHZxXRMFGgGVMfsODoOJojIXfh")
    
    // Search artists and releases by text.
    search, err := client.Search("Океан Ельзи")
    if err != nil {
        panic(err)
    }

    println("Search results: ")
    for i := 0; i < len(search.Results); i++ {
        track := search.Results[i]
        println("Title ", track.Title, " type: ", track.Type)
	}

    // Get specified artist by ID.
    artist, err := client.GetArtistByID("125748")
    if err != nil {
        panic(err)
    }

    println("Name ", artist.Name, " Profile: ", artist.Profile)

    // Get all releases of specified specified artist by artist ID.
    releases, err := client.GetReleasesByArtistID("125748")
    if err != nil {
        panic(err)
    }

    println("Releases: ")
    for i := 0; i < len(releases.Releases); i++ {
        release := releases.Releases[i]
        println("Title ", release.Title, " year: ", release.Year)
	}
}




