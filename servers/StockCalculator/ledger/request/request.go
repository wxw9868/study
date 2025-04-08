package request

type LedgerId struct {
	Id uint `uri:"id" form:"id" json:"id" binding:"required"`
}

// Ledger 创建账本请求
type Ledger struct {
	UserId     uint    `json:"user_id"`
	Name       string  `json:"name" binding:"required"`
	Date       string  `json:"date"`
	CurrencyId uint    `json:"currency_id" binding:"required"`
	Amount     float64 `json:"amount" binding:"required"`
}

// AssetRecord 记录资产请求
type AssetRecord struct {
	LedgerId uint    `json:"ledger_id" binding:"required"`
	Amount   float64 `json:"amount" binding:"required"`
	Date     string  `json:"date"`
}
