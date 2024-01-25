package dbs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slothmakeout/movies-project/data/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dbInstance *gorm.DB

func GetDB() *gorm.DB {
	return dbInstance
}

func setupPostgres() (*gorm.DB, error) {

	// newLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	// 	logger.Config{
	// 		SlowThreshold:             time.Second, // Slow SQL threshold
	// 		LogLevel:                  logger.Info, // Log level
	// 		IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
	// 		Colorful:                  true,        // Disable color
	// 	},
	// )
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	fmt.Println(dbHost)
	//postgres://postgres:1234@localhost:5432/estate
	connectionString := fmt.Sprintf("%s://%s:%s@localhost:%s/%s",
		dbHost,
		dbUser,
		dbPassword,
		dbPort,
		dbName)
	fmt.Println(connectionString)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InitializeDatabaseLayer() error {

	var db *gorm.DB
	var err error

	db, err = setupPostgres()

	if err != nil {
		return err
	}

	err = models.AutoMigrate(db)
	if err != nil {
		return err
	}
	dbInstance = db
	fmt.Println("migration completed")
	return nil
}
