package track

import "fmt"

type Track struct {
	Title  string
	Artist string
	Album  string
}

type TrackFieldName string

func GetCurrentTrack() (Track, error) {
	track := Track{}

	artist, err := getFieldOfCurrentTrack(Artist)
	if err != nil {
		return track, fmt.Errorf("unable to get artist of current track: %w", err)
	}
	track.Artist = artist

	title, err := getFieldOfCurrentTrack(Title)
	if err != nil {
		return track, fmt.Errorf("unable to get title of current track: %w", err)
	}
	track.Title = title

	album, err := getFieldOfCurrentTrack(Album)
	if err != nil {
		return track, fmt.Errorf("unable to get album of current track: %w", err)
	}
	track.Album = album

	return track, nil
}
