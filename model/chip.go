package model

import "github.com/jinzhu/gorm"

type ChipType struct {
	gorm.Model
	Name  string `gorm:"type:varchar(20);not null"`
	Chips []Chip
	//Chips []Chip `gorm:"foreignKey:ChipId"` //自定义外键
}

type Chip struct {
	gorm.Model
	Frame      string
	Length     uint
	MaxNum     uint
	MinNum     uint
	ChipTypeID uint `json:"chipTypeID"`
}
