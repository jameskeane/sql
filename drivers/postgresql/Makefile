include $(GOROOT)/src/Make.inc

TARG=sql/postgresql
GOFILES=driver.go connection.go statement.go
 
CGOFILES=low.go
CGO_CFLAGS=-I$(shell pg_config --includedir)
CGO_LDFLAGS=-L$(shell pg_config --libdir) -lpq

# this is the local install
PKGROOT=../..

# compile/link against local install
GC+= -I $(PKGROOT)/_obj
export GL=$(PKGROOT)/gotest_link_hack.sh
LD+= -L $(PKGROOT)/_obj

# copy to local install
$(PKGROOT)/_obj/$(TARG).a: _obj/$(TARG).a
	mkdir -p $(dir $@); cp -f $< $@
package: $(PKGROOT)/_obj/$(TARG).a

include $(GOROOT)/src/Make.pkg

# local dependencies
_obj/$(TARG).a: $(PKGROOT)/_obj/sql.a
