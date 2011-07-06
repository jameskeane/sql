package sql

import (
	"testing"
	"os"
	"reflect"
)

type DSNTest struct {
	in string
	dsn *DSN
	err  os.Error
}

var dsn_tests []DSNTest = []DSNTest {
	// bad inputs
	{"", nil, InvalidDSN},
	{"nocolon", nil, InvalidDSN},
	{"driver:", nil, InvalidDSN},
	{"driver:@", nil, InvalidDSN},
	{"driver:host:badport", nil, InvalidDSN},
	{"driver:user:password@host:30/?p", nil, InvalidDSN},
	{"driver:user:password@host:30/?p=", nil, InvalidDSN},
	{"driver:user:password@host:30/?p=&p2=value", nil, InvalidDSN},
	{"driver:user:password@host:30/?p=true&p=false", nil, InvalidDSN},
	
	// good inputs
	{"driver:test.db", &DSN{Driver:"driver", Host:"test.db", Parameters:map[string] string{}}, nil},
	{"driver:user:pass@host:50", &DSN{Driver:"driver", User:"user", Password:"pass", Host:"host", Port:50, Parameters:map[string] string{}}, nil},
	{"sqlite3:../_test/test.db/dbname?p1=1&p2=2", &DSN{Driver:"sqlite3", Host:"../_test/test.db", Database:"dbname", Parameters:map[string] string{"p1":"1","p2":"2"}}, nil},
	{"driver:host?p=1", &DSN{Driver:"driver", Host:"host", Parameters:map[string] string{"p":"1"}}, nil},
}

func TestDSN(t *testing.T) {
	for i, test := range dsn_tests {
		ret, err := parseDSN(test.in)
		if err != test.err {
			t.Error(err)
		}
		
		if !reflect.DeepEqual(ret, test.dsn) {
			t.Log("Test %d does not match:", i)
			t.Log("\treturned: %v", ret)
			t.Log("\texpected: %v", test.dsn)
			t.Fail()
		}
	}
}
