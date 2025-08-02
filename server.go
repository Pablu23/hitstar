package spotify

import (
	spotifyclient "github.com/pablu23/hipstar/spotify_client"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/spotify"
)

type Playlist struct {
	Id          string
	TrackNumber int
}

type Server struct {
	Sessions map[string]*spotifyclient.SpotifyClient

	States map[string]string

	Lobbies map[string]*Lobby

	oauthConf oauth2.Config
}

func NewServer(clientId string, clientSecret string, redirectUri string) *Server {
	return &Server{
		Sessions: make(map[string]*spotifyclient.SpotifyClient),
		States:   make(map[string]string),
		Lobbies:  make(map[string]*Lobby),
		oauthConf: oauth2.Config{
			ClientID:     clientId,
			ClientSecret: clientSecret,
			Endpoint:     spotify.Endpoint,
			RedirectURL:  redirectUri,
			Scopes: []string{
				"playlist-read-private",
				"user-modify-playback-state",
			},
		},
	}

}
