package postgresql

import (
	"http"
	"sql"
	"strings"
	"fmt"
	"os"
)

type Driver struct {
}


func parseDataSource(dsn *http.URL) (string, os.Error) {
    // Grab the basic stuff
    host := dsn.Host
    db := dsn.Path[1:]
    
    // Parse the authentication info
    var user, pwd string
    authinfo := strings.Split(dsn.RawUserinfo, ":", 2)
    if len(authinfo) >= 1 {
        user = authinfo[0]
        if len(authinfo) >= 2 {
            pwd = authinfo[1]
        }
    }
   
    // Parse the options and build the connect string
    options, err := http.ParseQuery(dsn.RawQuery)
	if err != nil {
		return "", err
	}
	
	opt := ""
	for key, val := range options {
		opt += fmt.Sprintf("%s=%s ", key, val[0])
	}

    return fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s %s", host, "", db, user, pwd, opt), nil
}


// connect string:
// 'postgresql://user:pwd@host:port/dbname?opt1=value
// for addition options see: http://www.postgresql.org/docs/8.1/static/libpq.html#LIBPQ-CONNECT
func (self *Driver) Connect(dsn *http.URL) (sql.Connection, sql.Error) {
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
