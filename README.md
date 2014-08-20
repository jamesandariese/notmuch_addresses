# notmuch\_addresses

Parses and prints back addresses for use with notmuch.

Written with golang for epic speeds and also for kicks and things.

I use this by adding the following to my .emacs:

    ; this can be done with M-x customize-variable notmuch-address-command
    ; (and is actually how I do it)
    (custom-set-variables
          '(notmuch-address-command "~/bin/query_notmuch_addresses"))

    ; this goes AFTER setting notmuch-address-command
    (notmuch-address-message-insinuate)

Build:

    make

Install:

    make install

Uninstall:

    make uninstall

Clean:

    make clean