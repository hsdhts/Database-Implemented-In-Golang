package Database_go

import (
	"context"
	"database/sql"
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

// query
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
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birth_date sql.NullTime
		var created_at time.Time
		var married bool
		err = rows.Scan(&id, &name, &email, &balance, &rating, &birth_date, &married, &created_at)
		if err != nil {
			panic(err)
		}
		fmt.Println("===============")
		fmt.Println("Id: ", id)
		fmt.Println("Name: ", name)
		if email.Valid {
			fmt.Println("Email: ", email.String)
		}
		fmt.Println("balance: ", balance)
		fmt.Println("rating:", rating)
		if birth_date.Valid {
			fmt.Println("BirthDate: ", birth_date.Time)
		}
		fmt.Println("married: ", married)
		fmt.Println("created_at: ", created_at)
	}
}

// sql injection
func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin"
	password := "admin123"

	script := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, script, username, password)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Login berhasil", username)
	} else {
		fmt.Println("Gagal login")
	}
}

// sql parameter
func TestExecSqlParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "batman"
	password := "batman123"

	script := "INSERT INTO  user(username, password) VALUES (?, ?)"
	_, err := db.ExecContext(ctx, script, username, password)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new user")
}

// auto increment
func TestAutoincrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "batman@gmail.com"
	comment := "test email batman"

	script := "INSERT INTO  comments(email, comment) VALUES (?, ?)"
	result, err := db.ExecContext(ctx, script, email, comment)

	if err != nil {
		panic(err)
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert new comment with id", insertId)
}
