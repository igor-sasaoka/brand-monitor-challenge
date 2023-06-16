package database

import (
	"github.com/igor-sasaoka/brand-monitor-front/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB 

func Init() {
    var err error
    db, err = gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database.")
    }

    db.AutoMigrate(&model.User{})
}

func DB() *gorm.DB {
    return db
}
