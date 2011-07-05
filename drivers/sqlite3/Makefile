include $(GOROOT)/src/Make.inc

TARG=sql/sqlite3
GOFILES=connection.go conversion.go driver.go error.go result_set.go statement.go

CGOFILES=low.go
CGO_LDFLAGS=-lsqlite3

# this is the local install
PKGROOT=../..

# compile/link against local install
GC+= -I $(PKGROOT)/_obj
export GL=$(PKGROOT)/gotest_link_hack.sh
LD+= -L $(PKGROOT)/_obj
	
#  copy to local install
$(PKGROOT)/_obj/$(TARG).a: _obj/$(TARG).a
	mkdir -p $(dir $@); cp -f $< $@
package: $(PKGROOT)/_obj/$(TARG).a

include $(GOROOT)/src/Make.pkg

# local dependencies
_obj/$(TARG).a: $(PKGROOT)/_obj/sql.a
