include $(GOROOT)/src/Make.inc

TARG=sql/postgresql
CGOFILES=low.go
GOFILES=driver.go
CGO_CFLAGS=-I$(shell pg_config --includedir)
CGO_LDFLAGS=-L$(shell pg_config --libdir) -lpq
#CLEANFILES+=example test.db

include $(GOROOT)/src/Make.pkg

#example: install test.db example.go
#	$(GC) example.go
#	$(LD) -o $@ example.$O

#test.db:
#	sqlite3 test.db <create_db.sql
