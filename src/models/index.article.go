package models

import (
	"gorm.io/gorm"
  )

func Migrate(conn *gorm.DB){
	conn.AutoMigrate(&Articles{})
}