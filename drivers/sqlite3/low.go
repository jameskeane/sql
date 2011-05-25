// Copyright 2009 Peter H. Froehlich. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlite3

/*
#include <stdlib.h>
#include <sqlite3.h>

// needed since sqlite3_column_text() and sqlite3_column_name()
// return const unsigned char* for some wack-a-doodle reason
const char *wsq_column_text(sqlite3_stmt *statement, int column)
{
	return (const char *) sqlite3_column_text(statement, column);
}
const char *wsq_column_name(sqlite3_stmt *statement, int column)
{
        return (const char *) sqlite3_column_name(statement, column);
}
int wsq_sqlite3_prepare_v2(sqlite3 *db, const char *zSql, int nByte, sqlite3_stmt **ppStmt, char **pzTail) 
{
	return sqlite3_prepare_v2(db, zSql, nByte, ppStmt, (const char**)pzTail);
}



// needed to work around the void(*)(void*) callback that is the
// last argument to sqlite3_bind_text(); SQLITE_TRANSIENT forces
// SQLite to make a private copy of the data
int wsq_bind_text(sqlite3_stmt *statement, int i, const char* text, int n)
{
	return sqlite3_bind_text(statement, i, text, n, SQLITE_TRANSIENT);
}

// needed to work around the ... argument of sqlite3_config(); if
// we ever require an option with parameters, we'll have to add more
// wrappers
int wsq_config(int option)
{
	return sqlite3_config(option);
}
*/
import "C"
import "unsafe"

// The type codes returned by sqlite3_column_type().
const (
	_ = iota;
	sqlIntegerType;
	sqlFloatType;
	sqlTextType;
	sqlBlobType;
	sqlNullType;
)

// These constants can be or'd together and passed as the
// "flags" option to Open(). Some of them only apply if
// the "vfs" option is also passed. See SQLite documentation
// for details. Note that we always force OpenFullMutex,
// so passing OpenNoMutex has no effect. See also FlagURL().
const (
	OpenReadOnly		= 0x00000001;
	OpenReadWrite		= 0x00000002;
	OpenCreate		= 0x00000004;
	OpenDeleteOnClose	= 0x00000008;	// VFS only
	OpenExclusive		= 0x00000010;	// VFS only
	OpenMainDb		= 0x00000100;	// VFS only
	OpenTempDb		= 0x00000200;	// VFS only
	OpenTransientDb		= 0x00000400;	// VFS only
	OpenMainJournal		= 0x00000800;	// VFS only
	OpenTempJournal		= 0x00001000;	// VFS only
	OpenSubJournal		= 0x00002000;	// VFS only
	OpenMasterJournal	= 0x00004000;	// VFS only
	OpenNoMutex		= 0x00008000;
	OpenFullMutex		= 0x00010000;
	OpenSharedCache		= 0x00020000;
	OpenPrivateCache	= 0x00040000;
)

// If something goes wrong on this level, we simply bomb
// out, there's no use trying to recover; note that most
// calls to sqlPanic() are for things that can never,
// ever, ever happen anyway. For regular "errors" status
// codes are returned.

func sqlPanic(str string) {
	panic("sqlite3 fatal error: " + str + "!")
}

// Wrappers around the most important SQLite types.

type sqlConnection struct {
	handle *C.sqlite3;
}

type sqlStatement struct {
	handle *C.sqlite3_stmt;
}

type sqlValue struct {
	handle *C.sqlite3_value;
}

type sqlBlob struct {
	handle *C.sqlite3_blob;
}

// Wrappers around the most important SQLite functions.

func sqlConfig(option int) int {
	return int(C.wsq_config(C.int(option)));
}

func sqlVersion() string {
	cp := C.sqlite3_libversion();
	if cp == nil {
		// The call can't really fail since it returns
		// a string constant, but let's be safe...
		sqlPanic("can't get library version");
	}
	return C.GoString(cp);
}

func sqlVersionNumber() int {
	return int(C.sqlite3_libversion_number());
}

func sqlSourceId() string {
	// SQLite 3.6.18 introduced sqlite3_sourceid(), see
	// http://www.hwaci.com/sw/sqlite/changes.html for
	// details; we can't expect wide availability yet,
	// for example Debian Lenny ships SQLite 3.5.9 only.
	if sqlVersionNumber() < 3006018 {
		return "unknown source id";
	}

	cp := C.sqlite3_sourceid();
	if cp == nil {
		// The call can't really fail since it returns
		// a string constant, but let's be safe...
		sqlPanic("can't get library sourceid");
	}
	return C.GoString(cp);
}

func sqlOpen(name string, flags int, vfs string) (conn *sqlConnection, rc int) {
	conn = new(sqlConnection);

	p := C.CString(name);
	if len(vfs) > 0 {
		q := C.CString(vfs);
		rc = int(C.sqlite3_open_v2(p, &conn.handle, C.int(flags), q));
		C.free(unsafe.Pointer(q));
	} else {
		rc = int(C.sqlite3_open_v2(p, &conn.handle, C.int(flags), nil))
	}
	C.free(unsafe.Pointer(p));

	// We could get a handle even if there's an error, see
	// http://www.sqlite.org/c3ref/open.html for details.
	// Initially we didn't want to return a connection on
	// error, but we actually have to since we want to fill
	// in a SystemError struct. Sigh.
//	if rc != StatusOk && conn.handle != nil {
//		_ = conn.sqlClose();
//		conn = nil;
//	}

	return;
}

