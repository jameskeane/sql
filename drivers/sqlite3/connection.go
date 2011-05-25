package sqlite3

import (
	"sql"
)

type Connection struct {
	handle *sqlConnection;
}

func (self *Connection) Query(sql string, params ...interface{}) (sql.ResultSet, sql.Error) {
	stmt, err := self.Prepare(sql)
	if err != nil {
		return nil, err	
	}

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
	sqlStmt, rc := self.handle.sqlPrepare(query)
	if rc != StatusOk {
		return nil, sql.NewError(self.handle.sqlErrorMessage())	
	}

	stmt := new(Statement)
	stmt.handle = sqlStmt
	stmt.connection = self
	return stmt, nil
}

func (self *Connection) Close() sql.Error {
	rc := self.handle.sqlClose()
	if rc != StatusOk {
		return sql.NewError(self.handle.sqlErrorMessage())	
	}
	return nil
}

