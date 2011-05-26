package sqlite3

import (
	"sql"
)

type ResultSet struct {
	statement *Statement
	connection *Connection
}


func (self *ResultSet) Next() bool {

	rc := self.statement.handle.sqlStep()
	if rc == StatusDone {
		return false
	}
	if rc == StatusRow {
		return true
	}
	
	// TODO: ??
	// Some error happened
	//return sql.Error(self.connection.handle.sqlErrorMessage())
	return false
}

// Scan the current row of results by column index.
func (self *ResultSet) Scan(refs ...interface{}) sql.Error {
	// TODO: is this the best way?
	columnCount := self.statement.handle.sqlColumnCount()

	if len(refs) > columnCount {
		return sql.NewError("Trying to scan more columns than exist!")
	}
	
	for i, val := range refs {
		switch v := val.(type) {
			case nil:
				// Do nothing
				continue
			case *int:
				*v = self.statement.handle.sqlColumnInt(i)
			case *string:
				*v = self.statement.handle.sqlColumnText(i)
			default:
				return sql.NewError("Attempting to scan an unrecognized type")				
		}
	}
	return nil
}

// Scan the current row of results by column name.
func (self *ResultSet) NamedScan(refs ...interface{}) sql.Error {
	return nil
}

// Release the resources held by a ResultSet.
func (self *ResultSet) Close() sql.Error {
	// ??
	return nil
}
