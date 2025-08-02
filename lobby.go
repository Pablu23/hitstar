package spotify

import (
	"crypto/rand"
	"encoding/json"
	"net/http"

	"github.com/pablu23/spotify/models"
)

type Lobby struct {
	Host    *Player
	Players []*Player

	Playlist models.Playlist
	// Playlist which is played
	// Scores
	// Other settings
}

func (s *Server) SetPlaylistForLobby(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session")
	if err != nil {
		http.Error(w, "unable to get session", http.StatusUnauthorized)
		return
	}

	lobbyId := r.PathValue("lobby")
	lobby, ok := s.Lobbies[lobbyId]
	if !ok {
		http.Error(w, "unable to find lobby", http.StatusNotFound)
		return
	}

	if lobby.Host.session != session.Value {
		http.Error(w, "only host can change Playlist", http.StatusForbidden)
		return
	}

	playlistId := r.PathValue("id")
	client, ok := s.Sessions[session.Value]
	if !ok {
		http.Error(w, "unable to find session", http.StatusNotFound)
		return
	}

	playlist, err := client.GetPlaylist(playlistId, "id", "tracks.total", "name")
	if err != nil {
		http.Error(w, "unable to get playlist", http.StatusInternalServerError)
		return
	}

	lobby.Playlist = playlist

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&playlist)
}

func (s *Server) CreateLobby(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session")
	if err != nil {
		http.Error(w, "unable to get session", 401)
		return
	}

	lobbyCode := rand.Text()

	host := NewPlayer("Host", s.Sessions[session.Value], session.Value)
	lobby := Lobby{
		Host: &host,
		Players: []*Player{
			&host,
		},
	}

	s.Lobbies[lobbyCode] = &lobby
	w.Write([]byte(lobbyCode))
}
