package api

import (
	"errors"
	"gorm.io/gorm"
	"nft-loans/contracts"
	"nft-loans/model"
	"strconv"
	"time"
)

var maps map[string]string

// IncomeRunP 收益跑p
func IncomeRunP(db *gorm.DB) {
	//更新质押记录数据
	//ChainUnix
	//更新所有质押记录
	err := RenewAddressLoansList(db)
	if err != nil {
		return
	}
	//跑p查询全部质押记录
	err = db.Transaction(func(tx *gorm.DB) error {
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
			//区块未确认 跳过
			if v.ChainUnix == 0 {
				continue
			}
			//reward := 1.001 //奖励数量
			nftId, err := strconv.Atoi(v.PledgeId)
			if err != nil {
				return err
			}
			_, reward, _ := GetInterestRate(nftId, tx)
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
				Num:         strconv.FormatInt(reward, 10),
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

// RenewAddressLoansList 更新所有用户的质押记录
func RenewAddressLoansList(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		uList, err := model.SelectAllUser(tx)
		if err != nil {
			return err
		}
		//循环平台钱包
		for _, user := range uList {
			loansList, err := contracts.GetAddressLoansList(user.WalletAddress)
			if err != nil {
				return err
			}
			if loansList == nil {
				continue
			}
			if len(loansList) == 0 {
				continue
			}
			if err = RenewUserLoansList(tx, user, loansList); err != nil {
				return err
			}
		}
		return nil
	})
}

func RenewUserLoansList(db *gorm.DB, user model.User, loansList []contracts.NGTTokenData) error {
	for _, v := range loansList {
		//fmt.Println(v.Flag.String())
		if v.Flag.Int64() != 1 {
			continue
		}
		err := db.Model(&model.Covenant{}).
			Where("pledge_id = ? and flag = '1' and owner_id = ?",
				v.TokenId.String(), user.ID).
			Updates(map[string]interface{}{
				"chain_unix": v.ReceivedTime.Int64(),
			}).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func GetInterestRate(id int, tx *gorm.DB) (string, int64, float64) {
	//var interest InterestRate
	//Token ID 1-500 朱雀 AAA级 0.7%日        NGT 1000分比，1000*0.007=7 单利。
	//Token ID 501-3500  白虎  AA级  0.6%日   NGT 1000分比，1000*0.006=6 单利。
	//Token ID 3501-10000玄武  A级别 0.5%日   NGT 1000分比，1000*0.005=5 单利。
	name := ""
	num := int64(0)
	f64 := 0.0
	var ni = model.NftInfo{}
	if id >= 1 && id <= 500 {
		ni.TypeNum = 7
	} else if id >= 501 && id <= 3500 {
		ni.TypeNum = 6
		//name = "白虎"
		//num = 6
		//f64 = 0.006
	} else if id >= 3501 && id <= 10000 {
		ni.TypeNum = 5
		//name = "玄武"
		//num = 5
		//f64 = 0.005
	}
	err := ni.GetByTypeNum(tx)
	if err != nil {
		return "", 0, 0
	}
	name = ni.Name
	num = ni.TypeNum
	f64 = ni.DayRate
	return name, num, f64
}

func SelectChainData(db *gorm.DB, userId uint) error {
	fun := func() error {
		var user = model.User{}
		err := db.First(&user, userId).Error
		if err != nil {
			return err
		}
		loansList, err := contracts.GetAddressLoansList(user.WalletAddress)
		if err != nil {
			return err
		}
		if loansList == nil {
			return errors.New("loansList == nil")
		}
		if len(loansList) == 0 {
			return errors.New("len(loansList) == 0")
		}
		if err = RenewUserLoansList(db, user, loansList); err != nil {
			return err
		}
		return nil
	}
	i := 0
	for {
		if i == 20 {
			return nil
		}
		err := fun()
		if err == nil {
			return nil
		}
		i++
		time.Sleep(time.Second * 10)
	}
}
