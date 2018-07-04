package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"database/sql"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func main() {
	database()

	engine := gin.Default()

	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": message()})
	})

	engine.GET("/api/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, AllBooks(db))
	})

	engine.GET("/api/books/:isbn", func(c *gin.Context) {
		isbn := c.Params.ByName("isbn")
		book, found := GetBook(db, isbn)
		if found {
			c.JSON(http.StatusOK, book)
		} else {
			c.AbortWithStatus(http.StatusNotFound)
		}
	})

	engine.Run(port())
}

func database() {
	url := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv("DB_ENGINE"), os.Getenv("PROJECT_DB_USERNAME"), os.Getenv("PROJECT_DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("PROJECT_DB_NAME"))
	log.Printf("Connecting to DB %s\n", url)

	db, err := sql.Open("postgres", url)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	log.Println("Connected.")
}

func message() string {
	msg := os.Getenv("MESSAGE")
	if len(msg) == 0 {
		msg = "Hello Cloud Native Night Muc with SquareScale!"
	}
	return msg
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "9090"
	}
	return ":" + port
}
