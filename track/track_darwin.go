//go:build darwin

package track

import (
	"fmt"

	"github.com/andybrewer/mack"
)

const (
	Title  TrackFieldName = "name"
	Artist TrackFieldName = "artist"
	Album  TrackFieldName = "album"
)

func getFieldOfCurrentTrack(field TrackFieldName) (string, error) {

	val, err := mack.Tell("Spotify", fmt.Sprintf("%s of current track as string", field))
	if err != nil {
		return "", err
	}

	return val, nil
}
