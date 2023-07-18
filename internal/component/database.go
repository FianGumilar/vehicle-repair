package component

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/FianGumilar/vehicle-repair/internal/config"
	_ "github.com/go-sql-driver/mysql"
)

func GetDatabaseConnection(conf config.Config) *sql.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		conf.DB.User,
		conf.DB.Pass,
		conf.DB.Host,
		conf.DB.Port,
		conf.DB.Name)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed Connect to Database: %s", err.Error())
	}

	return db
}
