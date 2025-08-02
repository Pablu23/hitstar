package spotifyclient

import (
	"encoding/json"
	"net/http"
)

type SpotifyClient struct {
	client *http.Client
}

func NewClient(client *http.Client) *SpotifyClient {
	return &SpotifyClient{
		client: client,
	}
}

func (c *SpotifyClient) getJsonFor(uri string, v any) error {
	resp, err := c.client.Get(uri)
	if err != nil {
		return err
	}

	err = json.NewDecoder(resp.Body).Decode(v)
	if err != nil {
		return err
	}

	return nil
}
