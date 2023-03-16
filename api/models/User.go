package models

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name      Name      `gorm:"embedded;embeddedPrefix:name_" json:"name"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave() error {
	hashedPassword, err := HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return err
}

func (u *User) Prepare() {
	CleanName(&u.Name)
	u.Email = strings.TrimSpace(u.Email)
	u.CreatedAt = time.Now().UTC()
	u.UpdatedAt = time.Now().UTC()
}

func (u *User) Validate(action string) error {
	switch action {
	case "update":
		if err := ValidateName(&u.Name); err != nil {
			return err
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := ValidateEmail(u.Email); err != nil {
			return err
		}
		return nil
	case "login":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	default:
		if err := ValidateName(&u.Name); err != nil {
			return err
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := ValidateEmail(u.Email); err != nil {
			return err
		}
		return nil
	}
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	err := db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) GetAllUsers(db *gorm.DB) (*[]User, error) {
	users := []User{}
	err := db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

func (u *User) GetUser(db *gorm.DB, uid uint32) (*User, error) {
	user := User{}
	err := db.Debug().First(&user, uid).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &User{}, errors.New("User Not Found")
	}

	if err != nil {
		return &User{}, err
	}

	return &user, err
}

func (u *User) UpdateUser(db *gorm.DB, uid uint32) (*User, error) {
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().First(&User{}, uid).Updates(*u)

	if db.Error != nil {
		return &User{}, db.Error
	}

	err = db.Debug().Model(User{ID: uid}).Take(u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (s *User) DeleteUser(db *gorm.DB, uid uint32) (int64, error) {
	db = db.Debug().Delete(&User{ID: uid})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
