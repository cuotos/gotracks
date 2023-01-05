package cmd

import (
	"fmt"
	"net/url"

	"github.com/andybrewer/mack"
	"github.com/cuotos/gotracks/track"
	"github.com/spf13/cobra"
)

const (
	UGSearchURL = "https://www.ultimate-guitar.com/search.php"
)

var TabCmd = &cobra.Command{
	Use: "tab",
	RunE: func(cmd *cobra.Command, args []string) error {
		ugc := UGClient{
			baseSearchURL: UGSearchURL,
		}

		track, err := track.GetCurrentTrack()
		if err != nil {
			return err
		}

		url, err := ugc.generateSearchURL(track)
		if err != nil {
			return err
		}

		_, err = mack.Tell("Google Chrome", fmt.Sprintf(`open location "%v"`, url.String()))

		if err != nil {
			return fmt.Errorf("unable to open browser: %w", err)
		}
		return nil
	},
}

type UGClient struct {
	baseSearchURL string
}

func (ugClient *UGClient) generateSearchURL(track track.Track) (*url.URL, error) {

	searchText := url.QueryEscape(fmt.Sprintf("%v %v", track.Artist, track.Title))

	ugUrl, err := url.Parse(ugClient.baseSearchURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ultimate guitar search url: %w", err)
	}

	q := ugUrl.Query()
	q.Add("search_type", "title")
	q.Add("value", searchText)
	ugUrl.RawQuery = q.Encode()

	return ugUrl, nil
}
