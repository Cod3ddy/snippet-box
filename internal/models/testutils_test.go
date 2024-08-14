package models

import (
	"database/sql"
	"os"
	"testing"
)

func newTestDB(t *testing.T) *sql.DB {
	 // Establish a sql.DB connection pool for our test database. Because our
    // setup and teardown scripts contains multiple SQL statements, we need
    // to use the "multiStatements=true" parameter in our DSN. This instructs
    // our MySQL database driver to support executing multiple SQL statements
    // in one db.Exec() call.

	db, err := sql.Open("mysql", "test_web:pass@/test_snippetbox?parseTime=true&multiStatements=true")
    if err != nil {
        t.Fatal(err)
    }

	script, err := os.ReadFile("./testdata/setup.sql")
	if err != nil{
		db.Close()
		t.Fatal(err)
	}

	_, err = db.Exec(string(script))
	if err != nil{
		db.Close()
		t.Fatal(err)
	}

	 // Use t.Cleanup() to register a function *which will automatically be
    // called by Go when the current test (or sub-test) which calls newTestDB() 
    // has finished*. In this function we read and execute the teardown script, 
    // and close the database connection pool.

	t.Cleanup(func() {
		defer db.Close()

		script, err := os.ReadFile("./testdata/teardown.sql")
		if err != nil{
			t.Fatal(err)
		}

		_, err = db.Exec(string(script))
		if err != nil{
			t.Fatal(err)
		}
	})

	// Return the database connection pool
	return db
}