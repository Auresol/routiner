package db

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	//"gorm.io/gorm/taskger"

	"github.com/joho/godotenv"
)

func NewDBConnection(env string) *gorm.DB {

	err := godotenv.Load(".env." + env)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbName := os.Getenv("DB_NAME")

	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{
		//Taskger: taskger.Default.TaskMode(taskger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
