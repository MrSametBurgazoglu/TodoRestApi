package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Context string `json:"context"`
	Done    bool   `json:"done"` //is this task completed
}
