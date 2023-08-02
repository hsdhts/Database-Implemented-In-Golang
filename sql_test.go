package Database_go

import (
	"context"
	"fmt"
	"testing"
)

func TestExecSqlCreate(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO  customer(id, name) VALUES ('wonderwomen', 'Wonderwomen')"
	_, err := db.ExecContext(ctx, script)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert customer")
}

func TestExecSqlRead(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name  FROM customer WHERE id= 'batman'"

	var id, name string
	err := db.QueryRowContext(ctx, query).Scan(&id, &name)

	if err != nil {
		panic(err)
	}

	fmt.Println("Id: ", id)
	fmt.Println("Name: ", name)
}
