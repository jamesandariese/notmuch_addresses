package main

import (
	"flag"
	"fmt"
	"github.com/jamesandariese/notmuch_addresses"
	"net/mail"
	"os"
)

var muttMode bool

func init() {
	flag.BoolVar(&muttMode, "mutt", false, "Query in Mutt Mode, response for $query_command [see muttrc(5)]")
}

func main() {
	flag.Parse()

	err := notmuch_addresses.InitDatabase(os.ExpandEnv("$HOME/.notmuch_addresses.sqlite3"))
	if err != nil {
		panic("Couldn't open database: " + err.Error())
	}

	if muttMode {
		fmt.Println("Query Notmuch Addresses Output")
		ch, cherr := notmuch_addresses.QueryToChannel(flag.Arg(0))
		for {
			select {
			case raw, ok := <-ch:
				if !ok {
					return
				}
				addr, err := mail.ParseAddress(raw)
				if err != nil {
					continue
				}
				fmt.Printf("%s\t%s\t\n", addr.Address, addr.Name)
			case <-cherr:
			}
		}
	} else {
		err = notmuch_addresses.QueryToStdout(flag.Arg(0))
		if err != nil {
			panic("Couldn't query database: " + err.Error())
		}
	}
}
