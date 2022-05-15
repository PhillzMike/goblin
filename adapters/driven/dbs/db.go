package dbs

import (
	"github.com/Zaida-3dO/goblin/pkg/errs"
	"gorm.io/gorm"
)

var DB *gorm.DB

type SQLDBSetup interface {
	NewDB() (db *gorm.DB, err *errs.Err)
}

func InitDB(name string) {
	client, err := NewClient(name)
	if err != nil {
		panic("cannot start client")
	}

	DB, err = client.NewDB()
	if err != nil {
		panic("cannot initiate DB")
	}
}

func NewClient(name string) (SQLDBSetup, *errs.Err) {
	switch name {
	case "psql":
		return &PSQLInit{}, nil
	default:
		return nil, errs.NewInternalServerErr("invalid client name", nil)
	}
}

func GetInstance(name string) *gorm.DB {
	if DB == nil {
		InitDB(name)
	}
	return DB
}
