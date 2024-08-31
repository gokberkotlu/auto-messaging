package database

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Database struct {
	connection *gorm.DB
}

type dbConfig struct {
	host     string
	user     string
	password string
	dbname   string
	port     uint16
}

var (
	DatabaseInstance *Database
	lock             = &sync.Mutex{}
)

func GetDB() (*gorm.DB, error) {
	if DatabaseInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if DatabaseInstance == nil {
			database, err := newDBConn()
			if err != nil {
				return nil, err
			}
			DatabaseInstance = database
		}
	}

	return DatabaseInstance.connection, nil
}

func newDBConn() (*Database, error) {
	intPortValue, _ := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	dbConfig := dbConfig{
		host:     "localhost",
		user:     os.Getenv("POSTGRES_USER"),
		password: os.Getenv("POSTGRES_PASSWORD"),
		dbname:   os.Getenv("POSTGRES_DB"),
		port:     uint16(intPortValue),
	}

	db, err := gorm.Open(postgres.Open(dbConfig.getDSN()), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return &Database{
		connection: db,
	}, nil
}

func (db *Database) Close() error {
	sqlDB, err := db.connection.DB()
	if err != nil {
		return fmt.Errorf("failed to get sqlDB: %w", err)
	}
	return sqlDB.Close()
}

func (dbConfig *dbConfig) getDSN() string {
	fmt.Printf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable\n", dbConfig.host, dbConfig.user, dbConfig.password, dbConfig.dbname, dbConfig.port)
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", dbConfig.host, dbConfig.user, dbConfig.password, dbConfig.dbname, dbConfig.port)
}
