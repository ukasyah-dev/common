package testkit

import (
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/ukasyah-dev/common/db/pool"
	"gorm.io/gorm"
)

var db *gorm.DB
var dbName string

func CreateTestDB() {
	// Extract info from DATABASE_URL
	info, err := url.Parse(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	dbName = fmt.Sprintf(
		"%s-test-%d",
		strings.Replace(info.Path, "/", "", 1),
		time.Now().Unix(),
	)

	// Connect to postgres db
	info.Path = "/postgres"
	db, err = pool.Open(info.String())
	if err != nil {
		panic(err)
	}

	// Create new test db
	if err := db.Exec(fmt.Sprintf("CREATE DATABASE \"%s\"", dbName)).Error; err != nil {
		panic(err)
	}

	// Update DATABASE_URL with test db name
	info.Path = "/" + dbName
	os.Setenv("DATABASE_URL", info.String())
}

func DestroyTestDB() {
	// Drop test db
	if err := db.Exec(fmt.Sprintf("DROP DATABASE \"%s\"", dbName)).Error; err != nil {
		panic(err)
	}

	// Close postgres db connection
	sql, _ := db.DB()
	sql.Close()
}
