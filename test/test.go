package test

import (
	"fmt"
	"sql"
	"os"
	_ "sql/sqlite3"
)

func fib(i int) int {
	if i <= 1 {
		return i
	}
	return fib(i-1)+fib(i-2)
}

func test_sql(dsn string) os.Error{
	conn, err := sql.Connect(dsn)
	if err != nil {
		fmt.Print(err)
		return nil
	}
	
	
	err = conn.Execute("CREATE TABLE fib (pos INTEGER, val INTEGER);")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	
	stmt, _ := conn.Prepare("INSERT INTO fib VALUES(?, ?);")
	
	for i := 0; i < 100; i++ {
		stmt.Execute(i, fib(i))
	}
	stmt.Close()
	

	conn.Close()
	return nil
}

func main() {
	// test the sqlite3 driver
	fmt.Print("Test sql/sqlite3: \n")
	_ = test_sql("sqlite3://test.db")
}
