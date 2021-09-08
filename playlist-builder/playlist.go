package playlist

import (
	"context"
	"fmt"
	"github.com/zmb3/spotify/v2"
	"log"
)

func BuildPlaylist(client *spotify.Client, ctx context.Context, songList []string) {
	var uriList []spotify.ID
	user, err := client.CurrentUser(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("You are logged in as:", user.ID)

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
	newPlaylist, err := client.CreatePlaylistForUser(ctx, user.ID, "Top 100 (according to Golang)", "Billboard Top 100 songs, compiled by a Golang application", true, false)
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
