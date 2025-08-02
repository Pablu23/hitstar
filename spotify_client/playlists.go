package spotifyclient

import (
	"strconv"

	"github.com/pablu23/spotify/models"
)

// limit=20 0-50
// offset=0
func (c *SpotifyClient) GetUserPlaylists(limit int, offset int, fields ...string) (models.Playlist, error) {
	uri := api("me", "playlists").WithQuery("limit", strconv.Itoa(limit), "offset", strconv.Itoa(offset)).WithFields(fields...).String()

	var playlist models.Playlist
	err := c.getJsonFor(uri, &playlist)

	return playlist, err
}

func (c *SpotifyClient) GetPlaylist(playlistId string, fields ...string) (models.Playlist, error) {
	uri := api("playlists", playlistId).WithFields(fields...).String()
	
	var playlist models.Playlist
	err := c.getJsonFor(uri, &playlist)

	return playlist, err
}

func (c *SpotifyClient) GetPlaylistItems(playlistId string, limit int, offset int, fields ...string) (models.PlaylistItems, error) {
	uri := api("playlists", playlistId, "tracks").
		WithQuery("limit", strconv.Itoa(limit), "offset", strconv.Itoa(offset)).
		WithFields(fields...).
		String()

	var playlistItems models.PlaylistItems
	err := c.getJsonFor(uri, &playlistItems)

	return playlistItems, err
}
