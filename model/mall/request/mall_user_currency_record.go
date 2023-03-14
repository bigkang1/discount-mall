package request

type AddUserCurrencyRecord struct {
	CurrencyAmount float64 `json:"currencyAmount"`
	//CurrencyType   int     `json:"currencyType"`
}
