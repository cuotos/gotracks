//go:build darwin

package cmd

import (
	"fmt"
	"net/url"

	"github.com/andybrewer/mack"
)

func openBrowser(url *url.URL) error {

	_, err := mack.Tell("Google Chrome", fmt.Sprintf(`open location "%v"`, url.String()))
	if err != nil {
		return fmt.Errorf("unable to open browser: %w", err)
	}

	return nil
}
