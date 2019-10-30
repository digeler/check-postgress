package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func main() {

	host := os.Getenv("host")
	port := 5432
	user := os.Getenv("user")
	password := os.Getenv("password")
	dbname := os.Getenv("dbname")

	for {
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s",
			host, port, user, password, dbname)

		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			fmt.Println(err)

		}
		defer db.Close()

		err = db.Ping()
		if err != nil {
			fmt.Println(err)
			//log.Fatal("failed db connection")
		}

		fmt.Println("real error ", err, time.Now().UTC())

		time.Sleep(2 * time.Second)
	}
}
