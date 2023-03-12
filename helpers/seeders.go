package helpers

import (
	"log"
	"main/api/models"

	"gorm.io/gorm"
)

func SeedOneUser(db *gorm.DB) (models.User, error) {
	user := models.User{
		Name: models.Name{
			First: "Mickey",
			Last:  "Mouse",
		},
		Email:    "mickeytest@example.com",
		Password: "password",
	}

	err := db.Model(&models.User{}).Create(&user).Error
	if err != nil {
		log.Fatalf("Can't seed one user: %v", err)
	}
	return user, nil
}

func SeedUsers(db *gorm.DB) error {
	users := []models.User{
		{
			Name: models.Name{
				First: "Kevin",
				Last:  "Simpson",
			},
			Email:    "kevintest@example.com",
			Password: "password",
		},
		{
			Name: models.Name{
				First: "James",
				Last:  "Doucet",
			},
			Email:    "jamestest@example.com",
			Password: "password",
		},
		{
			Name: models.Name{
				First: "Mario",
				Last:  "Bros",
			},
			Email:    "mario@example.com",
			Password: "password",
		},
	}

	err := db.Model(&models.User{}).CreateInBatches(users, 100).Error
	if err != nil {
		return err
	}
	return nil
}
