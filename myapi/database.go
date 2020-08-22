package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var err error
var db *sql.DB

func init() {
	url := os.Getenv("DATABASE_URL")
	db, err = sql.Open("postgres", url)

	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("Connected")
}

func createTable() {
	// url := os.Getenv("DATABASE_URL")
	// db, err := sql.Open("postgres", url)

	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
	// defer db.Close()

	// log.Println("Connected")

	createTb := `
	CREATE TABLE IF NOT EXISTS todos(
		id SERIAL PRIMARY KEY,
		title TEXT,
		status TEXT
	);
	`

	_, err = db.Exec(createTb)

	if err != nil {
		log.Fatal("can't create table", err)
	}

	log.Println("Create table success")
}

func insertTodo(title, status string) {
	// db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	// if err != nil {
	// 	log.Fatal("Connect to database error", err)
	// }
	// defer db.Close()
	row := db.QueryRow("INSERT INTO todos (title, status) values ($1, $2)  RETURNING id", title, status)
	var id int
	err = row.Scan(&id)
	if err != nil {
		fmt.Println("can't scan id", err)
		return
	}
	fmt.Println("insert todo success id : ", id)
}

func queryTodo(ID int) {
	// db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	// if err != nil {
	// 	log.Fatal("Connect to database error", err)
	// }

	stmt, err := db.Prepare("SELECT id, title,status FROM todos where id=$1")
	if err != nil {
		log.Fatal("can't prepare statement", err)
	}

	row := stmt.QueryRow(ID)
	var rowID int
	var title, status string

	err = row.Scan(&rowID, &title, &status)
	if err != nil {
		log.Fatal("can't scan row to variables", err)
	}
	fmt.Println("one row", rowID, title, status)
}

func updateTodo(ID int, status string) {
	// db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	// if err != nil {
	// 	log.Fatal("Connect to database error", err)
	// }

	stmt, err := db.Prepare("UPDATE todos SET status=$2 where id=$1")
	if err != nil {
		log.Fatal("can't prepare statement", err)
	}

	if _, err := stmt.Exec(ID, status); err != nil {
		log.Fatal("can't update record", err)
	}

	fmt.Println("update row id :", ID, "success")
}

func queryAllTodo() {
	stmt, err := db.Prepare("SELECT id, title,status FROM todos")
	if err != nil {
		log.Fatal("can't prepare statement", err)
	}

	row, err := stmt.Query()
	if err != nil {
		log.Fatal("can't execute query", err)

	}
	var rowID int
	var title, status string

	defer row.Close()
	for row.Next() {
		err = row.Scan(&rowID, &title, &status)
		if err != nil {
			log.Fatal("can't scan row to variables", err)
		}
		fmt.Println("Todo ID :", rowID, title, status)
		// fmt.Printf("%#v", row)
	}
}

func main() {
	// insertTodo("sleep", "active")
	// updateTodo(3, "inactive")
	// queryTodo(3)
	queryAllTodo()

	db.Close()
}
