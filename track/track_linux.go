//go:build linux

package track

import (
	"fmt"

	"github.com/godbus/dbus/v5"
)

const (
	Title  TrackFieldName = "title"
	Artist TrackFieldName = "artist"
	Album  TrackFieldName = "album"
)

func getFieldOfCurrentTrack(field TrackFieldName) (string, error) {
	var fieldResult string
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		return fieldResult, err
	}
	defer conn.Close()
	obj := conn.Object("org.mpris.MediaPlayer2.spotify", "/org/mpris/MediaPlayer2")
	var metadata map[string]dbus.Variant

	// Call the Get method on the Properties interface
	err = obj.Call("org.freedesktop.DBus.Properties.Get", 0, "org.mpris.MediaPlayer2.Player", "Metadata").Store(&metadata)
	if err != nil {
		return fieldResult, nil
	}

	metadataField := fmt.Sprintf("xesam:%s", field)
	if fieldResult, ok := metadata[metadataField]; ok {
		// value returns the underlying type, which might be an string or []string.
		// the dbus stuff returns an array of strings for artist so if it does
		// cast it to []string and get the first one
		switch v := fieldResult.Value().(type) {
		case []string:
			return v[0], nil
		case string:
			return v, nil
		}
		return fieldResult.String(), nil
	} else {
		return "", fmt.Errorf("unable to get field %s from spotify", field)
	}
}
