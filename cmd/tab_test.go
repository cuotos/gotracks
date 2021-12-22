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
				Artist: "mockArtist",
				Title:  "mockTitle",
			},
			expected: "mock?search_type=title&value=mockArtist%2BmockTitle",
		},
		{
			inputTrack: track.Track{
				Artist: "",
				Title:  "",
			},
			expected: "mock?search_type=title&value=%2B",
		},
	}

	for _, tc := range tcs {
		c := UGClient{
			baseSearchURL: "mock",
		}

		u, err := c.generateSearchURL(tc.inputTrack)
		require.NoError(t, err)
		assert.Equal(t, tc.expected, u.String())
	}
}
