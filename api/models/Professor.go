package models

import (
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Professor struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name      Name      `gorm:"embedded;embeddedPrefix:name_" json:"name"`
	Address   Address   `gorm:"embedded;embeddedPrefix:address_" json:"address"`
	Email     string    `gorm:"size:50;unique;not null" json:"email"`
	Position  string    `gorm:"size:50;not null" json:"position"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Courses   []*Course `gorm:"many2many:course_professors" json:"courses"`
}

func (p *Professor) Prepare() {
	CleanName(&p.Name)
	CleanAddress(&p.Address)
	p.Email = strings.TrimSpace(p.Email)
	p.Position = strings.TrimSpace(p.Position)

	p.CreatedAt = time.Now().UTC()
	p.UpdatedAt = time.Now().UTC()
}

func (p *Professor) Validate() error {
	if err := ValidateName(&p.Name); err != nil {
		return err
	}

	if err := ValidateEmail(p.Email); err != nil {
		return err
	}
	return nil
}

func (p *Professor) SaveProfessor(db *gorm.DB) (*Professor, error) {
	err := db.Debug().Create(&p).Error
	if err != nil {
		return &Professor{}, err
	}

	return p, nil
}

func (p *Professor) GetAllProfessors(db *gorm.DB) (*[]Professor, error) {
	professors := []Professor{}
	err := db.Debug().Preload("Courses").Limit(100).Find(&professors).Error

	if err != nil {
		return &[]Professor{}, err
	}

	return &professors, nil
}

func (p *Professor) GetProfessor(db *gorm.DB, pid uint32) (*Professor, error) {
	professor := Professor{}
	err := db.Debug().Preload("Courses").First(&professor, pid).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &Professor{}, errors.New("Professor Not Found")
	}

	if err != nil {
		return &Professor{}, err
	}

	return &professor, nil
}

func (p *Professor) UpdateProfessor(db *gorm.DB, pid uint32) (*Professor, error) {
	db = db.Debug().First(&Professor{}, pid).Updates(*p)

	if db.Error != nil {
		return &Professor{}, db.Error
	}

	if err := db.Debug().Model(Professor{ID: pid}).Preload("Courses").Take(p).Error; err != nil {
		return &Professor{}, err
	}

	return p, nil
}

func (s *Professor) DeleteProfessor(db *gorm.DB, pid uint32) (int64, error) {
	err := db.Debug().Model(&Professor{ID: pid}).Association("Courses").Clear()
	if err != nil {
		return 0, err
	}

	db = db.Debug().Delete(&Professor{ID: pid})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
