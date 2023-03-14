package response

type MallUserDetailResponse struct {
	UserId         int     `json:"userId"`
	NickName       string  `json:"nickName"`
	LoginName      string  `json:"loginName"`
	IntroduceSign  string  `json:"introduceSign"`
	Currency       float64 `json:"currency"`
	BankCard       int     `json:"bankCard"`
	PayPasswordMd5 string  `json:"payPasswordMd5"`
}

// 返回银行卡
type MallCardInfoRes struct {
	BankCard   string `json:"bankCard"`
	Cardhilder string `json:"cardhilder"`
}
