package sqlite3

import (
	"sql"
)

type Statement struct {
	handle *sqlStatement
	connection	*Connection;
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
