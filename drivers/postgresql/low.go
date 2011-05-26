package postgresql

/*
#include <stdlib.h>
#include <libpq-fe.h>
*/
import "C"
import (
	"unsafe"
	"os"
	"strconv"
)

// If something goes wrong on this level, we simply bomb
// out, there's no use trying to recover; note that most
// calls to sqlPanic() are for things that can never,
// ever, ever happen anyway. For regular "errors" status
// codes are returned.

func pqPanic(str string) {
	panic("postgresql fatal error: " + str + "!")
}

// Wrappers around the most important postgresql types.
type pqConnection struct {
	handle *C.PGconn
}

type pqStatement struct {
	// postgresql keeps track of prepared statements by name ? weird I know
	// easiest way is to generate some random string
	stmtName string
}

func pqConnect(conninfo string) (conn *pqConnection) {
	conn = new(pqConnection);
	
	// Prepare the parameter for calling into libpq
	p := C.CString(conninfo)
	
	// Try to connect
	conn.handle = C.PQconnectdb(p)
	
	// Free the parameter
	C.free(unsafe.Pointer(p))
	
	return conn
}

func (self *pqConnection) pqClose() {
	C.PQfinish(self.handle)
	// ensure we don't use this handle again
	self.handle = nil
}

func (self *pqConnection) pqPrepare(query string, types ...interface{}) *pqStatement {
	res := new(pqStatement)
	// TODO: is the current system time goog enough?
	secs, _, _ := os.Time()
	res.stmtName = strconv.Itoa64(secs)
	
	nameC := C.CString(res.stmtName)
	queryC := C.CString(query)
	
	// TODO: error check here :| yikes
	_ = C.PQprepare(self.handle, nameC, queryC, 0, nil)
	
	C.free(unsafe.Pointer(nameC))
	C.free(unsafe.Pointer(queryC))
	return res
}
