package models

import "gorm.io/gorm"

// Book represents a book entity in the database.
type Books struct {
	ID        uint    `gorm:"primary key;autoIncrement" json:"id"`
	Author    *string `json:"author"`
	Title     *string `json:"title"`
	Publisher *string `json:"publisher"`
}

// MigrateBooks automatically migrates the 'Books' model to the database.
func MigrateBooks(db *gorm.DB) error {
	err := db.AutoMigrate(&Books{})
	return err
}
