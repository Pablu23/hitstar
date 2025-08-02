package models

type Playlist struct {
	Id    string        `json:"id"`
	Name  string        `json:"name"`
	Items PlaylistItems `json:"tracks"`
}
