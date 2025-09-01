package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/pablu23/hipstar"
)

const redirectURI = "http://api.hitstar.xyz:8080/callback"

func readSecrets(path string) (map[string]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	result := make(map[string]string)
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		splits := strings.SplitN(line, "=", 2)

		name := strings.TrimSpace(splits[0])
		value := strings.TrimSpace(splits[1])
		result[name] = value
	}

	return result, nil
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://hitstar.xyz:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	secrets, err := readSecrets(".secret")
	if err != nil {
		log.Fatalf("unable to read secrets, err: %s", err)
	}

	server := spotify.NewServer(secrets["client_id"], secrets["client_secret"], redirectURI)

	mux := http.NewServeMux()

	mux.HandleFunc("/login", server.Login)
	mux.HandleFunc("/callback", server.CompleteAuth)

	// POST
	mux.HandleFunc("/createLobby", server.CreateLobby)
	mux.HandleFunc("/playlists", server.GetPlaylistsForUser)

	// Lobby
	mux.HandleFunc("/lobby/{lobby}/join", server.JoinLobby)
	mux.HandleFunc("/lobby/{lobby}/players/list", server.ListLobbyPlayers)
	mux.HandleFunc("/lobby/{lobby}/playlist/{id}", server.SetPlaylistForLobby)
	mux.HandleFunc("/lobby/{lobby}/track", server.GetRandomTrack)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request for:", r.URL.String())
	})

	fmt.Println("Starting Server on :8080")
	http.ListenAndServe(":8080", corsMiddleware(mux))
}
