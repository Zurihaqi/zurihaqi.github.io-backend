package model

import (
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	ID           uint         `gorm:"primaryKey"`
	Title        string       `json:"title"`
	Description  string       `json:"description"`
	Repo         string       `json:"repo"`
	Live         string       `json:"live"`
	Thumbnail    []string     `gorm:"-" json:"thumbnail"`
	Images       []Image      `json:"images" gorm:"foreignKey:ProjectID"`
	ImgAlt       []ImgAlt     `json:"img_alt" gorm:"foreignKey:ProjectID"`
	Categories   []Category   `json:"category" gorm:"many2many:project_categories"`
	Technologies []Technology `json:"tech" gorm:"many2many:project_technologies"`
}

type Image struct {
	gorm.Model
	URL       string
	ProjectID uint
}

type ImgAlt struct {
	gorm.Model
	Text      string
	ProjectID uint
}

type Category struct {
	gorm.Model
	Name     string
	Projects []Project `gorm:"many2many:project_categories"`
}

type Technology struct {
	gorm.Model
	Name     string
	Projects []Project `gorm:"many2many:project_technologies"`
}
