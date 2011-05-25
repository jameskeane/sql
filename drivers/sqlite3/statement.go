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
			default:
				// TODO: Error checking
				return sql.InvalidBindType
		}
		
		if rc != StatusOk {
			// TODO: Error checking
			return sql.InvalidBindType
		}	
	}
	return nil
}



func (self *Statement) Query(params ...interface{}) (sql.ResultSet, sql.Error)  {
	err := self.handle.BindParams(params...)
	if err != nil {
		return err
	}

	// TODO: Implement this
	self.handle.sqlReset()
	self.handle.sqlClearBindings()
		
	return nil, nil
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
