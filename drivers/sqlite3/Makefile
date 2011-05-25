# mostly copied from Eden Li's mysql interface
# "Who is supposed to grok this mess?" --- phf

include $(GOROOT)/src/Make.inc

TARG=sql/sqlite3
CGOFILES=low.go
GOFILES=driver.go error.go connection.go statement.go
CGO_LDFLAGS=-lsqlite3
#CLEANFILES+=example test.db

include $(GOROOT)/src/Make.pkg

#example: install test.db example.go
#	$(GC) example.go
#	$(LD) -o $@ example.$O

#test.db:
#	sqlite3 test.db <create_db.sql
