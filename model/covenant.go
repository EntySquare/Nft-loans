package model

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

// Covenant struct
type Covenant struct {
	gorm.Model
	Name              string     //产品名字
	AccumulatedIncome float64    //累计收益
	Duration          string     //时长
	StartTime         *time.Time //开始执行时间
	Flag              string     // 启用标志(1-启用 0-未启用)
	OwnerId           uint       //用户id
	Owner             User       //用户
}

func NewCovenant(id int64) Covenant {
	return Covenant{Model: gorm.Model{ID: uint(id)}}
}

func (c *Covenant) GetById(db *gorm.DB) error {
	return db.First(&c, c.ID).Error
}

func (c *Covenant) UpdateCovenant(db *gorm.DB) error {
	return db.Model(&c).Updates(c).Error
}
func (c *Covenant) InsertNewCovenant(db *gorm.DB) error {
	return db.Create(c).Error
}

// SelectMyCovenant
//
//	@Description:
//	@param db
//	@return userId
//	@return err

func (c *Covenant) SelectMyCovenant(db *gorm.DB) (cs []Covenant, err error) {
	cs = make([]Covenant, 0)
	err = db.Model(&Covenant{}).Where("owner_id = ?", c.OwnerId).Find(&cs).Error
	return cs, err
}

// GetAllOnTimePower 统计全网算力
func (c *Covenant) GetAllOnTimePower(db *gorm.DB) (int64, error) {
	var allPower sql.NullInt64
	err := db.Model(&c).Select("sum(power)").Where("flag = '1'").Scan(&allPower).Error
	if err != nil {
		return 0, err
	}
	if allPower.Valid {
		return allPower.Int64, nil
	} else {
		return 0, err
	}
}

// GetUserTotalPower 统计用户总算力
func (c *Covenant) GetUserTotalPower(db *gorm.DB, uid uint) (int64, error) {
	var totalPower sql.NullInt64
	err := db.Model(&c).Select("sum(power)").Where("owner_id = ? ", uid).Scan(&totalPower).Error
	if err != nil {
		return 0, err
	}
	if totalPower.Valid {
		return totalPower.Int64, nil
	} else {
		return 0, err
	}
}
