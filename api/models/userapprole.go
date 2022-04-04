package models

import (
	"errors"
	"strings"

	"github.com/jinzhu/gorm"
)

func (UserAppRole) TableName() string {
	return "usersappsroles"
}

type UserAppRole struct {
	Id        int     `json:"Id"`
	Idapprole int     `json:"IdApp"`
	Iduser    int     `json:"IdUser"`
	AppRole   AppRole `gorm:"foreignKey:Idapprole"`
	User      User    `gorm:"foreignKey:Iduser"`
}

func (a *UserAppRole) SaveUserAppRole(db *gorm.DB) (*UserAppRole, error) {
	err := db.Debug().Create(&a).Error
	if err != nil {
		return &UserAppRole{}, err
	}
	return a, nil
}

func (a *UserAppRole) FindAllUsersAppsRoles(db *gorm.DB) (*[]UserAppRole, error) {
	usersappsroles := []UserAppRole{}
	err := db.Preload("AppRole.App").
		Preload("AppRole.Role").
		Preload("User").Debug().Model(&UserAppRole{}).Limit(100).Find(&usersappsroles).Error
	if err != nil {
		return &[]UserAppRole{}, err
	}
	return &usersappsroles, err
}

func (u *UserAppRole) FindUserAppRoleByID(db *gorm.DB, id uint32) (*UserAppRole, error) {
	err := db.Preload("AppRole.App").
		Preload("AppRole.Role").
		Preload("User").Debug().Model(UserAppRole{}).Where("id = ?", id).Take(&u).Error
	if err != nil {
		return &UserAppRole{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &UserAppRole{}, errors.New("UserAppRole Not Found")
	}
	return u, err
}

func (a *UserAppRole) UpdateUserAppRole(db *gorm.DB, id uint32) (*UserAppRole, error) {
	db = db.Debug().Model(&UserAppRole{}).Where("id = ?", id).Take(&UserAppRole{}).UpdateColumns(
		map[string]interface{}{
			"idapprole": a.Idapprole,
			"iduser":    a.Iduser,
		},
	)

	if db.Error != nil {
		return &UserAppRole{}, db.Error
	}

	// This is the display the updated user
	err := db.Debug().Model(&UserAppRole{}).Where("id = ?", id).Take(&a).Error
	if err != nil {
		return &UserAppRole{}, err
	}
	return a, nil
}

func (a *UserAppRole) DeleteUserAppRole(db *gorm.DB, id uint32) (int64, error) {
	db = db.Debug().Model(&UserAppRole{}).Where("id = ?", id).Take(&UserAppRole{}).Delete(&UserAppRole{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (a *UserAppRole) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if a.Idapprole == 0 {
			return errors.New("required idapprole")
		}
		if a.Iduser == 0 {
			return errors.New("required iduser")
		}
		return nil
	default:
		if a.Idapprole == 0 {
			return errors.New("required idapprole")
		}
		if a.Iduser == 0 {
			return errors.New("required user")
		}
		return nil
	}
}
