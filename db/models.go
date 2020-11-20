package db

// Project struct
type Project struct {
	ID uint `json:"id" gorm:"primaryKey" gorm:"not null" gorm:"autoIncrement"`

	Name string `json:"name"  validate:"required" gorm:"not null" size:"20"`

	LiveURL string `json:"liveurl"`

	Github string `json:"gituhub"  validate:"required"  gorm:"not null"  size:"100"`

	BuiltWith []string `json:"built_with"  validate:"required" size:"20"`
}

// Projects : multiple projects
var Projects []Project
