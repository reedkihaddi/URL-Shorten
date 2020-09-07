package models

import (
	"database/sql"
	"fmt"
	"log"
	//PostgreSQL driver
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "saurabh"
	password = "yourpassword"
	dbname   = "mydb"
)

//InitDB initializes the  psql database.
func InitDB() (*sql.DB, error) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	
	if err != nil {
		log.Println("Error opening DB")
		return nil, err
	}
	
	err = db.Ping()
	if err != nil {
		log.Println("Error connecting to DB")
		panic(err)
	}
	
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS web_url(id text primary key not null, url text not null);")

	if err != nil {
		fmt.Println("Preparing DB statement")
		log.Println(err)
		return nil, err
	}
	
	_, err = stmt.Exec()
	
	if err != nil {
		fmt.Println("Executing DB statement")
		log.Println(err)
		return nil, err
	}
	
	return db, nil
}
