package main

import (
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {

	track, err := GetCurrentTrack()
	if err != nil {
		return err
	}

	openUG(track)
	return nil
}
