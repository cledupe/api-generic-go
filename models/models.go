package models

import "gorm.io/gorm"

type DayTask struct {
	gorm.Model
	Task string `json:"task" gorm:"text;not null;default:null`
	Date string `json:"date" gorm:"text;not null;default:null`
}
