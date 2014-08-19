package main

import (
	"github.com/jamesandariese/notmuch_addresses"
	//"net/mail"
	"fmt"
	"os"
	"flag"
)

func main() {
	flag.Parse()
	err := notmuch_addresses.InitDatabase(os.ExpandEnv("$HOME/.notmuch_addresses.sqlite3"))
	if err != nil {
		panic("Couldn't open database: " + err.Error())
	}

	for _, path := range flag.Args() {
		addresses, err := notmuch_addresses.GatherAddresses(path)
		if err != nil {
			panic("Couldn't gather email addresses: " + err.Error())
		}
		fmt.Printf("Saved %d addresses\n", addresses)
	}
	err = notmuch_addresses.Close()
	if err != nil {
		panic("Couldn't close database: " + err.Error())
	}
}


