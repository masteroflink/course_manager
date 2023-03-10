package models

import (
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Course struct {
	ID           uint32       `gorm:"primary_key;auto_increment" json:"id"`
	College      string       `gorm:"primary_key;size:50;not null" json:"college"`
	Department   string       `gorm:"primary_key;size:50;not null" json:"department"`
	CourseNumber string       `gorm:"primary_key;size:20;not null" json:"course_number"`
	MaxCapacity  int          `gorm:"not null" json:"max_capacity"`
	CreatedAt    time.Time    `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time    `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Professors   []*Professor `gorm:"many2many:course_professors" json:"professors"`
	Students     []*Student   `gorm:"many2many:course_students" json:"students"`
}

func (c *Course) Prepare() {
	c.Department = strings.TrimSpace(c.Department)
	c.CourseNumber = strings.TrimSpace(c.CourseNumber)
	c.College = strings.TrimSpace(c.College)

	c.CreatedAt = time.Now().UTC()
	c.UpdatedAt = time.Now().UTC()
}

func (c *Course) Validate() error {
	if len(c.Department) <= 0 {
		return errors.New("Course Department required.")
	}

	if len(c.CourseNumber) <= 0 {
		return errors.New("Course Number required.")
	}

	if len(c.CourseNumber) <= 0 {
		return errors.New("Course College required.")
	}

	return nil
}

func (c *Course) SaveCourse(db *gorm.DB) (*Course, error) {
	err := db.Debug().Create(&c).Error
	if err != nil {
		return &Course{}, err
	}

	return c, nil
}

func (c *Course) GetAllCourses(db *gorm.DB) (*[]Course, error) {
	courses := []Course{}
	err := db.Debug().Preload("Students").Preload("Professors").Limit(100).Find(&courses).Error

	if err != nil {
		return &[]Course{}, err
	}

	return &courses, nil
}

func (c *Course) GetCourse(db *gorm.DB, cid uint32) (*Course, error) {
	course := Course{}
	err := db.Debug().Preload("Students").Preload("Professors").First(&course, cid).Error
	if err != nil {
		return &Course{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &Course{}, errors.New("Course Not Found")
	}

	return &course, nil
}

func (c *Course) UpdateCourse(db *gorm.DB, cid uint32) (*Course, error) {
	db = db.Debug().First(&Course{}, cid).Updates(*c)

	if db.Error != nil {
		return &Course{}, db.Error
	}

	if err := db.Debug().Model(Course{ID: cid}).Preload("Students").Preload("Professors").Take(c).Error; err != nil {
		return &Course{}, err
	}

	return c, nil
}

func (s *Course) DeleteCourse(db *gorm.DB, cid uint32) (int64, error) {

	err := db.Debug().Model(&Course{ID: cid}).Association("Students").Clear()
	if err != nil {
		return 0, err
	}

	err = db.Debug().Model(&Course{ID: cid}).Association("Professors").Clear()
	if err != nil {
		return 0, err
	}

	db = db.Debug().Delete(&Course{ID: cid})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (c *Course) EnrollStudent(db *gorm.DB, cid uint32, sid uint32) (*Course, error) {
	err := db.Debug().Model(&Course{ID: cid}).Association("Students").Append(&Student{ID: sid})
	if err != nil {
		return &Course{}, err
	}

	course := Course{}
	err = db.Debug().First(&Course{}, cid).Preload("Students").Take(&course).Error
	if err != nil {
		return &Course{}, err
	}

	return &course, nil
}

func (c *Course) RemoveStudent(db *gorm.DB, cid uint32, sid uint32) (*Course, error) {
	err := db.Debug().Model(&Course{ID: cid}).Association("Students").Delete(Student{ID: sid})
	if err != nil {
		return &Course{}, err
	}

	course := Course{}
	err = db.Debug().First(&Course{}, cid).Preload("Students").Take(&course).Error
	if err != nil {
		return &Course{}, err
	}

	return &course, nil
}

func (c *Course) AssignProfessor(db *gorm.DB, cid uint32, pid uint32) (*Course, error) {
	err := db.Debug().Model(&Course{ID: cid}).Association("Professors").Append(&Professor{ID: pid})
	if err != nil {
		return &Course{}, err
	}

	course := Course{}
	err = db.Debug().First(&Course{}, cid).Preload("Professors").Take(&course).Error
	if err != nil {
		return &Course{}, err
	}

	return &course, nil
}

func (c *Course) RemoveProfessor(db *gorm.DB, cid uint32, sid uint32) (*Course, error) {
	err := db.Debug().Model(&Course{ID: cid}).Association("Professors").Delete(Professor{ID: sid})
	if err != nil {
		return &Course{}, err
	}

	course := Course{}
	err = db.Debug().First(&Course{}, cid).Preload("Professors").Take(&course).Error
	if err != nil {
		return &Course{}, err
	}

	return &course, nil
}
