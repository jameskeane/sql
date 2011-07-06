package postgresql

import (
	"sql"
	"strings"
	"fmt"
	"os"
)

type Driver struct {
}


func parseDataSource(dsn *http.URL) (string) {

	opt := ""
	for k, v := range dsn.Parameters {
		opt += fmt.Sprintf("%s=%s ", v, v)
	}

    return fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s %s", dsn.Host, "", dsn.Database, dsn.User, dsn.Password, opt), nil
}


// connect string:
// 'postgresql://user:pwd@host:port/dbname?opt1=value
// for addition options see: http://www.postgresql.org/docs/8.1/static/libpq.html#LIBPQ-CONNECT
func (self *Driver) Connect(dsn *sql.DSN) (sql.Connection, sql.Error) {
    ds, err := parseDataSource(dsn)
    if err != nil {
        return nil, sql.NewError(err.String())
    }

	fmt.Println(ds)

    sqlConn := pqConnect(ds)
    if sqlConn == nil {
    	//TODO: error checking
		return nil, sql.Busy
    }
    
    return nil, nil
}


func init() {
	sql.RegisterDriver("postgresql", new(Driver))
}
