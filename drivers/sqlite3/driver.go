package sqlite3

import (
	"sql"
)

type Driver struct {
}


func parseDataSource(dsn *sql.DSN) (file string, flags int, vfs string) {
    // TODO: figure out a good way to pass the flags, probably in the query
    //			sqlite3 supports uri filenames <http://www.sqlite.org/uri.html>
    return dsn.Host+"/"+dsn.Database, OpenReadWrite | OpenCreate, dsn.Parameters["vfs"]
}

func (self *Driver) Connect(dsn *sql.DSN) (sql.Connection, sql.Error) {
    file, flags, vfs := parseDataSource(dsn)
	
    sqlConn, rc := sqlOpen(file, flags, vfs)
    if rc != StatusOk {
    	//TODO: error checking
		return nil, sql.Busy
    }
    
    return &Connection{handle:sqlConn}, nil
}


func init() {
	sql.RegisterDriver("sqlite3", new(Driver))
}
