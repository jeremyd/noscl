package main

import (
	"fmt"
	"log"

	"github.com/docopt/docopt-go"
	"github.com/nbd-wtf/go-nostr"
)

func getContacts(opts docopt.Opts) {
	
	pubkey := getPubKey(config.PrivateKey)

	//pubkey := nip19.TranslatePublicKey(opts["<pubkey>"].(string))
	if pubkey == "" {
		log.Println("Profile key is empty! Exiting.")
		return
	}

	initNostr()

	//verbose, _ := opts.Bool("--verbose")

	var keys []string

	keys = append(keys, pubkey)

	_, all := pool.Sub(nostr.Filters{{Authors: keys}})
	for event := range nostr.Unique(all) {
		if event.Kind == nostr.KindContactList {
			fmt.Printf("found a contact list of %d follows\n", len(event.Tags))
			for _,contact := range event.Tags {
				fmt.Printf("key:%s, relay:%s\n", contact[1], contact[2])
				found := false
				for _,f := range config.Following {
					if f.Key == contact[1] {
						found = true
					}
				}
				if !found {
					fmt.Println("found a new follow.. adding")
					config.Following = append(config.Following, Follow{
						Key: contact[1],
					})
					saveConfig("/tmp/nostr.json")
				}
			}
		}
	}
}
