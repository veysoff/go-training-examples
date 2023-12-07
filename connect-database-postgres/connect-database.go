package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	HOST     = "localhost"
	PORT     = 5432
	USER     = "postgres"
	PASSWORD = "123456789"
	DBNAME   = "TestGoDB"
)

func main() {
	// Connection
	connString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, USER, PASSWORD, DBNAME,
	)
	db, err := sql.Open("postgres", connString)
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	// Insert
	var (
		id    int
		title string
	)

	sqlInsertStatement := `
	INSERT INTO tasks (title)
	VALUES ($1)
	RETURNING id`
	id = 0
	err = db.QueryRow(sqlInsertStatement, "Test insert").Scan(&id)
	check(err)
	fmt.Println("New record ID is:", id)

	// Select all
	rows, err := db.Query("select * from tasks")
	check(err)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &title)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, title)
	}
	err = rows.Err()
	check(err)

	// 	// Update
	sqlUpdateStatement := `
	UPDATE tasks
	SET title = $2
	WHERE id = $1;`
	res, err := db.Exec(sqlUpdateStatement, id, "NewTitle")
	check(err)

	count, err := res.RowsAffected()
	check(err)
	fmt.Println(count)

	// Get by id
	type Task struct {
		Id    int
		Title string
	}

	sqlStatement := buildGetByIdSql(1)
	var task Task
	row := db.QueryRow(sqlStatement)
	err = row.Scan(&task.Id, &task.Title)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return
	case nil:
		fmt.Println(task)
	default:
		panic(err)
	}

	// Delete
	sqlDeleteStatement := `
	DELETE FROM tasks
	WHERE id = $1;`
	_, err = db.Exec(sqlDeleteStatement, id)
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func buildGetByIdSql(id int) string {
	return fmt.Sprintf("SELECT * FROM tasks WHERE id='%d';", id)
}
