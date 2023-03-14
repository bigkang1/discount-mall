package request

type UpdateCurrencyRecordStatus struct {
	UserCurrencyRecordId int `json:"userCurrencyRecordId"`
	Status               int `json:"status"`
}

type AccessUserCurrency struct {
	UserCurrencyRecordId int `json:"userCurrencyRecordId"`
}
