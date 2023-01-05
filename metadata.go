package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/docopt/docopt-go"
	"github.com/nbd-wtf/go-nostr"
)

type Metadata struct {
	Name    string `json:"name,omitempty"`
	About   string `json:"about,omitempty"`
	Picture string `json:"picture,omitempty"`
	Nip05	string `json:"nip05,omitempty"`
}

func setMetadata(opts docopt.Opts) {
	initNostr()

	name, _ := opts.String("--name")
	about, _ := opts.String("--about")
	picture, _ := opts.String("--picture")
	nip05, _ := opts.String("--nip05")

	jmetadata, _ := json.Marshal(Metadata{
		Name:    name,
		About:   about,
		Picture: picture,
		Nip05: nip05,
	})

	event, statuses, err := pool.PublishEvent(&nostr.Event{
		PubKey:    getPubKey(config.PrivateKey),
		CreatedAt: time.Now(),
		Kind:      nostr.KindSetMetadata,
		Tags:      make(nostr.Tags, 0),
		Content:   string(jmetadata),
	})
	if err != nil {
		log.Printf("Error publishing: %s.\n", err.Error())
		return
	}

	printPublishStatus(event, statuses)
}
