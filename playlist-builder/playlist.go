package playlist

import (
	"context"
	"fmt"
	"github.com/zmb3/spotify/v2"
	"log"
)

func BuildPlaylist(client *spotify.Client, ctx context.Context, songList []string) {
	user, err := client.CurrentUser(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("You are logged in as:", user.ID)

	result, err := client.Search(ctx, songList[0], spotify.SearchTypeTrack)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result.Tracks.Tracks[0].SimpleTrack.URI)

	// newPlaylist, err := client.CreatePlaylistForUser(context.Background(), user.ID, "TEST GOLANG", "Test for my golang application", true, false)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(newPlaylist)
}
