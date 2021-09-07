package authorizeUser

import (
	"context"
	"fmt"
	"github.com/zmb3/spotify/v2/auth"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/zmb3/spotify/v2"
	"golang.org/x/oauth2/clientcredentials"
)

func AuthUser() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("SPOTIFY_ID"),
		ClientSecret: os.Getenv("SPOTIFY_SECRET"),
		TokenURL:     spotifyauth.TokenURL,
	}

	token, err := config.Token(ctx)
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}

	httpClient := spotifyauth.New().Client(ctx, token)
	client := spotify.New(httpClient)
	msg, page, err := client.FeaturedPlaylists(ctx)
	if err != nil {
		log.Fatalf("couldn't get features playlists: %v", err)
	}

	fmt.Println(msg)
	for _, playlist := range page.Playlists {
		fmt.Println("  ", playlist.Name)
	}
}
