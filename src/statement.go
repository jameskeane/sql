package sql

type Statement interface {
	// Execute the statement, return a result set and an error.
	Query(params ...interface{}) (ResultSet, Error)

    // Execute the statement, discard the result and return an error.
    Execute(params ...interface{}) Error

    // Free up any resources used by this statement.
    Close() Error
}
