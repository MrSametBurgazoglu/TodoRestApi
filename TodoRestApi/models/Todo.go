package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Context string `json:"context" validate:"required"`
	Done    bool   `json:"done"`
}
