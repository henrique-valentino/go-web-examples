package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	id        string
	userName  string
	password  string
	createdAt time.Time
}

var (
	username = "root"
	password = "root"
	host_db  = "localhost"
	port_db  = "3306"
	name_db  = "go-mysql"
)

func main() {

	// open connection with database mysql
	// to this was necessary get driver
	
	connectionStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
	 username, password, host_db, port_db, name_db)
	fmt.Println(connectionStr)
	db, err := sql.Open("mysql", connectionStr)
	if err != nil {
		log.Fatal(err)
	}

	// verify connection with database
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	{ //create new table
		query := `
			CREATE TABLE users (
				id INT AUTO_INCREMENT, 
				username TEXT NOT NULL, 
				password TEXT NOT NULL, 
				created_at DATETIME,
				PRIMARY KEY (id)
			);`

		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		}
	}

	{ // insert a new user
		userToInsert := User{
			userName:  "johndoe",
			password:  "secret",
			createdAt: time.Now(),
		}

		result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?,?,?)`, userToInsert.userName, userToInsert.password, userToInsert.createdAt)
		if err != nil {
			log.Fatal(err)
		}

		id, _ := result.LastInsertId()
		fmt.Println(id)
	}

	{ // query a single user
		var userResponse User
		query := "SELECT id, username, password, created_at FROM users WHERE id = ?"
		if err := db.QueryRow(query, 1).Scan(&userResponse.id, &userResponse.userName, &userResponse.password, &userResponse.createdAt); err != nil {
			log.Fatal(err)
		}
		fmt.Println(userResponse)
	}

	{ // query all users
		rows, err := db.Query("SELECT id, username, password, created_at FROM users")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var users []User
		for rows.Next() {
			var u User
			if err := rows.Scan(&u.id, &u.userName, &u.password, &u.createdAt); err != nil {
				log.Fatal(err)
			}

			users = append(users, u)
		}

		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%#v", users)
	}

	{ // delete
		_, err := db.Exec("DELETE FROM users WHERE id = ? ", 1)
		if err != nil {
			log.Fatal(err)
		}

	}

}
