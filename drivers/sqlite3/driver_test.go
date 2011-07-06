package sqlite3

import (
	"testing"
	"sql"
)


func TestConnect(t *testing.T) {
	con1, err := sql.Connect("sqlite3:_test/2/hello.db")
	if err != nil {
		t.Error(err)
	} else {
	
		con1.Close()
	}
}
