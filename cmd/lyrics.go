package cmd

import (
	"fmt"

	"github.com/cuotos/gotracks/track"
	"github.com/spf13/cobra"
)

var lyricsCmd = &cobra.Command{
	Use: "lyrics",
	RunE: func(cmd *cobra.Command, args []string) error {
		return getLyrics()
	},
}

func getLyrics() error {
	track, err := track.GetCurrentTrack()
	if err != nil {
		return err
	}

	fmt.Println(generateAZLyricsURL(track))
	return nil
}

func generateAZLyricsURL(track track.Track) (string, error) {
	urlFormat := "https://www.azlyrics.com/lyrics/%s/%s.html"

	url := fmt.Sprintf(urlFormat, track.Artist, track.Title)

	return url, nil
}
