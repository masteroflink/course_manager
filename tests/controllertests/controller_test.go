package controllertests

import (
	"fmt"
	"log"
	"main/api/controllers"
	"main/api/models"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var server = controllers.Server{}
var userInstance = models.User{}
var studentInstance = models.Student{}
var professorInstance = models.Professor{}
var courseInstance = models.Course{}

func TestMain(m *testing.M) {
	var err error
	err = godotenv.Load(os.ExpandEnv("../../.env"))
	if err != nil {
		log.Fatalf("Error gettting env %v\n", err)
	}
	Database()

	os.Exit(m.Run())
}

func Database() {
	var err error

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("TEST_DB_HOST"), os.Getenv("TEST_DB_PORT"), os.Getenv("TEST_DB_USER"), os.Getenv("TEST_DB_NAME"), os.Getenv("TEST_DB_PASSWORD"))

	server.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Cannot connect to PostgreSQL database.")
		log.Fatal("Error: ", err)
	} else {
		fmt.Printf("Connected to PostgreSQL database.\n")
	}
}
