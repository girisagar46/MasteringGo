package main

import (
	"database/sql"
	"fmt"
	"github.com/marianogappa/sqlparser"

	//_ "github.com/girisagar46/699b0eb1f1ae1269101ad9e982ee5b0ef508b7a9"
	//mydb "github.com/girisagar46/699b0eb1f1ae1269101ad9e982ee5b0ef508b7a9"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	master := master("root:@tcp(127.0.0.1:3306)/repl_example")
	replica := replica("root:@tcp(127.0.0.1:13306)/repl_example")
	fmt.Println(master.Ping(), replica.Ping())
	defer closeConn(master)
	defer closeConn(replica)
	simpleQuery(replica)
	simplePreparedStatementQuery(replica, 1)
	singleRowQuery(replica, 2)

	modifyData(master)
	deleteData(master, 6)
	transactions(master)
	preparedStatement(master)
	handleErrors(master)
	workingWithNullValues(master)
	workingWithUnknownColumns(master)
	sqlParser(master)

}

func sqlParser(db *sql.DB) {
	var sqls = []string{
		"SELECT * FROM users",
		"UPDATE test SET foo = 'foo', bar = 'bar' WHERE id = '10'",
		"INSERT INTO test (id, foo, bar) VALUES ('10', 'foo', 'bar')",
		"DELETE FROM test WHERE id = '10'",
	}
	for _, query := range sqls {
		query, err := sqlparser.Parse(query)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+#v\n", query)
	}

}

func workingWithUnknownColumns(db *sql.DB) {
	//err := db.QueryRow("SHOW PROCESSLIST")
	rows, err := db.Query("SHOW PROCESSLIST;")
	if err != nil {
		log.Fatal(err)
	}
	cols, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}
	vals := make([]interface{}, len(cols))

	for i := range cols {
		vals[i] = new(sql.RawBytes)
	}
	for rows.Next() {
		err = rows.Scan(vals...)
		if err != nil {
			log.Fatal(err)
		}
		// Now you can check each element of vals for nil-ness,
		// and you can use type introspection and type assertions
		// to fetch the column into a typed variable.
	}
}

func workingWithNullValues(db *sql.DB) {
	var s sql.NullString
	err := db.QueryRow("SELECT foo FROM test WHERE id = ?", 100).Scan(&s)
	if err != nil {
		log.Println(err)
	}
	if s.Valid {
		fmt.Println(s.String)
	} else {
		fmt.Println("NULL")
	}

	// BETTER APPROACH
	fmt.Println("Better approach")
	var foo, bar sql.NullString
	rows, _ := db.Query(`SELECT foo, COALESCE (bar, '') as bar FROM test WHERE id = ?`, 100)

	for rows.Next() {
		err := rows.Scan(&foo, &bar)
		fmt.Println(foo, bar)

		if err != nil {
			log.Println(err)
		}
	}
}

func handleErrors(db *sql.DB) {
	/*
		// Errors From Iterating Resultsets
		for rows.Next() {
			// ...
			// An abnormal termination automatically calls rows.Close()
		}
		if err = rows.Err(); err != nil { // could be the result of a variety of errors in the rows.Next() loop
			// handle the error here
		}
	*/

	/*
		// Errors From Closing Resultsets
		for rows.Next() {
			// ...
			break; // whoops, rows is not closed! memory leak...
		}
		// do the usual "if err = rows.Err()" [omitted here]...
		// it's always safe to [re?]close here:
		if err = rows.Close(); err != nil {
			// but what should we do if there's an error? Ans: Either log it. Or panic. Or ignore it.
			log.Println(err)
		}
	*/

	/*
		// Errors From QueryRow()
		var name string
		err = db.QueryRow("select name from users where id = ?", 1).Scan(&name)

		// What if there was no user with id = 1?
		// Go defines a special error constant, called sql.ErrNoRows, which is returned from QueryRow() when the result is empty.
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(name)

		// BETTER APPROACH:
		var name string
		err = db.QueryRow("select name from users where id = ?", 1).Scan(&name)
		if err != nil {
			if err == sql.ErrNoRows {
				// there were no rows, but otherwise no error occurred
			} else {
				log.Fatal(err)
			}
		}
		fmt.Println(name)
	*/

	/*
		//Identifying Specific Database Errors

		// Code we tend to write:
		rows, err := db.Query("SELECT someval FROM sometable")
		// err contains:
		// ERROR 1045 (28000): Access denied for user 'foo'@'::1' (using password: NO)
		if strings.Contains(err.Error(), "Access denied") {
			// Handle the permission-denied error
		}

		// BETTER APPROACH:
		// Handle such error from DB Driver's perspective
		if driverErr, ok := err.(*mysql.MySQLError); ok { // Now the error number is accessible directly
			if driverErr.Number == 1045 { // number is DB specific and also the 1045 is a magic number which is a code smell
				// Handle the permission-denied error
			}
		}

		// MORE BETTER APPROACH:
		if driverErr, ok := err.(*mysql.MySQLError); ok {
			if driverErr.Number == mysqlerr.ER_ACCESS_DENIED_ERROR { // mysqlerr is another package -> https://github.com/VividCortex/mysqlerr
				// Handle the permission-denied error
			}
		}
	*/

	/*
		// Handling Connection Errors
		//What if your connection to the database is dropped, killed, or has an error?
		//You don’t need to implement any logic to retry failed statements when this happens.
		//As part of the connection pooling in database/sql, handling failed connections is built-in.
		//If you execute a query or other statement and the underlying connection has a failure,
		//Go will reopen a new connection (or just get another from the connection pool) and retry, up to 10 times.
	*/
}

