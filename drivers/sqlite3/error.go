// Copyright 2009 Peter H. Froehlich. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlite3

import "fmt"

// Error in the database driver itself, *not* the database
// system we talk to.
type DriverError struct {
	message string;
}

// Textual description of the error.
// Implements os.Error interface.
func (self DriverError) String() string	{ return self.message }

// Basic SQLite status codes as returned by almost
// every SQLite operation. Note that SQLite calls
// these "result codes" but we use "Result" for a
// higher-level concept already.
const (
	StatusOk		= iota;	// Successful result
	StatusError;		// SQL error or missing database
	StatusInternal;		// Internal logic error in SQLite
	StatusPerm;		// Access permission denied
	StatusAbort;		// Callback routine requested an abort
	StatusBusy;		// The database file is locked
	StatusLocked;		// A table in the database is locked
	StatusNoMem;		// A malloc() failed
	StatusReadOnly;		// Attempt to write a readonly database
	StatusInterrupt;	// Operation terminated by sqlite3_interrupt()
	StatusIoErr;		// Some kind of disk I/O error occurred
	StatusCorrupt;		// The database disk image is malformed
	StatusNotFound;		// NOT USED. Table or record not found
	StatusFull;		// Insertion failed because database is full
	StatusCantOpen;		// Unable to open the database file
	StatusProtocol;		// NOT USED. Database lock protocol error
	StatusEmpty;		// Database is empty
	StatusSchema;		// The database schema changed
	StatusTooBig;		// String or BLOB exceeds size limit
	StatusConstraint;	// Abort due to constraint violation
	StatusMismatch;		// Data type mismatch
	StatusMisuse;		// Library used incorrectly
	StatusNoLfs;		// Uses OS features not supported on host
	StatusAuth;		// Authorization denied
	StatusFormat;		// Auxiliary database format error
	StatusRange;		// 2nd parameter to sqlite3_bind out of range
	StatusNotADb;		// File opened that is not a database file
	StatusRow		= 100;	// sqlite3_step() has another row ready
	StatusDone		= 101;	// sqlite3_step() has finished executing
)

// Extended SQLite status codes for StatusIoErr.
// Provide additional information on top of basic status codes.
const (
	_				= iota;
	StatusIoErrRead			= StatusIoErr | (iota << 8);
	StatusIoErrShortRead		= StatusIoErr | (iota << 8);
	StatusIoErrWrite		= StatusIoErr | (iota << 8);
	StatusIoErrFSync		= StatusIoErr | (iota << 8);
	StatusIoErrDirFSync		= StatusIoErr | (iota << 8);
	StatusIoErrTruncate		= StatusIoErr | (iota << 8);
	StatusIoErrFStat		= StatusIoErr | (iota << 8);
	StatusIoErrUnlock		= StatusIoErr | (iota << 8);
	StatusIoErrRdlock		= StatusIoErr | (iota << 8);
	StatusIoErrDelete		= StatusIoErr | (iota << 8);
	StatusIoErrBlocked		= StatusIoErr | (iota << 8);
	StatusIoErrNoMem		= StatusIoErr | (iota << 8);
	StatusIoErrAccess		= StatusIoErr | (iota << 8);
	StatusIoErrCheckReservedBlock	= StatusIoErr | (iota << 8);
	StatusIoErrLock			= StatusIoErr | (iota << 8);
	StatusIoErrClose		= StatusIoErr | (iota << 8);
	StatusIoErrDirClose		= StatusIoErr | (iota << 8);
)

// Extended SQLite status codes for StatusLocked.
// Provide additional information on top of basic status codes.
const (
	_			= iota;
	StatusLockedSharedCache	= StatusLocked | (iota << 8);
)

// Error in the database system we talk to.
// SQLite has basic and extended status codes
// in addition to textual messages.
type SystemError struct {
	message		string;
	basic		int;
	extended	int;
}

// Textual description of the error.
// Implements os.Error interface.
func (self SystemError) String() string {
	return fmt.Sprintf("%s (%d:%d)", self.message, self.basic, self.extended)
}

// Basic SQLite status code. These are plain
// integers.
func (self SystemError) Basic() int	{ return self.basic }

// Extended SQLite status code. These are OR'd
// together from various bits and pieces on top
// of basic status codes.
func (self SystemError) Extended() int	{ return self.extended }
