package model

import (
	"errors"
	"gorm.io/gorm"
)

// User struct
type User struct {
	gorm.Model
	Token            string
	RecommendCode    string //推荐码
	RecommendId      uint
	AccumulatedInput float64
	Email            string //邮箱
	EmailFlag        string //邮箱 0-未开启 1-已开启
	Password         string //支付密码
	Phone            string
	Flag             string // 启用标志(1-启用 0-停用)
	Area             string // 地区码
	GoogleKey        string // Google密钥
	GoogleFlag       string // Google验证开启状态 0-未开启 1-已开启
	BiologyKey       string // 生物验证密钥
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

// GetByPhone
//
//	@Description: 根据手机号查询用户 1 允许为空 0 不允许为空
//	@receiver u
//	@param db
//	@param isEmptyAllow
//	@return error
func (u *User) GetByPhone(db *gorm.DB, isEmptyAllow string) error {
	db = db.Where("phone = ?", u.Phone)
	if isEmptyAllow == "1" {
		return db.Limit(1).Find(&u).Error
	} else {
		return db.Take(&u).Error
	}
}

// QueryUserPwdByPhone
//
//	@Description: 根据手机号查询用户支付密码
//	@receiver u
//	@param db
//	@param isEmptyAllow
//	@return error
func (u *User) QueryUserPwdByPhone(db *gorm.DB) (pwd string, err error) {
	if err := db.Select("password").Find(&u).Error; err != nil {
		return "error", err
	}
	return pwd, nil
}

// QueryUserInfoByPhone
//
//	@Description: 根据手机号查询用户信息和账户信息 1 允许为空 0 不允许为空
//	@receiver u
//	@param db
//	@param isEmptyAllow
//	@return error
func (u *User) QueryUserInfoByPhone(db *gorm.DB) error {
	err := db.Where("phone = ?", u.Phone).Joins("Account").Limit(1).Find(&u).Error
	return err
}

// GetByRecommendId 根据推荐人Id查询用户
func (u *User) GetByRecommendId(db *gorm.DB) error {
	return db.Where("recommend_id = ?", u.RecommendId).First(&u).Error
}

// GetByRecommendCode 根据推荐码查询用户
func GetByRecommendCode(db *gorm.DB, r string) (*User, error) {
	u := User{}
	if err := db.Where("recommend_code = ?", r).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (u *User) UpdateUser(db *gorm.DB) error {
	res := db.Model(&u).Updates(u)
	if res.RowsAffected == 0 {
		return errors.New("res.RowsAffected == 0")
	}
	return nil
}

// InsertNewUser 新增用户
func (u *User) InsertNewUser(db *gorm.DB) error {
	return db.Create(u).Error
}

// UserSelectIdByToken token查询用户数据 token = "HASH"
func UserSelectIdByToken(db *gorm.DB, token string) (userId int64, tokenData string, err error) {
	err = db.Table("users").
		Select("id", "token").
		Where("token LIKE ?", token+":%").
		Row().Scan(&userId, &tokenData)
	return
}

func UserSelectIdByPhone(db *gorm.DB, phone string) (userId int64, err error) {
	err = db.Table("users").
		Select("id", "phone").
		Where("phone LIKE ?", phone+"%").
		Row().Scan(&userId)
	return
}

// UserRefreshToken
// @Description: 修改指定用户的token数据
// @param token 数据格式 <token_value:timestamp>
// @return err
func UserRefreshToken(db *gorm.DB, userId int64, token string) (err error) {
	res := db.Model(&User{}).Where("id = ?", userId).Update("token", token)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("res.RowsAffected == 0")
	}
	return nil
}

// UpdatePayPwd
//
//	@Description: 修改用户支付密码
//	@param db
//	@param userId
//	@param pwd
//	@return err
func (u *User) UpdatePayPwd(db *gorm.DB, userId int64, pwd string) error {
	res := db.Model(&u).Where("id = ?", userId).Update("password", pwd)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("UpdatePayPwd res.RowsAffected == 0")
	}
	return nil
}

// 修改用户电话号码
func (u *User) UpdateUserPhone(db *gorm.DB, userId int64, phone string) error {
	res := db.Model(&u).Where("id = ?", userId).Update("phone", phone)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("res.RowsAffected == 0")
	}
	return nil
}

// 修改用户邮箱
func (u *User) UpdateUserEmail(db *gorm.DB, userId int64, email string) error {
	res := db.Model(&u).Where("id = ?", userId).Update("email", email)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("res.RowsAffected == 0")
	}
	return nil
}

// 查询所有用户中单个邀请人Id和表中用户Id相匹配的用户
func SelectUserByRecommendId(db *gorm.DB, userId int64) (us []User, err error) {
	us = make([]User, 0)
	if err := db.Model(&User{}).Where("recommend_id = ?", userId).Find(&us).Error; err != nil {
		return nil, err
	}
	return us, nil
}

// 查询所有用户中多个邀请人Id和表中用户Id相匹配的用户
func SelectUserByRecommendList(db *gorm.DB, userIdList []int64) (us []User, err error) {
	us = make([]User, 0)
	if err := db.Model(&User{}).Where("recommend_id in ?", userIdList).Find(&us).Error; err != nil {
		return nil, err
	}
	return us, nil
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

// 查询用户Google验证是否开启
func SelectUserGoogleFlag(db *gorm.DB, uid int64) (flag string, err error) {
	err = db.Table("users").Where("id = ?", uid).Select("google_flag").Row().Scan(&flag)
	if err != nil {
		return "-1", err
	}
	return flag, err
}

// 根据手机号查询用户Google密钥
func SelectUserGoogleKeyByPhone(db *gorm.DB, phone string) (key string, err error) {
	err = db.Table("users").Where("phone = ?", phone).Select("google_key").Row().Scan(&key)
	if err != nil {
		return "", err
	}
	return key, err
}

// 更新用户Google密钥
func (u *User) UpdateUserGoogleSecret(db *gorm.DB, uid int64) error {
	res := db.Model(&u).Where("id = ?", uid).Updates(User{GoogleKey: u.GoogleKey, GoogleFlag: u.GoogleFlag})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("res.RowsAffected == 0")
	}
	return nil
}

// 查询用户生物验证是否开启
func SelectUserBioFlag(db *gorm.DB, uid int64) (flag string, err error) {
	err = db.Table("users").Where("id = ?", uid).Select("biology_flag").Row().Scan(&flag)
	if err != nil {
		return "-1", err
	}
	return flag, err
}

// 更新用户生物密钥
func (u *User) UpdateUserBioSecret(db *gorm.DB, uid int64) error {
	res := db.Model(&u).Where("id = ?", uid).Updates(User{BiologyKey: u.BiologyKey})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("res.RowsAffected == 0")
	}
	return nil
}
