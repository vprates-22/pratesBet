package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Dsn() (string, error) {
	var dsn string
	var err error

	err = godotenv.Load()
	if err != nil {
		return dsn, err
	}

	m := map[string]string{
		"host":     "POSTGRES_HOST",
		"user":     "POSTGRES_USER_NAME",
		"password": "POSTGRES_PASSWORD",
		"dbname":   "POSTGRES_DATABASE",
		"port":     "POSTGRES_PORT",
	}

	for k, kEnv := range m {
		v := os.Getenv(kEnv)

		if v == "" {
			err = fmt.Errorf("ERROR: Missing %s value in the .env file", kEnv)
			return "", err
		}

		dsn += k + "=" + v + " "
	}

	return dsn, err
}

func New() error {
	var dsn string
	var err error

	dsn, err = Dsn()
	if err != nil {
		return err
	}

	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		return err
	}

	return nil
}
