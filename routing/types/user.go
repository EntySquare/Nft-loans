package types

type SengMsg struct {
	Msg   string `json:"msg"`
	Phone string `json:"phone"`
	Area  string `json:"area"`
}

// 登录注册请求参数
type LoginAndRegisterReq struct {
	Phone         string `json:"phone"`
	Tag           string `json:"tag"`
	Code          string `json:"code"`
	RecommendCode string `json:"recommend_code"`
	Area          string `json:"area"`
	BiologyKey    string `json:"biology_key"`
}

// 用户信息返回数据
type UserInfoRes struct {
	Phone         string `json:"phone"`
	Email         string `json:"email"`
	EmailFlag     string `json:"email_flag"` //邮箱是否开启 0-未开启 1-已开启
	PwdFlag       string `json:"pwd_flag"`
	RecommendCode string `json:"recommend_code"`
	GoogleFlag    string `json:"google_flag"`
}

// 修改支付密码请求参数
type UpdatePwdReq struct {
	Phone  string `json:"phone"`
	NewPwd string `json:"new_pwd"`
	OldPwd string `json:"old_pwd"`
}

// 用户名下成员返回数据
type UserMember struct {
	Phone        string   `json:"phone"`
	FirstMember  []Member `json:"first_member"`
	SecondMember []Member `json:"second_member"`
}

type Member struct {
	ID    uint   `json:"id"`
	Phone string `json:"phone"`
}

// FromRealNameAttestationReq 实名认证数据接收
type FromRealNameAttestationReq struct {
	Name   string `json:"name"`
	IDCode string `json:"id_code"`
}

type GoogleAuth struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
	Key   string `json:"key"`
}

type BiologyAuth struct {
	Phone      string `json:"phone"`
	DeviceCode string `json:"device_code"`
	Key        string `json:"key"`
}

type SendMailboxCodeRes struct {
	Tag string `json:"tag"` // 验证码tag
}

type BindingEmailReq struct {
	Tag   string `json:"tag"`   // 验证码tag
	Code  string `json:"code"`  // 邮箱验证码
	Email string `json:"email"` // 邮箱
}
