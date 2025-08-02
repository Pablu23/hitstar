package spotifyclient

import (
	"strings"
	"testing"
)

func TestApi(t *testing.T) {
	subUri := "playlists"
	want := apiBase + "/" + subUri

	if api(subUri).String() != want {
		t.Errorf(`api("%s").String() != %s`, subUri, want)
	}
}

func TestApiWithFields(t *testing.T) {
	subUri := "playlists"
	fields := []string{
		"id",
		"name",
		"track(name)",
	}
	want := apiBase + "/" + subUri + "?fields=" + strings.Join(fields, ",")

	if api(subUri).WithFields(fields...).String() != want {
		t.Errorf(`api("%s").WithFields(%s).String() != %s`, subUri, strings.Join(fields, ", "), want)
	}
}

func TestApiWithQuery(t *testing.T) {
	subUri := "playlists"
	want := apiBase + "/" + subUri + "?length=50&offset=100&name=test"

	result := api(subUri).WithQuery("length", "50", "offset", "100", "name", "test").String()

	if result != want {
		t.Errorf(`api("%s").WithQuery("length", "50", "offset", "100", "name", "test").String() = %s but was not %s`, subUri, result, want)
	}
}

func TestApiWithQueryAndFields(t *testing.T) {
	subUri := "playlists"
	want := apiBase + "/" + subUri + "?length=50&offset=100&name=test" + "&fields=id,name"

	res := api(subUri).WithQuery("length", "50", "offset", "100", "name", "test").WithFields("id", "name").String()
	if res != want {
		t.Errorf(`api("%s").WithQuery("length", "50", "offset", "100", "name", "test").WithFields("id", "name").String() = %s but was not %s`, subUri, res, want)
	}
}

func TestApiWithFieldsAndQuery(t *testing.T) {
	subUri := "playlists"
	want := apiBase + "/" + subUri + "?fields=id,name" + "&length=50&offset=100&name=test"

	res := api(subUri).WithFields("id", "name").WithQuery("length", "50", "offset", "100", "name", "test").String()
	if res != want {
		t.Errorf(`api("%s").WithFields("id", "name").WithQuery("length", "50", "offset", "100", "name", "test").String() = %s but was not %s`, subUri, res, want)
	}
}
