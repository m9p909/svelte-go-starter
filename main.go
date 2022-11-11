package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
  "github.com/gofiber/fiber/v2/middleware/logger"
	"database/sql"
	"log"
	"fmt"
	 _ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "Yipyapyop1"
	dbname = "science"
)

func getDBConnection()  (*sql.DB){
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	connection, err := sql.Open("postgres", psqlconn)
  if err != nil {
		log.Fatalf("could not connect to db %s", err)
	}


	return connection;
}



func main() {

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000, http://localhost:5173",
	}))

	db := getDBConnection();

	app.Static("/", "./client/dist")
	app.Use(logger.New())

	app.Get("/api/helloWorld", func(c *fiber.Ctx) error {
		row := db.QueryRow("select key from hello_world")
		var key string
		err := row.Scan(&key)
		if err != nil {
			log.Fatalf("db error %s", err)
		}
		return c.SendString(key)
	})

	app.Listen(":3000")
}
