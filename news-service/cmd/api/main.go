package main

import (
	"database/sql"
	"embed"
	"fmt"
	"gopkg.in/reform.v1/dialects/postgresql"
	"log"
	"news-service/data"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"gopkg.in/reform.v1"
)

const webPort = "85"

var counts int64

//go:embed migrations/*.sql
var EmbedMigrations embed.FS

type Config struct {
	Repo   data.Repository
	Client *fiber.App
}

func main() {
	log.Println("Starting financial service")

	err := godotenv.Load("example.env")
	if err != nil {
		log.Panic("Error loading .env file", err)
	}

	conn := connectToDB()
	if conn == nil {
		log.Panic("Can't connect to Postgres!")
	}

	db := reform.NewDB(conn, postgresql.Dialect, nil)

	goose.SetBaseFS(EmbedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		log.Panic(err)
	}

	migrationsDir := os.Getenv("GOOSE_MIGRATION_DIR")
	if err := goose.Up(conn, migrationsDir); err != nil {
		log.Panic(err)
	}

	app := fiber.New()

	config := Config{
		Client: app,
	}
	config.setupRepo(db)

	config.routes(app)

	log.Printf("Starting server on port %s\n", webPort)
	if err := app.Listen(fmt.Sprintf(":%s", webPort)); err != nil {
		log.Panic(err)
	}
}

// openDB opens connection with the database
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// connectToDB connects to the database
func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")
	log.Println(dsn)

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Postgres not yet ready ...")
			log.Println(err)
			counts++
		} else {
			log.Println("Connected to Postgres!")
			return connection
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for two seconds....")
		time.Sleep(2 * time.Second)
		continue
	}
}

// setupRepo sets up the Repo
func (app *Config) setupRepo(db *reform.DB) {
	repo := data.NewPostgresRepository(db)
	app.Repo = repo
}
