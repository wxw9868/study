package server

import "fmt"

const (
	CommissionRate = 0.03 / 100  // 沪深两市成交金额的佣金费率 0.3% (起点5元)
	StampDuty      = 0.05 / 100  // 印花税为成交金额的 0.05% （卖出时收取）
	TransferFee    = 0.001 / 100 // 过户费为成交金额的 0.001% （卖出时收取）
	// 温馨提示：佣金、过户费双向收取，印花税向卖方收取。
)

// 新增交易方向类型
type TradeDirection int

const (
	Buy  TradeDirection = iota // 买入
	Sell                       // 卖出
)

type Stock struct {
	CommissionRate float64 // 佣金费率
	StampDuty      float64 // 印花税
	TransferFee    float64 // 过户费
}

func NewStock(commissionRate float64) *Stock {
	if commissionRate <= 0 {
		commissionRate = CommissionRate
	}
	return &Stock{
		CommissionRate: commissionRate,
		StampDuty:      StampDuty,
		TransferFee:    TransferFee,
	}
}

func (s *Stock) CalculateCommission(amount float64) float64 {
	commission := max(amount*s.CommissionRate, 5)
	return commission + 0.01
}

func (s *Stock) CalculateStampDuty(amount float64, direction TradeDirection) float64 {
	if direction == Sell {
		return amount * s.StampDuty
	}
	return 0
}

func (s *Stock) CalculateTransferFee(amount float64) float64 {
	return amount * s.TransferFee
}

func (s *Stock) CalculateTotalFee(amount float64, direction TradeDirection) float64 {
	commission := s.CalculateCommission(amount)
	transferFee := s.CalculateTransferFee(amount)
	stampDuty := s.CalculateStampDuty(amount, direction)
	return commission + stampDuty + transferFee
}

func (s *Stock) CalculateTotalFeeByCount(price float64, count float64, direction TradeDirection) float64 {
	commission := s.CalculateCommission(price * count)
	stampDuty := s.CalculateStampDuty(price*count, direction)
	transferFee := s.CalculateTransferFee(price * count)
	return commission + stampDuty + transferFee
}

type Transaction struct {
	Stock       *Stock
	Price       float64
	Count       float64
	Direction   TradeDirection
	TotalAmount float64
	TotalFee    float64
}

func NewTransaction(stock *Stock, price, count float64, direction TradeDirection) *Transaction {
	totalAmount := price * count
	totalFee := stock.CalculateTotalFee(totalAmount, direction)
	return &Transaction{
		Stock:       stock,
		Price:       price,
		Count:       count,
		Direction:   direction,
		TotalAmount: totalAmount,
		TotalFee:    totalFee,
	}
}

func (t *Transaction) String() string {
	directionStr := "买入"
	if t.Direction == Sell {
		directionStr = "卖出"
	}

	return fmt.Sprintf("价格: %.2f | 数量: %.0f | 方向: %s | 金额: %.2f | 费用: %.2f",
		t.Price, t.Count, directionStr, t.TotalAmount, t.TotalFee)
}
