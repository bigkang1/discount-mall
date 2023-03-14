package request

// 用户注册
type RegisterUserParam struct {
	NickNam    string `json:"nickName"`
	LoginName  string `json:"loginName"`
	Password   string `json:"password"`
	InviteCode string `json:"inviteCode"`
}

// 更新用户信息
type UpdateUserInfoParam struct {
	NickName      string `json:"nickName"`
	RawPassword   string `json:"rawPassword"`
	PasswordMd5   string `json:"passwordMd5"`
	IntroduceSign string `json:"introduceSign"`
}

type UserLoginParam struct {
	LoginName   string `json:"loginName"`
	PasswordMd5 string `json:"passwordMd5"`
}

// 更新用户信息
type UpdateUserBankCard struct {
	BankCard    string `json:"bankCard"`
	Cardhilder  string `json:"cardhilder"`
	PayPassword string `json:"payPassword"`
}

type ValidateUserPayPassword struct {
	PayPasswordMd5 string `json:"payPasswordMd5"`
}
