package model

import "github.com/jinzhu/gorm"

type Class struct {
	gorm.Model
	ClassName int `json:"className"`
	Students  []Student
}

type Student struct {
	gorm.Model
	StudentName string    `json:"studentName"`
	ClassId     uint      `json:"classId"`
	IDCard      IDCard    `json:"IDCard"`
	Teachers    []Teacher `gorm:"many2many:student_teachers;"`
}

type IDCard struct {
	gorm.Model
	Num       int  `json:"num"`
	StudentId uint `json:"studentId"`
}

type Teacher struct {
	gorm.Model
	TeacherName string    `json:"teacherName"`
	Students    []Student `gorm:"many2many:student_teachers;"`
}
