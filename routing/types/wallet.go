package types

type WithdrawNgtReq struct {
	Num     int64  `json:"num"`
	Address string `json:"address"`
	Hash    string `json:"hash"`
	Chain   string `json:"chain"`
}

type DepositNgtReq struct {
	Num     int64  `json:"num"`
	Address string `json:"address"`
	Hash    string `json:"hash"`
	Chain   string `json:"chain"`
}
