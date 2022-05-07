package dbs

import (
	"fmt"
	"log"
	"os"

	"github.com/Zaida-3dO/goblin/config"
	"github.com/Zaida-3dO/goblin/internal/dtos"
	"github.com/Zaida-3dO/goblin/pkg/errs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PSQLInit struct{}

func (p *PSQLInit) NewDB() (*gorm.DB, error) {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("path is: %s\n", path)
	config.LoadConfig(fmt.Sprintf("%s/config", path))
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable",
		config.Cfg.DBUsername, config.Cfg.DBName, config.Cfg.DBPassword, config.Cfg.DBHost)

	gormDB, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		fmt.Printf("database connection error: %v\n", err)
		return nil, errs.NewInternalServerErr("database connection error", nil)
	}

	autoMigrate(gormDB)

	fmt.Printf("database connected successfully!\n")
	return gormDB, nil
}

func autoMigrate(db *gorm.DB) {
	db.AutoMigrate([]dtos.User{})
	db.AutoMigrate([]dtos.UserToken{})
}
