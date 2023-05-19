package model

import (
	"gorm.io/gorm"
)

// Transactions struct
type Transactions struct {
	gorm.Model
	Hash      string //交易哈希
	Status    string //交易状态
	ChainName string //公链名
	Flag      string // // 启用标志(1-质押中 2-已完成 0-取消中)
}

func NewTransactions(id int64) Transactions {
	return Transactions{Model: gorm.Model{ID: uint(id)}}
}

func (txs *Transactions) GetById(db *gorm.DB) error {
	return db.First(&txs, txs.ID).Error
}

func (txs *Transactions) UpdateTransactions(db *gorm.DB) error {
	return db.Model(&txs).Updates(txs).Error
}
func (txs *Transactions) InsertNewTransactions(db *gorm.DB) error {
	return db.Create(txs).Error
}
