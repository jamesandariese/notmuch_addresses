package notmuch_addresses

import (
	"errors"
	"net/mail"
	"os"
	"code.google.com/p/go-sqlite/go1/sqlite3"
	"log"
	"io"
	"fmt"
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

var (
	conn *sqlite3.Conn
)

func InitDatabase(path string) error {
	c, err := sqlite3.Open(path)
	if err != nil {
		return err
	}
	conn = c

	conn.Exec(`CREATE TABLE IF NOT EXISTS address (
                       address TEXT,
                       raw TEXT,
                       count INTEGER
                   );`)

	return nil
}

func Close() (err error) {
	err = conn.Close()
	conn = nil
	return
}

var ErrDatabaseNotOpen = errors.New("Database is not open")

//func GatherAddresses(path string) (addresses []*mail.Address, err error) {
func GatherAddresses(path string) (addresses int, err error) {
	if conn == nil {
		return 0, ErrDatabaseNotOpen
	}

	file, err := os.Open(path)
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
			for _, address := range tmp_addresses {
				//addresses = append(addresses, tmp_addresses...)
				for {
					var save bool
					conn.Begin()
					old_affected := conn.TotalRowsAffected()
					err = conn.Exec(`UPDATE address
                                                            SET count = count + 1
                                                          WHERE raw = ?`, 
						address.String())
					if err != nil {
						// Can't save this email address.
						// It wasn't because of a failed transaction.
						log.Print("Couldn't increment or save address", address, err)
						break
					}
					if conn.TotalRowsAffected() == old_affected {
						err = conn.Exec(`INSERT INTO address (address, raw, count) VALUES (?,?,1);`,
							address.Address,
							address.String())
						if err != nil {
							log.Print("Couldn't save address", address, err)
							break
						}
						save = true
					} else {
						save = false
					}
					if conn.Commit() == nil {
						if save {
							log.Print("Saved new address ", address)
						} else {
							log.Print("Incremented address ", address)
						}
						break
					} else {
						log.Print("Couldn't commit transaction.  Retrying.")
					}
				}
			}
			addresses += 1
		}
	}
	return
}

func QueryToStdout(substring string) error {
	if conn == nil {
		return ErrDatabaseNotOpen
	}
	
	stmt, err := conn.Query(`SELECT raw FROM address WHERE raw LIKE ?;`, "%" + substring + "%")
	if err != nil {
		if err == io.EOF {
			return nil
		}
		return err
	}
	for ; err == nil; err = stmt.Next() {
		var raw string;
		errb := stmt.Scan(&raw)
		if errb != nil {
			log.Print("Couldn't scan row from SQLite3: ", err)
		} else {
			fmt.Printf("%s\n", raw)
		}
	}
	if err == io.EOF {
		return nil
	}
	return err
}
