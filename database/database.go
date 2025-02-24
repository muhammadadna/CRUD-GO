// package database

// import (
// 	"database/sql"

// 	_ "github.com/go-sql-driver/mysql"
// )

// func InitDatabase() *sql.DB {
// 	dsn := "root@tcp(localhost:3306)/CRUD-GO"
// 	db, err := sql.Open("mysql", dsn)
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = db.Ping()
// 	if err != nil {
// 		panic(err)
// 	}

// 	return db
// }

package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func InitDatabase() *sql.DB {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	return db
}
