package Database_go

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// create
func TestExecSqlCreate(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO  customer(id, name) VALUES ('captainAmerica', 'CaptainAmerica')"
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

	query := "SELECT id, name  FROM customer WHERE id= 'superman'"

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

func TestQuerySwl(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id: ", id)
		fmt.Println("Name: ", name)
	}

}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name, email string
		var balance int32
		var rating float64
		var birth_date, created_at time.Time
		var married bool
		err = rows.Scan(&id, &name, &email, &balance, &rating, &birth_date, &married, &created_at)
		if err != nil {
			panic(err)
		}
		fmt.Println("===============")
		fmt.Println("Id: ", id)
		fmt.Println("Name: ", name)
		fmt.Println("Email: ", email)
		fmt.Println("balance: ", balance)
		fmt.Println("rating:", rating)
		fmt.Println("birthDate: ", birth_date)
		fmt.Println("married: ", married)
		fmt.Println("created_at: ", created_at)
	}
}
