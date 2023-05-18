package model

import (
	"gorm.io/gorm"
)

// Account struct
type Account struct {
	gorm.Model
	UserId  uint
	Balance float64
	Flag    string // 启用标志(1-启用 0-停用)
}

func (u *Account) GetById(db *gorm.DB) error {
	return db.First(&u, u.ID).Error
}
func (a *Account) GetByUserId(db *gorm.DB) error {
	return db.Model(&a).Where("user_id = ? ", a.UserId).Take(&a).Error
}
