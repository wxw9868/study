package main

import (
	"fmt"
	"study/StockCalculator/ledger/request"
	"study/StockCalculator/ledger/service"
)

func main() {
	s := new(service.LedgerService)

	if err := s.CreateLedger(&request.Ledger{}); err != nil {
		panic(err)
	}

	if err := s.RecordAsset(&request.AssetRecord{}); err != nil {
		panic(err)
	}

	data, err := s.GetLedgerAssets(0)
	if err != nil {
		panic(err)
	}
	fmt.Printf("data: %+v\n", data)

	res, err := s.GetLedgerTotalAssets(0)
	if err != nil {
		panic(err)
	}
	fmt.Printf("data: %+v\n", res)
}
