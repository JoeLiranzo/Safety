package models

import (
	"errors"
	"html"
	"strings"

	"github.com/jinzhu/gorm"
)

type App struct {
	Id      uint64 `json:"Id"`
	Appname string `json:"AppName"`
	Appdesc string `json:"AppDesc"`
}

func (a *App) Prepare() {
	a.Id = 0
	a.Appname = html.EscapeString(strings.TrimSpace(a.Appname))
	a.Appdesc = html.EscapeString(strings.TrimSpace(a.Appdesc))
}

func (a *App) SaveApp(db *gorm.DB) (*App, error) {
	err := db.Debug().Create(&a).Error
	if err != nil {
		return &App{}, err
	}
	return a, nil
}

func (a *App) FindAllApps(db *gorm.DB) (*[]App, error) {
	apps := []App{}
	err := db.Debug().Model(&App{}).Limit(100).Find(&apps).Error
	if err != nil {
		return &[]App{}, err
	}
	return &apps, err
}

func (a *App) FindAppByID(db *gorm.DB, uid uint32) (*App, error) {
	err := db.Debug().Model(App{}).Where("id = ?", uid).Take(&a).Error
	if err != nil {
		return &App{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &App{}, errors.New("App Not Found")
	}
	return a, err
}

func (a *App) UpdateApp(db *gorm.DB, aid uint32) (*App, error) {
	db = db.Debug().Model(&App{}).Where("id = ?", aid).Take(&App{}).UpdateColumns(
		map[string]interface{}{
			"appname": a.Appname,
			"appdes":  a.Appdesc,
		},
	)

	if db.Error != nil {
		return &App{}, db.Error
	}

	// This is the display the updated user
	err := db.Debug().Model(&App{}).Where("id = ?", aid).Take(&a).Error
	if err != nil {
		return &App{}, err
	}
	return a, nil
}

func (a *App) DeleteApp(db *gorm.DB, uid uint32) (int64, error) {
	db = db.Debug().Model(&App{}).Where("id = ?", uid).Take(&App{}).Delete(&App{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (a *App) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if a.Appname == "" {
			return errors.New("required appname")
		}
		if a.Appdesc == "" {
			return errors.New("required appdesc")
		}
		return nil
	default:
		if a.Appname == "" {
			return errors.New("required appname")
		}
		if a.Appdesc == "" {
			return errors.New("required appdesc")
		}
		return nil
	}
}
