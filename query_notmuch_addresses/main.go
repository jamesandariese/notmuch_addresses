package main

import (
	"github.com/jamesandariese/notmuch_addresses"
	"os"
	"flag"
)

func main() {
	flag.Parse()
	
	err := notmuch_addresses.InitDatabase(os.ExpandEnv("$HOME/.notmuch_addresses.sqlite3"))
	if err != nil {
		panic("Couldn't open database: " + err.Error())
	}
	
	err = notmuch_addresses.QueryToStdout(flag.Arg(0))
	if err != nil {
		panic("Couldn't query database: " + err.Error())
	}
}


