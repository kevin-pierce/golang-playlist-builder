package authorizeUser

import (
	"context"
	"fmt"
	"github.com/zmb3/spotify/v2/auth"
	"log"
	"net/http"
	// "os"

	"github.com/joho/godotenv"
	"github.com/zmb3/spotify/v2"
)

const redirectURI = "http://localhost:8080/callback"

var (
	auth  = spotifyauth.New(spotifyauth.WithRedirectURL(redirectURI), spotifyauth.WithScopes(spotifyauth.ScopeUserReadPrivate, spotifyauth.ScopePlaylistModifyPublic))
	ch    = make(chan *spotify.Client)
	state = "abc123"
)

func AuthUser() (*spotify.Client, context.Context) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// first start an HTTP server
	http.HandleFunc("/callback", completeAuth)
	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	url := auth.AuthURL(state)

	// Consult on fix for this issue
	// tempSplit := strings.split(url, "client_id=")
	// newUrl = tempSplit[0] + "client_id=" + os.Getenv("SPOTIFY_ID")

	fmt.Println("Please log in to Spotify by visiting the following page in your browser:\n", url)

	// wait for auth to complete
	client := <-ch

	// use the client to make calls that require authorization
	return client, context.Background()
}

func completeAuth(w http.ResponseWriter, r *http.Request) {
	tok, err := auth.Token(r.Context(), state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, state)
	}

	// use the token to get an authenticated client
	client := spotify.New(auth.Client(r.Context(), tok))
	fmt.Fprintf(w, "Login Completed!")
	ch <- client
}
