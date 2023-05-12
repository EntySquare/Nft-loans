package model

import (
	"errors"
	"gorm.io/gorm"
)

// User struct
type User struct {
	gorm.Model
	WalletAddress string
	Token         string
	Account       Account `gorm:"foreignKey:UserId"`
	Flag          string  // 启用标志(1-启用 0-停用)
}

type APIUser struct {
	Phone string
}

type UserBranch struct {
	ID uint
}

func (u *User) GetById(db *gorm.DB) error {
	return db.First(&u, u.ID).Error
}

func QueryUserCount(db *gorm.DB) (uCount int64, err error) {
	if err := db.Model(&User{}).Count(&uCount).Error; err != nil {
		return 0, err
	}
	return uCount, nil
}

// SelectAllUser 查询所有用户
func SelectAllUser(db *gorm.DB) (us []User, err error) {
	if err := db.Model(&User{}).Order("id").Find(&us).Error; err != nil {
		return nil, err
	}
	return us, nil
}

// SelectAllUserID SelectAllUser 查询所有用户ID
func SelectAllUserID(db *gorm.DB) (us []uint, err error) {
	us = make([]uint, 0)
	if err := db.Model(&User{}).Select("id").Order("id").Find(&us).Error; err != nil {
		return nil, err
	}
	return us, nil
}

// InsertNewUser 新增用户
func (u *User) InsertNewUser(db *gorm.DB) error {
	return db.Create(u).Error
}

// 更新用户Token
func (u *User) UpdateUserToken(db *gorm.DB, uid int64) error {
	res := db.Model(&u).Where("id = ?", uid).Update("token", "")
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("res.RowsAffected == 0")
	}
	return nil
}
