// This file handles all the conversions used to convert go's types into sqlite3 types

package sqlite3

import (
	"sql"
)

// TODO: don't use int!!it can change on each machine.. must convert to int32 or int64 first

// Wrap this on the sqlStatement type to avoid exporting it
func (self *sqlStatement) BindParams(params ...interface{}) sql.Error {
	var rc int
	
	stmtCount := self.sqlBindParameterCount()
	paramCount := len(params)
	
	if( paramCount != stmtCount ) {
		return sql.InvalidBindType
	}
	
	if( paramCount < 1 ) {
		return nil
	}
	
	for pos, param := range params {
		switch param.(type) {
			case nil:
				rc = self.sqlBindNull(pos)
			case int:
				rc = self.sqlBindInt(pos, param.(int))
			case int64:
				rc = self.sqlBindInt64(pos, param.(int64))
			case string:
				rc = self.sqlBindText(pos, param.(string))
			case float32:
				rc = self.sqlBindFloat32(pos, param.(float32))
			case float64:
				rc = self.sqlBindFloat64(pos, param.(float64))
			case bool:
				// http://www.sqlite.org/datatype3.html
				// SQLite does not have a separate Boolean storage class. Instead, Boolean values are stored as integers 0 (false) and 1 (true).
				var b int
				if param.(bool) {
					b = 1
				} else {
					b = 0
				}
				rc = self.sqlBindInt(pos, b)
			default:
				return sql.NewError("Parameter " + string(pos+1) + "is an unrecognized type")
		}
		
		if rc != StatusOk {
			return sql.NewError("Could not bind parameter " + string(pos+1))
		}	
	}
	return nil
}


func (self *ResultSet) getColumn(pos int, ptr interface{}) sql.Error {
	switch v := ptr.(type) {
		case nil:
			// Do nothing
			break
		case *int:
			*v = self.statement.handle.sqlColumnInt(pos)
		case *int64:
			*v = self.statement.handle.sqlColumnInt64(pos)
		case *float32:
			*v = self.statement.handle.sqlColumnFloat32(pos)
		case *float64:
			*v = self.statement.handle.sqlColumnFloat64(pos)
		case *string:
			*v = self.statement.handle.sqlColumnText(pos)
		default:
			return sql.NewError("Attempting to scan an unrecognized type")				
	}
	return nil
}
