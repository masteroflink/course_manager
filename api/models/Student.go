package models

import (
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Student struct {
	ID               uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name             Name      `gorm:"embedded;embeddedPrefix:name_" json:"name"`
	Address          Address   `gorm:"embedded;embeddedPrefix:address_" json:"address"`
	Courses          []*Course `gorm:"many2many:course_students" json:"courses"`
	Email            string    `gorm:"size:256;not null;unique" json:"email"`
	Phone            string    `gorm:"size:256" json:"phone"`
	GPA              float32   `json:"gpa"`
	Credits          float32   `json:"credits"`
	AttemptedCredits float32   `json:"attempted_credits"`
	DegreeLevel      string    `gorm:"size:256" json:"degree_level"`
	FieldOfStudy     string    `gorm:"size:256" json:"field_of_study"`
	CreatedAt        time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (s *Student) Prepare() {
	CleanName(&s.Name)
	CleanAddress(&s.Address)
	s.Email = strings.TrimSpace(s.Email)
	s.Phone = strings.TrimSpace(s.Phone)
	s.DegreeLevel = strings.TrimSpace(s.DegreeLevel)
	s.FieldOfStudy = strings.TrimSpace(s.FieldOfStudy)

	s.CreatedAt = time.Now().UTC()
	s.UpdatedAt = time.Now().UTC()
}

func (s *Student) Validate() error {
	if err := ValidateName(&s.Name); err != nil {
		return err
	}

	if err := ValidateEmail(s.Email); err != nil {
		return err
	}

	return nil
}

// When does SaveStudent get called?
func (s *Student) SaveStudent(db *gorm.DB) (*Student, error) {
	err := db.Debug().Create(&s).Error
	if err != nil {
		return &Student{}, err
	}

	return s, nil
}

func (s *Student) GetAllStudents(db *gorm.DB) (*[]Student, error) {
	students := []Student{}
	err := db.Debug().Preload("Courses").Limit(100).Find(&students).Error
	if err != nil {
		return &[]Student{}, err
	}
	return &students, err
}

func (s *Student) GetStudent(db *gorm.DB, sid uint32) (*Student, error) {
	student := Student{}
	err := db.Debug().Preload("Courses").First(&student, sid).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &Student{}, errors.New("Student Not Found")
	}
	if err != nil {
		return &Student{}, err
	}
	return s, err
}

func (s *Student) UpdateStudent(db *gorm.DB, sid uint32) (*Student, error) {
	db = db.Debug().First(&Student{}, sid).Updates(*s)
	if db.Error != nil {
		return &Student{}, db.Error
	}

	err := db.Debug().Model(Student{ID: sid}).Preload("Courses").Take(s).Error
	if err != nil {
		return &Student{}, err
	}

	return s, nil
}

func (s *Student) DeleteStudent(db *gorm.DB, sid uint32) (int64, error) {
	err := db.Debug().Model(&Student{ID: sid}).Association("Courses").Clear()
	if err != nil {
		return 0, err
	}

	db = db.Debug().Delete(&Student{ID: sid})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
