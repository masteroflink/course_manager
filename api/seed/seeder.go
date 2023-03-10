package seed

import (
	"log"
	"main/api/models"

	"gorm.io/gorm"
)

var users = []models.User{
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

var courses = []models.Course{
	{
		College:      "Engineering",
		Department:   "Computer Science",
		CourseNumber: "101",
		MaxCapacity:  300,
	},
	{
		College:      "Liberal Arts & Sciences",
		Department:   "English",
		CourseNumber: "103",
		MaxCapacity:  350,
	},
	{
		College:      "Liberal Arts & Sciences",
		Department:   "Biology",
		CourseNumber: "105",
		MaxCapacity:  2,
	},
}

var students = []models.Student{
	{
		Name: models.Name{
			First: "James",
			Last:  "Doucet",
		},
		Address: models.Address{
			Raw: "136 Highland Dr Burkburnett, TX, 76354",
		},
		Email:            "jamestest@example.com",
		Phone:            "(940) 569-3810",
		GPA:              3.15,
		Credits:          20,
		AttemptedCredits: 20,
		DegreeLevel:      "bachelors",
		FieldOfStudy:     "Mathematics",
	},
	{
		Name: models.Name{
			First:  "Mickey",
			Middle: "Disney",
			Last:   "Mouse",
		},
		Address: models.Address{
			StreetNumber: "136",
			StreetName:   "Highland Dr",
			City:         "Burkburnett",
			State:        "Texas",
			Country:      "United States of America",
			CountryCode:  "US",
			PostalCode:   "76354",
		},
		Email:            "mickeytest@example.com",
		Phone:            "(940) 569-3810",
		GPA:              3.55,
		Credits:          32,
		AttemptedCredits: 32,
		DegreeLevel:      "bachelors",
		FieldOfStudy:     "Computer Science",
	},
	{
		Name: models.Name{
			First:  "Donald",
			Middle: "Disney",
			Last:   "Duck",
		},
		Address: models.Address{
			StreetNumber: "136",
			StreetName:   "Highland Dr",
			City:         "Burkburnett",
			State:        "Texas",
			Country:      "United States of America",
			CountryCode:  "US",
			PostalCode:   "76354",
		},
		Email:            "donaldtest@example.com",
		Phone:            "(940) 569-3810",
		GPA:              3.55,
		Credits:          32,
		AttemptedCredits: 32,
		DegreeLevel:      "bachelors",
		FieldOfStudy:     "Computer Science",
	},
}

var professors = []models.Professor{
	{
		Name: models.Name{
			First:  "Mickey",
			Middle: "Disney",
			Last:   "Mouse",
		},
		Address: models.Address{
			StreetNumber: "136",
			StreetName:   "Highland Dr",
			City:         "Burkburnett",
			State:        "Texas",
			Country:      "United States of America",
			CountryCode:  "US",
			PostalCode:   "76354",
		},
		Email:    "drmouse@example.com",
		Position: "Associate Professor",
	},
	{
		Name: models.Name{
			First:  "Donald",
			Middle: "Disney",
			Last:   "Duck",
		},
		Address: models.Address{
			StreetNumber: "136",
			StreetName:   "Highland Dr",
			City:         "Burkburnett",
			State:        "Texas",
			Country:      "United States of America",
			CountryCode:  "US",
			PostalCode:   "76354",
		},
		Email:    "donald@example.com",
		Position: "Associate Professor",
	},
	{
		Name: models.Name{
			First:  "Mario",
			Middle: "Nin",
			Last:   "Bro",
		},
		Address: models.Address{
			StreetNumber: "136",
			StreetName:   "Highland Dr",
			City:         "Burkburnett",
			State:        "Texas",
			Country:      "United States of America",
			CountryCode:  "US",
			PostalCode:   "76354",
		},
		Email:    "drmario@example.com",
		Position: "Associate Professor",
	},
}

func Load(db *gorm.DB) {
	err := db.Debug().Migrator().DropTable(&models.User{}, &models.Student{}, &models.Professor{}, &models.Course{}, "course_professors", "course_students")
	if err != nil {
		log.Fatalf("Error while dropping table: %v", err)
		return
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Student{}, &models.Professor{}, &models.Course{})
	if err != nil {
		log.Fatalf("Error while migrating: %v", err)
		return
	}

	db.Model(&models.User{}).CreateInBatches(users, 100)
	db.Model(&models.Student{}).CreateInBatches(students, 100)
	db.Model(&models.Professor{}).CreateInBatches(professors, 100)
	db.Model(&models.Course{}).CreateInBatches(courses, 100)

	db.Model(&models.Professor{ID: uint32(1)}).Association("Courses").Append(&courses[0])
	db.Model(&models.Student{ID: uint32(1)}).Association("Courses").Append(&courses[0])
}
