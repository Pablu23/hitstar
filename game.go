package spotify

import (
	"encoding/json"
	"log/slog"
	"math/rand"
	"net/http"
)

func (s *Server) GetRandomTrack(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session")
	if err != nil {
		slog.Error(
			"unable to get session from cookie",
			"error", err,
		)
		http.Error(w, "unable to get session", http.StatusUnauthorized)
		return
	}

	lobbyId := r.PathValue("lobby")
	lobby, ok := s.Lobbies[lobbyId]
	if !ok {
		slog.Error(
			"unable to find lobby",
			"lobby", lobbyId,
			"error", err,
		)
		http.Error(w, "unable to find lobby", http.StatusNotFound)
		return
	}

	client, ok := s.Sessions[session.Value]
	if !ok {
		slog.Error(
			"unable to find session",
			"session", session.Value,
			"error", err,
		)
		http.Error(w, "unable to find session", http.StatusNotFound)
		return
	}

	chosen := rand.Intn(lobby.Playlist.Items.Total)

	playlistItems, err := client.GetPlaylistItems(lobby.Playlist.Id, 1, chosen)
	if err != nil {
		
		slog.Error(
			"unable to get track",
			"track_num", chosen,
			"error", err,
		)
		http.Error(w, "unable to get track", http.StatusInternalServerError)
		return
	}

	for _, player := range lobby.Players {
		err := player.spotifyClient.PostPlayerQueue(playlistItems.Items[0].Track.Uri)
		if err != nil {
			slog.Error(
				"unable to enqueue track",
				"error", err,
			)
			http.Error(w, "unable to enqueue track", http.StatusInternalServerError)
			return
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&playlistItems)
}
