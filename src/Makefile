include $(GOROOT)/src/Make.inc

TARG=sql
GOFILES=connection.go driver.go error.go manager.go result_set.go statement.go dsn.go

# gb: this is the local install
PKGROOT=..

# gb: compile/link against local install
GC+= -I $(PKGROOT)/_obj
LD+= -L $(PKGROOT)/_obj

# gb: copy to local install
$(PKGROOT)/_obj/$(TARG).a: _obj/$(TARG).a
	mkdir -p $(dir $@); cp -f $< $@
package: $(PKGROOT)/_obj/$(TARG).a

include $(GOROOT)/src/Make.pkg
