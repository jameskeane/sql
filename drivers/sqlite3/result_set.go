package sqlite3


type ResultSet struct {
	statement *Statement
	connection *Connection
}


func (self *ResultSet) RowCount() (uint64, sql.Error){
	return nil, nil
}

func (self *ResultSet) Next bool {

}

// Scan the current row of results by column index.
func (self *ResultSet) Scan(refs ...interface{}) sql.Error {

}

// Scan the current row of results by column name.
func (self *ResultSet) NamedScan(refs ...interface{}) sq.Error {

}

// Release the resources held by a ResultSet.
func (self *ResultSet) Close() sql.Error {

}
