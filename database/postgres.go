package database

import (
	"fmt"
	"log"
	"os"

	// "gorm.io/gorm/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DBConn *gorm.DB

func ConnectDB() {

	// Access DB credentials from environment
	host := os.Getenv("db_host")
	user := os.Getenv("db_user")
	password := os.Getenv("db_password")
	dbname := os.Getenv("db_name")
	dbport := os.Getenv("db_port")

	fmt.Println("Starting connection with Postgres Db")
	dsn := user + "://postgres:" + password + "@" + host + ":" + dbport + "/" + dbname + "?sslmode=disable"

	//db, err := gorm.Open(postgres.Open(dsn) , &gorm.Config{})
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		}})

	if err != nil {
		panic("Database connection failed.")
	}

	log.Println("Connection successful.")

	DBConn = db

	// Migrate your models
	// err = db.AutoMigrate(
	// 	&model.Login{},
	// 	&model.Item{},
	// )
	// if err != nil {
	// 	log.Fatal("failed to migrate models: ", err)
	// }

	fmt.Println("Data Migration complete.")
}

func TableExists(db *gorm.DB, tableName string) bool {
	var exists bool
	query := fmt.Sprintf("SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_name = '%s')", tableName)
	err := db.Raw(query).Scan(&exists).Error
	if err != nil {
		log.Printf("Error checking if table %s exists: %v", tableName, err)
		return false
	}
	return exists
}
