package sqlite3

import (
	"http"
	"sql"
	"fmt"
)

type Driver struct {
}

type DataSource struct {
	file string
	flags int
	vfs string
}

func parseDataSource(dsn *http.URL) (res *DataSource, err sql.Error) {
    res = new(DataSource)
    res.file = dsn.Host
    
    // TODO: figure out a good way to pass these options, probably in the query
    res.flags = OpenReadWrite | OpenCreate
    res.vfs = ""
    
    //res.options, err = http.ParseQuery(dsn.RawQuery)
    return
}

func (self *Driver) Connect(dsn *http.URL) (sql.Connection, sql.Error) {
    ds, err := parseDataSource(dsn)
    if err != nil {
        return nil, err
    }    
	
    sqlConn, rc := sqlOpen(ds.file, ds.flags, ds.vfs)
    if rc != StatusOk {
    	//TODO: error checking
		return nil, sql.Busy
    }
    
    conn := new(Connection)
    conn.handle = sqlConn
    
    return conn, nil
}


func init() {
	sql.RegisterDriver("sqlite3", new(Driver))
}
