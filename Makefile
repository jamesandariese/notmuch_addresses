PREFIX=$(HOME)/bin

ALL: bin/query_notmuch_addresses bin/save_notmuch_addresses

bin:
	mkdir bin

bin/save_notmuch_addresses: bin save_notmuch_addresses/save_notmuch_addresses
	cp -a save_notmuch_addresses/save_notmuch_addresses bin/save_notmuch_addresses

bin/query_notmuch_addresses: bin query_notmuch_addresses/query_notmuch_addresses
	cp -a query_notmuch_addresses/query_notmuch_addresses bin/query_notmuch_addresses

save_notmuch_addresses/save_notmuch_addresses:
	(cd save_notmuch_addresses && go build)

query_notmuch_addresses/query_notmuch_addresses:
	(cd query_notmuch_addresses && go build)

clean:
	go clean
	find . -name '*~' -exec rm {} \;
	rm -f bin/save_notmuch_addresses
	rm -f bin/query_notmuch_addresses
	[ -d bin ] && rmdir bin
	rm -f query_notmuch_addresses/query_notmuch_addresses
	rm -f save_notmuch_addresses/save_notmuch_addresses

install: bin/query_notmuch_addresses bin/save_notmuch_addresses
	mkdir -p $(PREFIX)
	cp -a bin/query_notmuch_addresses bin/save_notmuch_addresses $(PREFIX)

uninstall:
	rm -f $(PREFIX)/query_notmuch_addresses $(PREFIX)/save_notmuch_addresses
