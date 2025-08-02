package models

type Album struct {
	ReleaseDate string `json:"release_date"`
}

type Track struct {
	Name  string `json:"name"`
	Id    string `json:"id"`
	Uri   string `json:"uri"`
	Album Album  `json:"album"`
}

type PlaylistItems struct {
	Total int `json:"total"`
	Items []struct {
		Track Track `json:"track"`
	} `json:"items,omitempty"`
}
