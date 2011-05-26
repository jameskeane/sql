package sqlite3

import (
	"sql"
)

type Statement struct {
	handle *sqlStatement
	connection	*Connection;
}

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



func (self *Statement) Query(params ...interface{}) (sql.ResultSet, sql.Error)  {

	rc := self.handle.sqlReset()
	if rc != StatusOk {
		// Then there is some error
		return nil, sql.NewError(self.connection.handle.sqlErrorMessage())	
	}
	
	// TODO: Error Checking
	self.handle.sqlClearBindings()


	err := self.handle.BindParams(params...)
	if err != nil {
		return nil, err
	}
	
	rs := new(ResultSet)
	rs.statement = self
	rs.connection = self.connection
	
	return rs, nil
}


func (self *Statement) Execute(params ...interface{}) sql.Error {
	err := self.handle.BindParams(params...)
	if err != nil {
		return err
	}

	rc := self.handle.sqlStep() 
	if rc != StatusDone && rc != StatusRow {
		// Then there is some error
		return sql.NewError(self.connection.handle.sqlErrorMessage())	
	}
	
	rc = self.handle.sqlReset()
	if rc != StatusOk {
		// Then there is some error
		return sql.NewError(self.connection.handle.sqlErrorMessage())	
	}
	
	// TODO: Error Checking
	self.handle.sqlClearBindings()
	return nil
}

func (self *Statement) Close() sql.Error {
	rc := self.handle.sqlFinalize()
	if rc != StatusOk {
		return sql.NewError(self.connection.handle.sqlErrorMessage())	
	}
	return nil
}
