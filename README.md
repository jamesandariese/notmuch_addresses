# notmuch\_addresses

## Intro

Parses and prints back addresses for use with notmuch.  Saves addresses in a sqlite3 database at `~/.notmuch_addresses.sqlite3`.

Written with golang for epic speeds and also for kicks and things.

This will install to $HOME/bin/{query,save}_notmuch_addresses by default.

## Usage

I use this by adding the following to my .emacs:

    ; this can be done with M-x customize-variable notmuch-address-command
    ; (and is actually how I do it)
    (custom-set-variables
          '(notmuch-address-command "~/bin/query_notmuch_addresses"))

    ; this goes AFTER setting notmuch-address-command
    (notmuch-address-message-insinuate)

In order to save addresses, it will be necessary to grab them from incoming emails.
This is commonly done in a notmuch post-new hook.  For example, mine is located at `~/Mail/.notmuch/hooks/post-new`.
`save_notmuch_addresses` takes a list of files on the command line and gathers addresses from those files into the database for later
retrieval by query.  Here's mine:

    #!/bin/sh
    # ~/Mail/.notmuch/hooks/post-new

    notmuch search --format=text0 --output=files tag:new | xargs -0 ~/bin/save_notmuch_addresses
    
    notmuch tag -new -- tag:new

As well as the config to mark new mail to be picked up by this hook:

    #~/.notmuch-config

    ...
    [new]
    tags=unread;inbox;new;
    ...

The tags line above will likely only contain `unread;inbox;`.  You add
new so all new messages have a new tag and remove new once the new
messages have been processed by the post-new hook.  Woohoo.

## Building and installing

Build:

    make

Install:

    make install

Uninstall:

    make uninstall

Clean:

    make clean