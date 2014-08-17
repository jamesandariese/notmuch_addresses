package notmuch_addresses

import (
	"net/mail"
	"os"
)

var (
	headers = [...]string{
		"To",
		"Cc",
		"CC",
		"Bcc",
		"BCC",
		"From",
	}
)

func GatherAddresses(filename string) (addresses []*mail.Address, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	msg, err := mail.ReadMessage(file)
	if err != nil {
		return
	}
	for _,header := range headers {
		tmp_addresses, err := msg.Header.AddressList(header)
		if err == nil {
			addresses = append(addresses, tmp_addresses...)
		}
	}
	return 
}