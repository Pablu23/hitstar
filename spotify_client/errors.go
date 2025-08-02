package spotifyclient

import "fmt"

type SpotifyError struct {
	JsonError struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Reason  string `json:"reason"`
	} `json:"error"`
}

func (e *SpotifyError) Error() string {
	return fmt.Sprintf("Status: %d, Message: %s, Reason: %s", e.JsonError.Status, e.JsonError.Message, e.JsonError.Reason)
}
