package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/pablu23/spotify"
)

const redirectURI = "http://localhost:8080/callback"

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

func main() {
	secrets, err := readSecrets(".secret")
	if err != nil {
		log.Fatalf("unable to read secrets, err: %s", err)
	}

	server := spotify.NewServer(secrets["client_id"], secrets["client_secret"], redirectURI)
	http.HandleFunc("/login", server.Login)
	http.HandleFunc("/callback", server.CompleteAuth)

	// POST
	http.HandleFunc("/createLobby", server.CreateLobby)
	http.HandleFunc("/playlists", server.GetPlaylistsForUser)

	// Lobby
	http.HandleFunc("/lobby/{lobby}/playlist/{id}", server.SetPlaylistForLobby)
	http.HandleFunc("/lobby/{lobby}/track", server.GetRandomTrack)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request for:", r.URL.String())
	})

	http.ListenAndServe(":8080", nil)
}
