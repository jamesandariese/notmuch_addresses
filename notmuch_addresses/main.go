package main

import (
	"github.com/jamesandariese/notmuch_addresses"
	//"net/mail"
	"fmt"
)

func main() {
	addresses, err := notmuch_addresses.GatherAddresses("message")
	if err != nil {
		panic("Couldn't gather email addresses: " + err.Error())
	}
	for _, address := range addresses {
		fmt.Printf("%#v\n", address.String())
	}
}


