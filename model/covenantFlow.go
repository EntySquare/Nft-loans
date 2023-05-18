package model

import (
	"gorm.io/gorm"
)

// CovenantFlow struct
type CovenantFlow struct {
	gorm.Model
	AccountId   uint
	CovenantId  uint
	Covenant    Covenant
	Num         float64
	ReleaseDate int64
	Flag        string // 启用标志(1-启用 0-停用)
}

func (c *CovenantFlow) GetById(db *gorm.DB) error {
	return db.First(&c, c.ID).Error
}
func (c *CovenantFlow) GetByAccountIdAndReleaseDate(db *gorm.DB) ([]CovenantFlow, error) {
	flowList := make([]CovenantFlow, 0)
	db.Model(&c).Where("account_id = ? and release_date = ? ", c.AccountId, c.ReleaseDate).Find(&flowList)
	return flowList, nil
}
