package postgresql

import (
	"sql"
)

type Statement struct {
}

func (self *Statement) Query(params ...interface{}) (sql.ResultSet, sql.Error)  {
	return nil, nil
}


func (self *Statement) Execute(params ...interface{}) sql.Error {
	return nil
}

func (self *Statement) Close() sql.Error {
	return nil
}
