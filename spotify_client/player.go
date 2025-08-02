package spotifyclient

import (
	"encoding/json"

	"github.com/pablu23/spotify/models"
)

func (c *SpotifyClient) PostPlayerQueue(enqueUri string) error {
	uri := api("me/player/queue").WithQuery("uri", enqueUri).String()

	resp, err := c.client.Post(uri, "text/plain", nil)
	if err != nil {
		return err
	} else if resp.StatusCode != 200 {
		var err SpotifyError
		json.NewDecoder(resp.Body).Decode(&err)
		return &err
	}
	return err
}

func (c *SpotifyClient) GetAvailableDevices() (models.Devices, error) {
	uri := api("me/player/devices").String()

	var devices models.Devices
	err := c.getJsonFor(uri, &devices)

	return devices, err
}
