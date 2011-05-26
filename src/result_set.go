package sql

// API to be implemented for query resultsets.
type ResultSet interface {
    // Get the next row of results.
    Next() bool

    // Scan the current row of results by column index.
    Scan(refs ...interface{}) Error

    // Scan the current row of results by column name.
    NamedScan(refs ...interface{}) Error

    // Release the resources held by a ResultSet.
    Close() Error
}
