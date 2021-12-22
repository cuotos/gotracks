package track

import (
	"fmt"

	"github.com/andybrewer/mack"
)

type Track struct {
	Title  string
	Artist string
	Album  string
}

func GetCurrentTrack() (Track, error) {

	track := Track{}

	artist, err := getFieldOfCurrentTrack("artist")
	if err != nil {
		return track, fmt.Errorf("unable to get artist of current track: %w", err)
	}
	track.Artist = artist

	title, err := getFieldOfCurrentTrack("name")
	if err != nil {
		return track, fmt.Errorf("unable to get title of current track: %w", err)
	}
	track.Title = title

	album, err := getFieldOfCurrentTrack("album")
	if err != nil {
		return track, fmt.Errorf("unable to get album of current track: %w", err)
	}
	track.Album = album

	return track, nil

}

func getFieldOfCurrentTrack(field string) (string, error) {
	val, err := mack.Tell("Spotify", fmt.Sprintf("%s of current track as string", field))
	if err != nil {
		return "", err
	}

	return val, nil
}
