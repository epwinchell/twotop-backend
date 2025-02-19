package database

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Recipe struct {
	gorm.Model
	Name         string         `json:"name"`
	Author       Author         `json:"author" gorm:"serializer:json"`
	Description  string         `json:"description"`
	Keywords     pq.StringArray `json:"keywords" gorm:"type:text[]"`
	Yield        string         `json:"yield"`
	Category     string         `json:"category"`
	Cuisine      string         `json:"cuisine"`
	Ingredients  pq.StringArray `json:"ingredients" gorm:"type:text[]"`
	Instructions []Instruction  `json:"instructions" gorm:"serializer:json"`
}

type Instruction struct {
	Name string `json:"name"`
	Text string `json:"text"`
}

type Author struct {
	Name      string `json:"name"`
	Reference string `json:"reference"`
}
