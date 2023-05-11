package model

import (
	"gorm.io/gorm"
)

// Flow struct
type Flow struct {
	gorm.Model
	CovenantId uint
	Covenant   Covenant
	Balance    float64
	Flag       string // 启用标志(1-启用 0-停用)
}

func (u *Flow) GetById(db *gorm.DB) error {
	return db.First(&u, u.ID).Error
}
