package api

import (
	"errors"
	"gorm.io/gorm"
	"nft-loans/model"
	"strconv"
	"time"
)

var maps map[string]string

// IncomeRunP 收益跑p
func IncomeRunP(db *gorm.DB) {
	err := db.Transaction(func(tx *gorm.DB) error {
		dateStr := time.Now().Format("2006-01-02")
		err := tx.Create(&model.PLog{DateStr: dateStr}).Error
		if err != nil {
			return errors.New("ok")
		}

		//启用标志(1-质押中 2-已完成 0-取消中)
		list, err := model.SelectMyCovenantByFlag(tx, "1")
		if err != nil {
			panic(err)
		}
		for _, v := range list {
			reward := 1.001 //奖励数量
			//AccumulatedBenefit float64    //累计收益
			//PledgeFee          float64    //质押费用
			//ReleaseFee         float64    //释放费用
			err = tx.Model(&model.Covenant{}).
				Where("id = ?", v.ID).
				Updates(map[string]interface{}{"accumulated_benefit": gorm.Expr("accumulated_benefit + ?", reward),
					"release_fee": gorm.Expr("release_fee + ?", reward)}).Error
			if err != nil {
				return err
			}

			err = tx.Model(&model.Account{}).
				Where("id = ?", v.OwnerId).
				Update("balance", gorm.Expr("balance + ?", reward)).Error
			if err != nil {
				return err
			}

			if err = tx.Create(&model.CovenantFlow{
				AccountId:   v.OwnerId,
				CovenantId:  v.ID,
				Num:         strconv.FormatFloat(reward, 'f', 2, 64),
				ReleaseDate: time.Now().Unix(),
				Flag:        "1",
			}).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		if err.Error() != "ok" {
			panic(err)
		}
	}
}
