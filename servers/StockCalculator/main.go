package main

import (
	"flag"
	"fmt"
	"study/StockCalculator/server"
)

// go run stock_calculator.go -rate 0.0185 -buy 76.10 -sell 76.68 -count 200
// go run stock_calculator.go -rate 0.0185 -buy 3.69 -sell 3.74 -count 4500
// go run stock_calculator.go -rate 0.0185 -buy 1.84 -sell 1.88 -count 9300
func main() {
	fmt.Println("=== 股票计算器 ===")
	// 处理应用事件
	customRate := flag.Float64("rate", 0.03, "佣金费率")
	buyPrice := flag.Float64("buy", 1.0, "买入价格")
	sellPrice := flag.Float64("sell", 1.1, "卖出价格")
	stockCount := flag.Int("count", 10000, "股票数量")
	flag.Parse()

	// 参数验证
	if *buyPrice <= 0 || *sellPrice <= 0 || *stockCount <= 0 {
		fmt.Println("错误：所有参数都必须是正值！")
		return
	}

	stock := server.NewStock(*customRate / 100)
	buyTransaction := server.NewTransaction(stock, *buyPrice, float64(*stockCount), server.Buy)
	sellTransaction := server.NewTransaction(stock, *sellPrice, float64(*stockCount), server.Sell)

	// 新增收益率计算方法
	calculateYield := func(buy, sell *server.Transaction) float64 {
		return (sell.Price - buy.Price) / buy.Price
	}

	// 优化输出格式
	fmt.Println("\n=== 交易详情 ===")
	fmt.Printf("【买入】%s\n", buyTransaction.String())
	fmt.Printf("【卖出】%s\n\n", sellTransaction.String())

	fmt.Printf("手续费: %.2f CNY\n", buyTransaction.TotalFee+sellTransaction.TotalFee)
	fmt.Printf("收益率: %.2f%%\n", calculateYield(buyTransaction, sellTransaction)*100)
	fmt.Printf("净利润: %.2f CNY\n", sellTransaction.TotalAmount-buyTransaction.TotalAmount-(buyTransaction.TotalFee+sellTransaction.TotalFee))
}
