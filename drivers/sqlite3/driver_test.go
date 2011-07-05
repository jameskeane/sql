package sqlite3

import (
	"testing"
	"sql"
)


func TestConnect(t *testing.T) {
	con1, err := sql.Connect("sqlite3://_test/hello.db")
	if err != nil {
		t.Fail()
	}
	
	con1.Close()
}
