package playlist

import (
	"context"
	"fmt"
	"log"

	"github.com/zmb3/spotify/v2"
)

func BuildPlaylist(client *spotify.Client, ctx context.Context, songList []string, weekOf string) {
	var uriList []spotify.ID
	user, err := client.CurrentUser(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Logged in!")

	fmt.Println("Getting songs...")
	for _, songName := range songList {
		result, err := client.Search(ctx, songName, spotify.SearchTypeTrack)
		if err != nil {
			log.Fatal(err)
		}

		if len(result.Tracks.Tracks) > 0 {
			songURI := result.Tracks.Tracks[0].SimpleTrack.ID
			uriList = append(uriList, songURI)
		}
	}

	fmt.Println("Creating playlist...")
	descStr := fmt.Sprintf("Billboard Top 100 songs, compiled by a Golang application for the week of %s", weekOf)

	newPlaylist, err := client.CreatePlaylistForUser(ctx, user.ID, "Top 100 (according to Golang)", descStr, true, false)
	if err != nil {
		log.Fatal(err)
	}
	newPlaylistID := newPlaylist.SimplePlaylist.ID

	version, err := client.AddTracksToPlaylist(ctx, newPlaylistID, uriList...)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Succesfully created playlist!")
		fmt.Println(version)
	}
}
