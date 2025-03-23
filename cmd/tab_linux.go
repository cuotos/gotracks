//go:build linux

package cmd

import (
	"fmt"
	"net/url"
	"os/exec"
)

func openBrowser(url *url.URL) error {

	err := exec.Command("xdg-open", url.String()).Start()
	if err != nil {
		return fmt.Errorf("unable to open browser: %w", err)
	}

	return nil
}
