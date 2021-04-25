package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type Test struct {
	id      int32
	content string
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "user:password@(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		errors.Wrap(err, "Database Connect Fail!")
	}
	return db, nil
}

func main() {

	db, err := connect()
	defer db.Close()
	if err != nil {
		fmt.Printf("------>Error: %T %v\n\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("------>Trace: %+v\n", err)
		return
	}
	test, err := query(db, 10)
	if err != nil {
		fmt.Printf("------->Error: %T %v\n\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("------->Trace: %+v\n", err)
		return
	}
	fmt.Println(test)
}

func query(db *sql.DB, test_id int32) (Test, error) {
	var test Test
	sqlStatement := "select * from test1 where tid=?"
	rows, err := db.Query(sqlStatement, test_id)
	if err == sql.ErrNoRows {
		wrapContent := "No Rows. SqlStatement: " + sqlStatement
		errors.Wrap(err, wrapContent)
	}
	if err != nil {
		wrapContent := "sqlStatement: " + sqlStatement
		errors.Wrap(err, wrapContent)
	}
	for rows.Next() {
		_ = rows.Scan(&test)
	}
	return test, nil
}
