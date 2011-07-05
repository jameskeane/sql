include $(GOROOT)/src/Make.inc

.PHONY: all install test clean nuke

all:
	gomake -C src
	gomake -C drivers/sqlite3
	gomake -C drivers/postgresql

install: all
	gomake -C src install
	gomake -C drivers/sqlite3 install
	gomake -C drivers/postgresql install

test:
	gomake -C src
	gomake -C drivers/sqlite3 test
	gomake -C drivers/postgresql test

clean:
	gomake -C src clean
	gomake -C drivers/sqlite3 clean
	gomake -C drivers/postgresql clean
	rm -rf _obj
nuke:
	gomake -C src nuke
	gomake -C drivers/sqlite3 nuke
	gomake -C drivers/postgresql nuke
