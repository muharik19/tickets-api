package databases

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	cm "github.com/azura-labs/common"
)

func ConnectDb() (*sql.DB, error) {
	conString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cm.Config.DBUser, cm.Config.DBPassword, cm.Config.DBHost, cm.Config.DBPort, cm.Config.DBName)

	db, err := sql.Open("mysql", conString)
	if err != nil {
		fmt.Println("Failed to connect database!", err)
	}

	testPing := db.Ping()
	if testPing != nil {
		log.Fatal("Error pinging database: " + testPing.Error())
	}

	return db, err
}

func Query(sql string) (*sql.Rows, error) {
	ctx := context.Background()
	db, _ := ConnectDb()
	defer db.Close()
	return db.QueryContext(ctx, sql)
}

func QueryRow(sql string) *sql.Row {
	ctx := context.Background()
	db, _ := ConnectDb()
	defer db.Close()
	return db.QueryRowContext(ctx, sql)
}

func Exec(sql string) bool {
	ctx := context.Background()
	db, _ := ConnectDb()
	defer db.Close()
	_, err := db.ExecContext(ctx, sql)
	if err != nil {
		return false
	}
	return true
}

func QueryCount(queryCount string) int {
	var totalRows int
	rowsCount, _ := Query(queryCount)
	for rowsCount.Next() {
		errScanCount := rowsCount.Scan(&totalRows)
		if errScanCount != nil {
			return 0
		}
	}

	return totalRows
}