// Wrappers as connection methods.

func (self *sqlConnection) sqlClose() int {
	return int(C.sqlite3_close(self.handle));
}

func (self *sqlConnection) sqlChanges() int {
	return int(C.sqlite3_changes(self.handle));
}

func (self *sqlConnection) sqlLastInsertRowId() int64 {
	return int64(C.sqlite3_last_insert_rowid(self.handle));
}

func (self *sqlConnection) sqlBusyTimeout(milliseconds int) int {
	return int(C.sqlite3_busy_timeout(self.handle, C.int(milliseconds)));
}

func (self *sqlConnection) sqlExtendedResultCodes(on bool) int {
	v := map[bool]int{true: 1, false: 0}[on];
	return int(C.sqlite3_extended_result_codes(self.handle, C.int(v)));
}

func (self *sqlConnection) sqlErrorMessage() string {
	cp := C.sqlite3_errmsg(self.handle);
	if cp == nil {
		// The call can't really fail since it returns
		// a string constant, but let's be safe...
		sqlPanic("can't get error message");
	}
	return C.GoString(cp);
}

func (self *sqlConnection) sqlErrorCode() int {
	return int(C.sqlite3_errcode(self.handle));
}

func (self *sqlConnection) sqlExtendedErrorCode() int {
	// SQLite 3.6.5 introduced sqlite3_extended_errcode(),
	// see http://www.hwaci.com/sw/sqlite/changes.html for
	// details; we can't expect wide availability yet, for
	// example Debian Lenny ships SQLite 3.5.9 only.
	if sqlVersionNumber() < 3006005 {
		// just return the regular error code...
		return self.sqlErrorCode();
	}
	return int(C.sqlite3_extended_errcode(self.handle));
}

func (self *sqlConnection) sqlPrepare(query string) (stat *sqlStatement, rc int) {
	stat = new(sqlStatement);

	p := C.CString(query);
	// TODO: may need tail to process statement sequence? or at
	// least to generate an error that we missed some SQL?
	//
	// -1: process query until 0 byte
	// nil: don't return tail pointer
	rc = int(C.wsq_sqlite3_prepare_v2(self.handle, p, -1, &stat.handle, nil));
	C.free(unsafe.Pointer(p));

	// We are not supposed to get a handle on error. Since
	// sqlite3_open() follows a different rule, however, we
	// indulge in paranoia and check to make sure. We really
	// don't want to return a statement on error.
	if rc != StatusOk && stat.handle != nil {
		_ = stat.sqlFinalize();
		stat = nil;
	}

	return;
}

// Wrappers as statement methods.

func (self *sqlStatement) sqlBindParameterCount() int {
	return int(C.sqlite3_bind_parameter_count(self.handle));
}

func (self *sqlStatement) sqlBindText(slot int, value string) int {
	p := C.CString(value);
	// SQLite counts slots from 1 instead of 0; -1 means "until
	// end of string" here.
	rc := int(C.wsq_bind_text(self.handle, C.int(slot+1), p, C.int(-1)));
	C.free(unsafe.Pointer(p));
	return rc;
}

func (self *sqlStatement) sqlBindInt(slot int, value int) int {
	return int(C.sqlite3_bind_int(self.handle, C.int(slot+1), C.int(value)))
}

func (self *sqlStatement) sqlBindNull(slot int) int {
	return int(C.sqlite3_bind_null(self.handle, C.int(slot+1)))
}

func (self *sqlStatement) sqlClearBindings() int {
	return int(C.sqlite3_clear_bindings(self.handle));
}

func (self *sqlStatement) sqlStep() int {
	return int(C.sqlite3_step(self.handle));
}

func (self *sqlStatement) sqlSql() string {
	cp := C.sqlite3_sql(self.handle);
	if cp == nil {
		// The call shouldn't fail unless we forgot to
		// use sqlite3_prepare_v2()...
		sqlPanic("can't get SQL statement");
	}
	return C.GoString(cp);
}

func (self *sqlStatement) sqlFinalize() int {
	return int(C.sqlite3_finalize(self.handle));
}

func (self *sqlStatement) sqlReset() int {
	return int(C.sqlite3_reset(self.handle));
}

func (self *sqlStatement) sqlColumnCount() int {
	return int(C.sqlite3_column_count(self.handle));
}

func (self *sqlStatement) sqlColumnType(col int) int {
	return int(C.sqlite3_column_type(self.handle, C.int(col)));
}

func (self *sqlStatement) sqlColumnName(col int) string {
	cp := C.wsq_column_name(self.handle, C.int(col));
	if cp == nil {
		// TODO: not sure at all when and how this can
		// fail...
		sqlPanic("can't get column name");
	}
	return C.GoString(cp);
}

func (self *sqlStatement) sqlColumnText(col int) string {
	cp := C.wsq_column_text(self.handle, C.int(col));
	// Apparently this can return nil, for example if there
	// is no value in the column. So we can't sanity check
	// anything here...
//	if cp == nil {
//		sqlPanic("can't get column text");
//	}
	return C.GoString(cp);
}

func (self *sqlStatement) sqlColumnDeclaredType(col int) string {
	cp := C.sqlite3_column_decltype(self.handle, C.int(col));
	// This can return nil, for example if the column is an
	// SQL expression and not a "real" column in a table. So
	// again no sanity checks...
	return C.GoString(cp);
}
