package sql

type Connection interface {
	
	// Executes a query, returning a ResultSet.
	Query(sql string, params ...interface{}) (ResultSet, Error)

	// Executes a query, discarding any results.
	Execute(sql string, params ...interface{}) Error
	
	// Prepares an SQL statement for later execution.
	Prepare(sql string) (Statement, Error)

	// Close this connection to the database
	Close() Error
}
