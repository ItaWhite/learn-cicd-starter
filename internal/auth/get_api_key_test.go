package auth

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetAPIKey_Valid(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey 1488")

	s, err := GetAPIKey(headers)

	require.Equal(t, "1488", s)
	require.NoError(t, err)
}

func TestGetAPIKey_Empty(t *testing.T) {
	headers := http.Header{}

	s, err := GetAPIKey(headers)

	require.Equal(t, "", s)
	require.EqualError(t, err, ErrNoAuthHeaderIncluded.Error())
}

func TestGetAPIKey_MalformedHeader(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey")

	s, err := GetAPIKey(headers)

	require.Equal(t, "", s)
	require.EqualError(t, err, "malformed authorization header")
}
