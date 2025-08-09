package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/vprates-22/pratesBet/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func buildDsn() (string, error) {
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

func migrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Group{},
		&models.Permission{},
		&models.GroupPermission{},
		&models.Person{},
		&models.Address{},
		&models.State{},
		&models.Country{},
		&models.Configuration{},
		&models.Language{},
	)
}

func New() error {
	var dsn string
	var err error

	dsn, err = buildDsn()
	if err != nil {
		return err
	}

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	err = migrateDB(DB)
	if err != nil {
		return err
	}

	return nil
}
