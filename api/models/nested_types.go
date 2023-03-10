package models

import (
	"errors"
	"strings"
)

type Name struct {
	First  string `gorm:"size:100;not null" json:"first"`
	Middle string `gorm:"size:100" json:"middle"`
	Last   string `gorm:"size:100;not null" json:"last"`
}

type Address struct {
	Raw          string `gorm:"size:100" json:"raw"`
	StreetNumber string `gorm:"size:100" json:"street_number"`
	StreetName   string `gorm:"size:100" json:"street_name"`
	Unit         string `gorm:"size:100" json:"unit"`
	City         string `gorm:"size:100" json:"city"`
	State        string `gorm:"size:100" json:"state"`
	Country      string `gorm:"size:100" json:"country"`
	CountryCode  string `gorm:"size:100" json:"country_code"`
	PostalCode   string `gorm:"size:100" json:"postal_code"`
}

func CleanName(n *Name) {
	n.First = strings.TrimSpace(n.First)
	n.Middle = strings.TrimSpace(n.Middle)
	n.Last = strings.TrimSpace(n.Last)
}

func ValidateName(n *Name) error {
	if len(n.First) <= 0 {
		return errors.New("First Name required")
	}

	if len(n.Last) <= 0 {
		return errors.New("Last Name required")
	}
	return nil
}

func CleanAddress(a *Address) {
	a.Raw = strings.TrimSpace(a.Raw)
	a.StreetNumber = strings.TrimSpace(a.StreetNumber)
	a.StreetName = strings.TrimSpace(a.StreetName)
	a.Unit = strings.TrimSpace(a.Unit)
	a.City = strings.TrimSpace(a.City)
	a.State = strings.TrimSpace(a.State)
	a.Country = strings.TrimSpace(a.Country)
	a.CountryCode = strings.TrimSpace(a.CountryCode)
	a.PostalCode = strings.TrimSpace(a.PostalCode)
}
