package cmd

import (
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"

	"github.com/andybrewer/mack"
	"github.com/cuotos/gotracks/track"
	"github.com/spf13/cobra"
)

const (
	UGSearchURL = "https://www.ultimate-guitar.com/search.php"
)

type Type int

const (
	All       Type = 0
	Tab       Type = 200
	Chords    Type = 300
	Bass      Type = 400
	GuitarPro Type = 500
	Power     Type = 600
	Ukulele   Type = 800
	Official  Type = 900
)

func (t Type) String() string {
	return strconv.Itoa(int(t))
}

func lookupType(input string) Type {
	switch strings.ToLower(input) {
	case "all":
		return All
	case "bass", "b":
		return Bass
	case "tab", "t":
		return Tab
	case "chords", "c":
		return Chords
	case "guitarpro", "pro":
		return GuitarPro
	case "power", "p":
		return Power
	case "ukulele", "u", "uke":
		return Ukulele
	case "official", "o":
		return Official
	default:
		log.Printf("unknown type %s. defaulting to All", input)
		return All
	}
}

func NewTabCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "tab",
		RunE: OpenTab,
	}

	cmd.Flags().String("type", "all", `tab type to get`)
	cmd.Flags().Bool("dryrun", false, "print url and dont open chrome")
	cmd.Flags().MarkHidden("dryrun")

	return cmd
}

func OpenTab(cmd *cobra.Command, args []string) error {

	ugc := UGClient{
		baseSearchURL: UGSearchURL,
	}

	track, err := track.GetCurrentTrack()
	if err != nil {
		return err
	}

	t, err := cmd.Flags().GetString("type")
	if err != nil {
		return err
	}

	url, err := ugc.generateSearchURL(track, lookupType(t))
	if err != nil {
		return err
	}

	dryrun, _ := cmd.Flags().GetBool("dryrun")
	if dryrun {
		fmt.Println(url)
		return nil
	}

	_, err = mack.Tell("Google Chrome", fmt.Sprintf(`open location "%v"`, url.String()))

	if err != nil {
		return fmt.Errorf("unable to open browser: %w", err)
	}
	return nil
}

type UGClient struct {
	baseSearchURL string
}

func (ugClient *UGClient) generateSearchURL(track track.Track, t Type) (*url.URL, error) {

	searchText := fmt.Sprintf("%v %v", track.Artist, track.Title)

	ugUrl, err := url.Parse(ugClient.baseSearchURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ultimate guitar search url: %w", err)
	}

	q := ugUrl.Query()
	q.Add("search_type", "title")
	q.Add("value", searchText)

	if t != All {
		q.Add("type", t.String())
	}

	ugUrl.RawQuery = q.Encode()

	return ugUrl, nil
}
