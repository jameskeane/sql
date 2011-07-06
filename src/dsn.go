package sql

import (
	"bytes"
	"os"
	"strings"
	"strconv"
)

var (
	InvalidDSN = os.NewError("sql: Invalid DSN string")
)

// typical dsn string:
// 		'<driver>:<user>:<password>@<host>:<port>/<database>?<param>=<value>&<param>=<value>'
type DSN struct {
	Driver string
	Host string
	Port int
	User string
	Password string
	Database string
	Parameters map[string] string
}

// TODO: this should be valid 'driver:host?param=value'
func parseDSN(in string) (*DSN, os.Error) {
	var err os.Error
	var t string
	sr := bytes.NewBufferString(in)
	
	// init the dsn
	dsn := &DSN{}
	dsn.Parameters = map[string] string {}
	
	// Read the driver
	dsn.Driver, err = sr.ReadString(':')
	if err != nil {
		return nil, InvalidDSN
	}
	dsn.Driver = dsn.Driver[:len(dsn.Driver)-1]
	
	// look ahead to see if we have a user and password
	if i := strings.Index(sr.String(), "@"); i != -1 {
		// Read it this way since we already have the index we need
		temp := make([]byte, i)
		sr.Read(temp)
		if k := strings.Index(string(temp), ":"); k != -1 {
			dsn.User = string(temp[:k])
			dsn.Password = string(temp[k+1:])
		} else {
			dsn.User = string(temp)
		}
		
		// clear the '@'
		sr.ReadByte()
	}
	
	// now read the host and port and database
	t, err = sr.ReadString('?')
	if len(t) == 0 {
		return nil, InvalidDSN
	}	
	
	if err != os.EOF {
		if err == nil {
			// everything worked, but we have an extra '?' in our temp string remove it
			t = t[:len(t)-1]
		} else {
			// we got an error
			return nil, err
		}
	}
	
	// look ahead to see if theres a database name in that string
	if k := strings.LastIndex(t, "/"); k != -1 {
		dsn.Database = t[k+1:]
		t = t[:k]	
	}
	
	// look ahead to see if a port is defined
	if k := strings.Index(t, ":"); k != -1 {
		// we have a port
		dsn.Host = t[:k]
		var cerr os.Error
		dsn.Port, cerr = strconv.Atoi(t[k+1:])
		if cerr != nil {
			return nil, InvalidDSN
		}
	} else {
		dsn.Host = t
	}
	
	// try and read parameters?
	if err != os.EOF {
		for {
			t, err = sr.ReadString('&')
			if err != os.EOF {
				if err == nil {
					// everything worked, but we have an extra '&' in our temp string remove it
					t = t[:len(t)-1]
				} else {
					// we got an error
					return nil, err
				}
			}
		
			// split and add to the map
			k := strings.Index(t, "=")
			if k == -1 || k == len(t)-1 {
				return nil, InvalidDSN
			}
			
			// TODO: make sure the second part does not contain any extra '=', that would be invalid
			// Can't define the same parameter twice
			if _, ok := dsn.Parameters[t[:k]]; ok {
				return nil, InvalidDSN
			}
			
			dsn.Parameters[t[:k]] = t[k+1:]

			// terminate clause, go has no do ... while :(
			if err == os.EOF {
				break;	
			}
		}
	}
	return dsn, nil
}
