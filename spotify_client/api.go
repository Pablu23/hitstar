package spotifyclient

import (
	"fmt"
	"net/url"
	"slices"
	"strings"
)

type apiUrl struct {
	uri        string
	extraQuery []string
}

func api(subUri ...string) apiUrl {
	uri, _ := url.JoinPath(apiBase, subUri...)
	return apiUrl{
		uri:        uri,
		extraQuery: make([]string, 0),
	}
}

func (a apiUrl) String() string {
	complete := a.uri

	if len(a.extraQuery) == 0 {
		return complete
	}

	complete += "?"
	for i, query := range a.extraQuery {
		complete += query

		if i != len(a.extraQuery)-1 {
			complete += "&"
		}
	}
	 
	return complete
}

func (a apiUrl) WithFields(fields ...string) apiUrl {
	fieldsRaw := strings.Join(fields, ",")
	a.extraQuery = append(a.extraQuery, fmt.Sprintf("fields=%s", fieldsRaw))
	return a
}

func (a apiUrl) WithQuery(nameAndValue ...string) apiUrl {
	if len(nameAndValue) % 2 != 0 {
		panic("query parameters must be multiple of two, WithQuery(\"name\", \"value\")")
	}

	extraNeededSize := len(nameAndValue) / 2
	if len(a.extraQuery) == 0 {
		a.extraQuery = make([]string, 0, extraNeededSize)
	} else {
		newArr := make([]string, len(a.extraQuery), len(a.extraQuery) + extraNeededSize)
		copy(newArr, a.extraQuery)
		a.extraQuery = newArr
	}

	for pair := range slices.Chunk(nameAndValue, 2) {
		a.extraQuery = append(a.extraQuery, fmt.Sprintf("%s=%s", pair[0], pair[1]))
	}

	return a
}
