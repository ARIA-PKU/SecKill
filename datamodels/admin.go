package datamodels

import "github.com/jinzhu/gorm"

type Admin struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"varchar(11);not null;unique"`
	Password  string `gorm:"size:255;not null"`
	Authority string  `gorm:"type:varchar(11)"`
}
