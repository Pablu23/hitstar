package spotify

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/http"

	spotifyclient "github.com/pablu23/hipstar/spotify_client"
	"golang.org/x/oauth2"
)

type Player struct {
	Name          string
	spotifyClient *spotifyclient.SpotifyClient
	session       string
}

func NewPlayer(name string, client *spotifyclient.SpotifyClient, session string) Player {
	return Player{
		Name:          name,
		spotifyClient: client,
		session:       session,
	}
}

func (s *Server) GetPlaylistsForUser(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session")
	if err != nil {
		http.Error(w, "unable to get session", 401)
		return
	}

	client, ok := s.Sessions[session.Value]
	if !ok {
		http.Error(w, "unable to find session", http.StatusNotFound)
		return
	}

	playlists, err := client.GetUserPlaylists(20, 0)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&playlists)
}

func (s *Server) CompleteAuth(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	state := r.FormValue("state")

	verifier, ok := s.States[state]
	if !ok {
		http.NotFound(w, r)
	}

	tok, err := s.oauthConf.Exchange(context.TODO(), code, oauth2.VerifierOption(verifier))
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to exchange token: %s", err), 401)
		return
	}

	session := rand.Text()

	client := s.oauthConf.Client(context.TODO(), tok)
	s.Sessions[session] = spotifyclient.NewClient(client)
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    session,
		Quoted:   false,
		Secure:   false,
		HttpOnly: false,
	})

	delete(s.States, state)

	fmt.Fprintf(w, "logged in")
}

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	state := rand.Text()
	codeVerifier := oauth2.GenerateVerifier()
	authURL := s.oauthConf.AuthCodeURL(state, oauth2.S256ChallengeOption(codeVerifier))
	http.Redirect(w, r, authURL, http.StatusTemporaryRedirect)

	s.States[state] = codeVerifier
}
