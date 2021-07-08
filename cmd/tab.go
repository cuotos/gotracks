package cmd

import (
	"fmt"
	"net/url"

	"github.com/andybrewer/mack"
	"github.com/cuotos/gotracks/track"
	"github.com/spf13/cobra"
)

var tabCmd = &cobra.Command{
	Use: "tab",
	RunE: func(cmd *cobra.Command, args []string) error {
		return openUG()
	},
}

func openUG() error {
	track, err := track.GetCurrentTrack()
	if err != nil {
		return err
	}

	searchText := url.QueryEscape(fmt.Sprintf("%v %v", track.Artist, track.Title))
	searchUri := fmt.Sprintf("https://www.ultimate-guitar.com/search.php?search_type=title&value=%v", searchText)

	_, err = mack.Tell("Google Chrome", fmt.Sprintf(`open location "%v"`, searchUri))
	if err != nil {
		return fmt.Errorf("unable to open browser: %w", err)
	}

	return nil
}
