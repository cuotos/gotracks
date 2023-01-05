package cmd

import (
	"testing"

	"github.com/cuotos/gotracks/track"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateUGUrl(t *testing.T) {
	tcs := []struct {
		inputTrack track.Track
		expected   string
	}{
		{
			inputTrack: track.Track{
				Artist: "mock Artist",
				Title:  "mock Title",
			},
			expected: "mock?search_type=title&value=mock+Artist+mock+Title",
		},
		{
			inputTrack: track.Track{
				Artist: "",
				Title:  "",
			},
			expected: "mock?search_type=title&value=+",
		},
	}

	for _, tc := range tcs {
		c := UGClient{
			baseSearchURL: "mock",
		}

		u, err := c.generateSearchURL(tc.inputTrack, All)
		require.NoError(t, err)
		assert.Equal(t, tc.expected, u.String())
	}
}

func TestQuerySpecificType(t *testing.T) {
	tcs := []struct {
		track      track.Track
		searchType Type
		expected   string
	}{
		{
			track.Track{
				Artist: "artist",
				Title:  "title",
			},
			Bass,
			"?search_type=title&type=400&value=artist+title",
		},
		{
			track.Track{
				Artist: "artist",
				Title:  "title",
			},
			All,
			"?search_type=title&value=artist+title",
		},
	}

	for _, tc := range tcs {
		c := UGClient{
			baseSearchURL: "",
		}

		actual, err := c.generateSearchURL(tc.track, tc.searchType)
		assert.NoError(t, err)
		assert.Equal(t, tc.expected, actual.String())
	}
}

func TestLookupType(t *testing.T) {
	tcs := []struct {
		input    []string
		expected Type
	}{
		{[]string{"All"}, All},
		{[]string{"Bass", "bass", "b"}, Bass},
		{[]string{"Tab", "tab", "t"}, Tab},
		{[]string{"Chords", "chords", "c"}, Chords},
		{[]string{"GuitarPro", "guitarPro", "pro"}, GuitarPro},
		{[]string{"Power", "power", "p"}, Power},
		{[]string{"Ukulele", "ukulele", "uke", "u"}, Ukulele},
		{[]string{"Official", "official", "o"}, Official},
		{[]string{"garbage"}, All},
	}

	for _, tc := range tcs {
		for _, inputType := range tc.input {
			actual := lookupType(inputType)
			assert.Equal(t, tc.expected, actual)
		}
	}
}
