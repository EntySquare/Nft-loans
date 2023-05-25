package types

type WithdrawNgtReq struct {
	Address string  `json:"address"`
	Num     float64 `json:"num"`
	Hash    string  `json:"hash"`
	Chain   string  `json:"chain"`
}

type WithdrawNgtResp struct {
	Msg string `json:"msg"`
}

type DepositNgtReq struct {
	Address string  `json:"address"`
	Num     float64 `json:"num"`
	Hash    string  `json:"hash"`
	Chain   string  `json:"chain"`
}

type DepositNgtResp struct {
	Msg string `json:"msg"`
}
