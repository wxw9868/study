package service

import (
	"errors"
	"time"

	"study/StockCalculator/ledger/database"
	"study/StockCalculator/ledger/model"
	"study/StockCalculator/ledger/request"

	"gorm.io/gorm"
)

var db = database.DB()

type LedgerService struct{}

func (s *LedgerService) CreateLedger(req *request.Ledger) error {
	// 检查货币是否存在
	var currency model.Currency
	if err := db.Where("id = ?", req.CurrencyId).First(&currency).Error; err != nil {
		return err
	}

	var ledger model.Ledger
	if err := db.Where("name = ?", req.Name).First(&ledger).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("账本已存在")
	}

	return db.Transaction(func(tx *gorm.DB) error {
		// 创建账本
		ledger := &model.Ledger{UserId: req.UserId, Name: req.Name, Amount: req.Amount, Date: time.Now(), CurrencyId: currency.ID}
		assetRecord := &model.AssetRecord{Amount: req.Amount, Date: time.Now()}
		if req.Date != "" {
			loc, _ := time.LoadLocation("Asia/Shanghai")
			t, _ := time.ParseInLocation("2006-01-02 15:04:05", req.Date, loc)
			ledger.Date = t
			assetRecord.Date = t
		}

		if err := db.Create(ledger).Error; err != nil {
			return err
		}

		assetRecord.LedgerId = ledger.ID
		if err := db.Create(assetRecord).Error; err != nil {
			return err
		}
		return nil
	})
}

// RecordAsset 在账本上记录资产
func (s *LedgerService) RecordAsset(req *request.AssetRecord) error {
	// 检查账本是否存在
	var ledger model.Ledger
	if err := db.Where("id = ?", req.LedgerId).First(&ledger).Error; err != nil {
		return err
	}

	// 创建资产记录
	assetRecord := &model.AssetRecord{
		LedgerId:     req.LedgerId,
		Amount:       req.Amount,
		Date:         time.Now(),
		Profit:       req.Amount - ledger.Amount,
		AnnualReturn: (req.Amount - ledger.Amount) / ledger.Amount,
	}
	// 如果请求中包含日期，则使用请求中的日期
	if req.Date != "" {
		loc, _ := time.LoadLocation("Asia/Shanghai")
		t, _ := time.ParseInLocation("2006-01-02 15:04:05", req.Date, loc)
		assetRecord.Date = t
	}

	return db.Create(assetRecord).Error
}

// GetLedgerAssets 获取账本下的所有资产记录
func (s *LedgerService) GetLedgerAssets(ledgerId uint) ([]model.AssetRecord, error) {
	var assets []model.AssetRecord
	err := db.Where("ledger_id = ?", ledgerId).Order("date desc").Find(&assets).Error
	return assets, err
}

// GetLedgerTotalAssets 获取账本的总资产
func (s *LedgerService) GetLedgerTotalAssets(userId uint) (any, error) {
	var ledgers []model.Ledger
	if err := db.Model(&model.Ledger{}).Where("user_id = ?", userId).Find(&ledgers).Error; err != nil {
		return nil, err
	}

	var ids []uint
	var initAmount float64
	for _, ledger := range ledgers {
		ids = append(ids, ledger.ID)
		initAmount += ledger.Amount
	}

	type result struct {
		LedgerId     uint    `gorm:"column:ledger_id" json:"ledger_id"`
		Name         string  `gorm:"column:name" json:"name"`
		Amount       float64 `gorm:"column:amount" json:"amount"`
		Profit       float64 `gorm:"column:profit" json:"profit"`
		AnnualReturn float64 `gorm:"column:annual_return" json:"annual_return"`
	}
	var res []result

	if err := db.Raw(`
	SELECT ledger_id, name, tmp1.amount, profit, annual_return
	FROM (
		SELECT ledger_id, amount, profit, annual_return
		FROM (
			SELECT 
				ledger_id, amount, profit, annual_return,
				ROW_NUMBER() OVER (PARTITION BY ledger_id ORDER BY date DESC) as rank_no
			FROM video_asset_record
			WHERE ledger_id in ?
		) AS tmp 
		WHERE rank_no = 1
	) AS tmp1
	LEFT JOIN video_ledger AS tmp2 ON tmp1.ledger_id = tmp2.id
	`, ids).Scan(&res).Error; err != nil {
		return nil, err
	}

	var totalAmount float64
	for _, v := range res {
		totalAmount += v.Amount
	}
	profit := totalAmount - initAmount
	annualReturn := profit / initAmount

	res = append(res, result{
		LedgerId:     0,
		Name:         "汇总总资产(元)",
		Amount:       totalAmount,
		Profit:       profit,
		AnnualReturn: annualReturn,
	})
	return res, nil
}

// GetLedgerById 获取账本信息
func (s *LedgerService) GetLedgerById(ledgerId uint, ledger *model.Ledger) error {
	// 从数据库获取账本信息
	return db.Where("id = ?", ledgerId).First(ledger).Error
}
