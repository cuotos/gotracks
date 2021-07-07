package main

import (
	"fmt"
	"net/url"

	"github.com/andybrewer/mack"
)

func openUG(track Track) error {

	searchText := url.QueryEscape(fmt.Sprintf("%v %v", track.Artist, track.Title))
	searchUri := fmt.Sprintf("https://www.ultimate-guitar.com/search.php?search_type=title&value=%v", searchText)

	_, err := mack.Tell("Google Chrome", fmt.Sprintf(`open location "%v"`, searchUri))
	if err != nil {
		return fmt.Errorf("unable to open browser: %w", err)
	}

	return nil
}