func preparedStatement(db *sql.DB) {
	//At the database level, a prepared statement is bound to a single database connection.
	// The typical flow is that the
	// - client sends a SQL statement with placeholders to the server for preparation,
	// - the server responds with a statement ID, and
	// - then the client executes the statement by sending its ID and parameters.
}

func transactions(db *sql.DB) {
	// In Go, a transaction is essentially an object that reserves a connection to the datastore.
	// lets you do all the operations with a guarantee that they’ll be executed on the same connection.
	tx, err := db.Begin() // only the Tx object is in transaction, hence from hereon, only deal with the Tx object
	if err != nil {
		log.Fatal(err)
	}
	_, err = tx.Exec("INSERT INTO test (foo, bar) VALUES (?, ?)", "John", "Cena")
	if err != nil {
		log.Fatal(err)
	}
	_, err = tx.Exec("INSERT INTO test (foo, bar) VALUES (?, ?)", "Peace", "maker")
	if err != nil {
		log.Fatal(err)
	}
	err = tx.Commit() // or tx.Rollback()
	if err != nil {
		log.Fatal(err)
	}
}

func deleteData(db *sql.DB, id int) {
	_, err := db.Exec("DELETE FROM test where id = ?", id) // OK
	//_, err := db.Query("DELETE FROM test where id = ?", id) // DO NOT EVER DO THIS because Query() will return a sql.Rows which reserves a database connection until the sql.Rows is closed
	if err != nil {
		log.Fatal(err)
	}
}

func modifyData(db *sql.DB) {
	query := "INSERT INTO repl_example.test (foo, bar) VALUES ('foo1115', 'bar1116');"
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
}

func singleRowQuery(db *sql.DB, rowId int) {
	var (
		id       int
		foo, bar string
	)
	err := db.QueryRow("SELECT * FROM test WHERE id = ?", rowId).Scan(&id, &foo, &bar) // returns at most one row
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id, foo, bar)

	// call QueryRow() on a prepared statement
	stmt, err := db.Prepare("SELECT foo FROM test WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(stmt)
	var name string
	err = stmt.QueryRow(rowId).Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)
}

func simplePreparedStatementQuery(db *sql.DB, id int) {
	stmt, err := db.Prepare("SELECT * FROM test WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(stmt)
	rows, err := stmt.Query(id) // Under the hood, db.Query() actually prepares, executes, and closes a prepared statement.
	if err != nil {
		log.Fatal(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)
	for rows.Next() {
		var id int
		var foo, bar string
		err := rows.Scan(&id, &foo, &bar)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, foo, bar)
	}
}

func simpleQuery(db *sql.DB) {
	var (
		id       int
		foo, bar string
	)
	rows, err := db.Query("SELECT id, foo, bar FROM test;")
	if err != nil {
		log.Fatal(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)
	for rows.Next() {
		err := rows.Scan(&id, &foo, &bar) // read the columns in each row into variables
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, foo, bar)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func master(connStr string) *sql.DB {
	master, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return master
}

func replica(connStr string) *sql.DB {
	replica, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return replica
}

func closeConn(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
}
