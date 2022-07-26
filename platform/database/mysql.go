package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"strconv"
	"time"

	"gitlab.com/d6825/golang_template/pkg/utils"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql" // load driver for Mysql
)

// MysqlConnection func for connection to Mysql database.
func MysqlConnection() (*sqlx.DB, error) {
	// Define database connection settings.
	maxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	maxLifetimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))

	// Build Mysql connection URL.
	mysqlConnURL, err := utils.ConnectionURLBuilder("mysql")
	if err != nil {
		return nil, err
	}

	// Define database connection for Mysql.
	db, err := sqlx.Connect("mysql", mysqlConnURL)
	if err != nil {
		return nil, fmt.Errorf("error, not connected to database, %w", err)
	}

	// Set database connection settings:
	// 	- SetMaxOpenConns: the default is 0 (unlimited)
	// 	- SetMaxIdleConns: defaultMaxIdleConns = 2
	// 	- SetConnMaxLifetime: 0, connections are reused forever
	db.SetMaxOpenConns(maxConn)
	db.SetMaxIdleConns(maxIdleConn)
	db.SetConnMaxLifetime(time.Duration(maxLifetimeConn))

	// Try to ping database.
	if err := db.Ping(); err != nil {
		defer db.Close() // close database connection
		return nil, fmt.Errorf("error, not sent ping to database, %w", err)
	}

	return db, nil
}

func GormMysqlConnection() *gorm.DB {
	dsn := os.Getenv("GORM_MYSQL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
