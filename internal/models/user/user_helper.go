package user

import (
	"errors"
	"strings"

	"github.com/dita-daystaruni/verisafe/pkg/utils"
	"gorm.io/gorm"
)

// BeforeSave is a GORM hook that is called before saving a record
func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if u.FirstName == "" || u.LastName == "" || u.NationalID == "" || u.Email == "" || u.Password == "" {
		return errors.New("fields must not be empty")
	}
	u.FirstName = strings.TrimSpace(strings.Title(u.FirstName))
	u.LastName = strings.TrimSpace(strings.Title(u.LastName))
	u.NationalID = strings.TrimSpace(u.NationalID)
	u.Gender = strings.TrimSpace(u.Gender)
	u.Address = strings.TrimSpace(u.Address)
	u.Email = strings.TrimSpace(u.Email)
	u.ProfileURL = strings.TrimSpace(u.ProfileURL)
	u.Password, err = utils.HashPassword(strings.TrimSpace(u.Password))
	if err != nil {
		return err
	}
	return nil
}

// BeforeUpdate is a GORM hook that is called before updating a record
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if u.FirstName == "" || u.LastName == "" || u.NationalID == "" || u.Email == "" || u.Password == "" {
		return errors.New("fields must not be empty")
	}
	u.FirstName = strings.TrimSpace(strings.Title(u.FirstName))
	u.LastName = strings.TrimSpace(strings.Title(u.LastName))
	u.NationalID = strings.TrimSpace(u.NationalID)
	u.Gender = strings.TrimSpace(u.Gender)
	u.Address = strings.TrimSpace(u.Address)
	u.Email = strings.TrimSpace(u.Email)
	u.ProfileURL = strings.TrimSpace(u.ProfileURL)
	u.Password = strings.TrimSpace(u.Password)
	return nil
}
