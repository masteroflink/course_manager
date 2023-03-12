package helpers

import (
	"log"
	"main/api/models"

	"gorm.io/gorm"
)

func RefreshUserTable(db *gorm.DB) error {
	err := db.Migrator().DropTable(&models.User{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}
	log.Println("Sucessfully refreshed User table.")
	return nil
}

func RefreshStudentTable(db *gorm.DB) error {
	err := db.Migrator().DropTable(&models.Student{}, "course_students")
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&models.Student{})
	if err != nil {
		return err
	}
	log.Println("Sucessfully refreshed Student table.")
	return nil
}

func RefreshProcessorTable(db *gorm.DB) error {
	err := db.Migrator().DropTable(&models.Professor{}, "course_professors")
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&models.Professor{})
	if err != nil {
		return err
	}
	log.Println("Sucessfully refreshed Professor table.")
	return nil
}

func RefreshCourseTable(db *gorm.DB) error {
	err := db.Migrator().DropTable(&models.Course{}, "course_professors", "course_students")
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&models.Course{})
	if err != nil {
		return err
	}
	log.Println("Sucessfully refreshed Course table.")
	return nil
}
