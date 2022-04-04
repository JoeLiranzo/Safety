package models

import (
	"errors"
	"html"
	"strings"

	"github.com/jinzhu/gorm"
)

type Role struct {
	Id       int    `json:"Id"`
	Rolename string `json:"RoleName"`
	Roledesc string `json:"RoleDesc"`
}

func (r *Role) Prepare() {
	r.Id = 0
	r.Rolename = html.EscapeString(strings.TrimSpace(r.Rolename))
	r.Roledesc = html.EscapeString(strings.TrimSpace(r.Roledesc))
}

func (r *Role) SaveRole(db *gorm.DB) (*Role, error) {
	err := db.Debug().Create(&r).Error
	if err != nil {
		return &Role{}, err
	}
	return r, nil
}

func (a *Role) FindAllRoles(db *gorm.DB) (*[]Role, error) {
	roles := []Role{}
	err := db.Debug().Model(&Role{}).Limit(100).Find(&roles).Error
	if err != nil {
		return &[]Role{}, err
	}
	return &roles, err
}

func (r *Role) FindRoleByID(db *gorm.DB, uid uint32) (*Role, error) {
	err := db.Debug().Model(Role{}).Where("id = ?", uid).Take(&r).Error
	if err != nil {
		return &Role{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Role{}, errors.New("Role Not Found")
	}
	return r, err
}

func (r *Role) UpdateRole(db *gorm.DB, aid uint32) (*Role, error) {
	db = db.Debug().Model(&Role{}).Where("id = ?", aid).Take(&Role{}).UpdateColumns(
		map[string]interface{}{
			"rolename": r.Rolename,
			"roledesc": r.Roledesc,
		},
	)

	if db.Error != nil {
		return &Role{}, db.Error
	}

	// This is the display the updated user
	err := db.Debug().Model(&Role{}).Where("id = ?", aid).Take(&r).Error
	if err != nil {
		return &Role{}, err
	}
	return r, nil
}

func (r *Role) DeleteRole(db *gorm.DB, rid uint32) (int64, error) {
	db = db.Debug().Model(&Role{}).Where("id = ?", rid).Take(&Role{}).Delete(&Role{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (r *Role) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if r.Rolename == "" {
			return errors.New("required rolename")
		}
		if r.Roledesc == "" {
			return errors.New("required roledesc")
		}
		return nil
	default:
		if r.Rolename == "" {
			return errors.New("required rolename")
		}
		if r.Roledesc == "" {
			return errors.New("required roledesc")
		}
		return nil
	}
}
