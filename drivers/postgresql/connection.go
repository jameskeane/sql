package postgresql

import (
	"sql"
)

type Connection struct {
	handle *pqConnection;
}

func (self *Connection) Query(sql string, params ...interface{}) (sql.ResultSet, sql.Error) {
	stmt, err := self.Prepare(sql)
	if err != nil {
		return nil, err	
	}

	// TODO: need a way to finalize this statement ? but allow the result set to stay open
	//defer stmt.Close()
	return stmt.Query(params...)
}

func (self *Connection) Execute(sql string, params ...interface{}) sql.Error {
	stmt, err := self.Prepare(sql)
	if err != nil {
		return err	
	}

	defer stmt.Close()
	return stmt.Execute(params...)
}

func (self *Connection) Prepare(query string) (sql.Statement, sql.Error) {
	sqlStmt := self.handle.pqPrepare(query)
	// TODO: Error checking
	

	stmt := new(Statement)
	stmt.handle = sqlStmt
	stmt.connection = self
	return stmt, nil
}

func (self *Connection) Close() sql.Error {
	self.handle.pqClose()
	return nil
}

