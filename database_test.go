package Database_go

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/golang_db")

	if err != nil {
		panic(err)
	}

	// gunakan database

	defer db.Close()

}
