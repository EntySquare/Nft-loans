package types

type SengMsg struct {
	Msg   string `json:"msg"`
	Phone string `json:"phone"`
	Area  string `json:"area"`
}

// 登录注册请求参数
type LoginAndRegisterReq struct {
	WalletAddress string `json:"wallet_address"`
	Tag           string `json:"tag"`
	Code          string `json:"code"`
}

// 修改支付密码请求参数
type UpdatePwdReq struct {
	WalletAddress string `json:"wallet_address"`
	NewPwd        string `json:"new_pwd"`
	OldPwd        string `json:"old_pwd"`
}

type MyInvestmentResp struct {
	UID                    string               `json:"uid"`
	Level                  int64                `json:"level"`
	AccumulatedPledgeCount int64                `json:"accumulated_pledge_count"`
	InvestmentCount        int64                `json:"investment_count"`
	InvestmentAddress      string               `json:"investment_address"`
	InvestmentUsers        []InvestmentUserInfo `json:"investment_users"`
}
type InvestmentUserInfo struct {
	UID         string `json:"uid"`
	Level       int64  `json:"level"`
	PledgeCount int64  `json:"pledge_count"`
}
