package main

import (
    "bitbucket.org/Martinyuk/discogs/api"
)

func main() {
    client := api.NewDiscogsClient(nil, "HNmJpKxApdkwljeHZxXRMFGgGVMfsODoOJojIXfh")
    result, err := client.Search("Amy Winehouse")
    if err != nil {

    }

    println("Search results: ")

    for i := 0; i < len(result.Results); i++ {
        track := result.Results[i]
        println("Title ", track.Title, " type: ", track.Type)
	}
}




