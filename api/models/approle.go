package models

import (
	"errors"
	"strings"

	"github.com/jinzhu/gorm"
)

func (AppRole) TableName() string {
	return "appsroles"
}

type AppRole struct {
	Id     int  `json:"Id"`
	Idapp  int  `json:"IdApp"`
	Idrole int  `json:"IdRole"`
	App    App  `gorm:"foreignKey:Idapp"`
	Role   Role `gorm:"foreignKey:Idrole"`
}

func (a *AppRole) SaveAppRole(db *gorm.DB) (*AppRole, error) {
	err := db.Debug().Create(&a).Error
	if err != nil {
		return &AppRole{}, err
	}
	return a, nil
}

func (a *AppRole) FindAllAppsRoles(db *gorm.DB) (*[]AppRole, error) {
	appsroles := []AppRole{}
	err := db.Preload("App").Preload("Role").Debug().Model(&AppRole{}).Limit(100).Find(&appsroles).Error
	if err != nil {
		return &[]AppRole{}, err
	}
	return &appsroles, err
}

func (a *AppRole) FindAppRoleByID(db *gorm.DB, uid uint32) (*AppRole, error) {
	err := db.Preload("App").Preload("Role").Debug().Model(AppRole{}).Where("id = ?", uid).Take(&a).Error
	if err != nil {
		return &AppRole{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &AppRole{}, errors.New("AppRole Not Found")
	}
	return a, err
}

func (a *AppRole) UpdateAppRole(db *gorm.DB, id uint32) (*AppRole, error) {
	db = db.Debug().Model(&AppRole{}).Where("id = ?", id).Take(&AppRole{}).UpdateColumns(
		map[string]interface{}{
			"idapp":  a.Idapp,
			"idrole": a.Idrole,
		},
	)

	if db.Error != nil {
		return &AppRole{}, db.Error
	}

	// This is the display the updated user
	err := db.Debug().Model(&AppRole{}).Where("id = ?", id).Take(&a).Error
	if err != nil {
		return &AppRole{}, err
	}
	return a, nil
}

func (a *AppRole) DeleteAppRole(db *gorm.DB, id uint32) (int64, error) {
	db = db.Debug().Model(&AppRole{}).Where("id = ?", id).Take(&AppRole{}).Delete(&AppRole{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (a *AppRole) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if a.Idapp == 0 {
			return errors.New("required idapp")
		}
		if a.Idrole == 0 {
			return errors.New("required idrole")
		}
		return nil
	default:
		if a.Idapp == 0 {
			return errors.New("required idapp")
		}
		if a.Idrole == 0 {
			return errors.New("required idrole")
		}
		return nil
	}
}
