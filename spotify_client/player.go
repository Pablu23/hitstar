package spotifyclient

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/pablu23/hipstar/models"
)

func (c *SpotifyClient) EnqueuePlayer(enqueueUri string) error {
	uri := api("me/player/queue").WithQuery("uri", enqueueUri).String()

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

func (c *SpotifyClient) ResumePlayer() error {
	uri := api("me/player/play").String()

	request, err := http.NewRequest("PUT", uri, nil)
	if err != nil {
		return err
	}

	resp, err := c.client.Do(request)
	if err != nil {
		return err
	} else if resp.StatusCode != 204 {
		var err SpotifyError
		json.NewDecoder(resp.Body).Decode(&err)
		return &err
	}

	return nil
}

func (c *SpotifyClient) PlayTrack(enqueueUri string) error {
	uri := api("me/player/play").String()

	body := struct {
		Uris []string `json:"uris"`
	}{
		Uris: []string{enqueueUri},
	}
	
	buf, err := json.Marshal(&body)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("PUT", uri, bytes.NewBuffer(buf))
	if err != nil {
		return err
	}

	resp, err := c.client.Do(request)
	if err != nil {
		return err
	} else if resp.StatusCode != 204 {
		var err SpotifyError
		json.NewDecoder(resp.Body).Decode(&err)
		return &err
	}

	return nil
}
