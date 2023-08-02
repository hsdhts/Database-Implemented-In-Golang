package Database_go

import (
	"context"
	"fmt"
	"testing"
)

// create
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

// read
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

// update
func TestExecSqlUpdate(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "UPDATE  customer  SET name= 'KucingGarong' WHERE id= 'wonderwomen'"

	result, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("RowsAffected: ", rowsAffected)
}

// delete
func TestExecSqlDelete(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "DELETE  FROM customer WHERE id= 'batman'"

	result, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Success delete batman", rowsAffected)
	//fmt.Println("RowsAffected: ", rowsAffected)
}
