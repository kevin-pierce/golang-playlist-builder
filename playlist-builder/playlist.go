package playlist

import (
	"context"
	"fmt"
	"github.com/zmb3/spotify/v2"
)

func BuildPlaylist(client *spotify.Client, ctx context.Context) {
	fmt.Println(client)
	fmt.Println(ctx)
}
