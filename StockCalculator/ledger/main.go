package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/xuri/excelize/v2"
)

// 资产类型
type AssetType string

const (
	Stock    AssetType = "股票"
	Fund     AssetType = "基金"
	Bond     AssetType = "债券"
	Cash     AssetType = "现金"
	Property AssetType = "房产"
	Other    AssetType = "其他"
)

// 资产记录
type AssetRecord struct {
	ID            int       // 记录ID
	Type          AssetType // 资产类型
	Name          string    // 资产名称
	Amount        float64   // 资产金额
	InitialAmount float64   // 初始投资金额
	Date          time.Time // 记录日期
	Profit        float64   // 累计收益
	AnnualReturn  float64   // 年化收益率(%)
	Notes         string    // 备注
}

// 资产管理器
type AssetManager struct {
	records      []AssetRecord
	nextID       int
	dataFilePath string
}

// 创建新的资产管理器
func NewAssetManager() *AssetManager {
	// 获取用户主目录
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Printf("无法获取用户主目录: %v", err)
		homeDir = "."
	}

	// 创建数据目录
	dataDir := filepath.Join(homeDir, ".stockcalculator")
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		log.Printf("无法创建数据目录: %v", err)
	}

	return &AssetManager{
		records:      []AssetRecord{},
		nextID:       1,
		dataFilePath: filepath.Join(dataDir, "assets.xlsx"),
	}
}

// 添加资产记录
func (am *AssetManager) AddRecord(record AssetRecord) {
	record.ID = am.nextID
	am.nextID++

	// 计算累计收益
	record.Profit = record.Amount - record.InitialAmount

	// 计算年化收益率
	daysHeld := float64(time.Now().Sub(record.Date).Hours() / 24)
	if daysHeld > 0 && record.InitialAmount > 0 {
		// 年化收益率 = (当前金额/初始金额)^(365/持有天数) - 1
		annualReturn := (math.Pow(record.Amount/record.InitialAmount, 365.0/daysHeld) - 1) * 100
		// 保留两位小数
		record.AnnualReturn = math.Round(annualReturn*100) / 100
	} else {
		// 如果是当天添加的资产，设置年化收益率为0
		record.AnnualReturn = 0
	}

	am.records = append(am.records, record)
}

// 更新资产记录
func (am *AssetManager) UpdateRecord(id int, record AssetRecord) bool {
	for i, r := range am.records {
		if r.ID == id {
			record.ID = id

			// 计算累计收益
			record.Profit = record.Amount - record.InitialAmount

			// 计算年化收益率
			daysHeld := float64(time.Now().Sub(record.Date).Hours() / 24)
			if daysHeld > 0 && record.InitialAmount > 0 {
				// 年化收益率 = (当前金额/初始金额)^(365/持有天数) - 1
				annualReturn := (math.Pow(record.Amount/record.InitialAmount, 365.0/daysHeld) - 1) * 100
				// 保留两位小数
				record.AnnualReturn = math.Round(annualReturn*100) / 100
			} else {
				// 如果是当天添加的资产，设置年化收益率为0
				record.AnnualReturn = 0
			}

			am.records[i] = record
			return true
		}
	}
	return false
}

// 删除资产记录
func (am *AssetManager) DeleteRecord(id int) bool {
	for i, r := range am.records {
		if r.ID == id {
			am.records = append(am.records[:i], am.records[i+1:]...)
			return true
		}
	}
	return false
}

// 获取所有资产记录
func (am *AssetManager) GetAllRecords() []AssetRecord {
	return am.records
}

// 获取特定类型的资产记录
func (am *AssetManager) GetRecordsByType(assetType AssetType) []AssetRecord {
	var result []AssetRecord
	for _, r := range am.records {
		if r.Type == assetType {
			result = append(result, r)
		}
	}
	return result
}

// 获取特定ID的资产记录
func (am *AssetManager) GetRecordByID(id int) (AssetRecord, bool) {
	for _, r := range am.records {
		if r.ID == id {
			return r, true
		}
	}
	return AssetRecord{}, false
}

// 计算总资产
func (am *AssetManager) CalculateTotalAssets() float64 {
	total := 0.0
	for _, r := range am.records {
		total += r.Amount
	}
	return total
}

// 计算总收益
func (am *AssetManager) CalculateTotalProfit() float64 {
	total := 0.0
	for _, r := range am.records {
		total += r.Profit
	}
	return total
}

// 保存数据到Excel文件
func (am *AssetManager) SaveToFile() error {
	f := excelize.NewFile()

	// 创建资产记录表
	sheetName := "资产记录"
	f.NewSheet(sheetName)

	// 设置表头
	headers := []string{"ID", "资产类型", "资产名称", "当前金额", "初始金额", "记录日期", "累计收益", "年化收益率(%)", "备注"}
	for i, header := range headers {
		cell := fmt.Sprintf("%c1", 'A'+i)
		f.SetCellValue(sheetName, cell, header)
	}

	// 写入数据
	for i, record := range am.records {
		row := i + 2 // 从第2行开始写入数据
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), record.ID)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), record.Type)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), record.Name)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), record.Amount)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", row), record.InitialAmount)
		f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), record.Date.Format("2006-01-02"))
		f.SetCellValue(sheetName, fmt.Sprintf("G%d", row), record.Profit)
		// 确保Excel中的年化收益率也保持两位小数
		f.SetCellValue(sheetName, fmt.Sprintf("H%d", row), math.Round(record.AnnualReturn*100)/100)
		f.SetCellValue(sheetName, fmt.Sprintf("I%d", row), record.Notes)
	}

	// 删除默认的Sheet1
	f.DeleteSheet("Sheet1")

	// 保存文件
	if err := f.SaveAs(am.dataFilePath); err != nil {
		return err
	}

	return nil
}

// 从Excel文件加载数据
func (am *AssetManager) LoadFromFile() error {
	// 检查文件是否存在
	if _, err := os.Stat(am.dataFilePath); os.IsNotExist(err) {
		// 文件不存在，不需要加载
		return nil
	}

	// 打开Excel文件
	f, err := excelize.OpenFile(am.dataFilePath)
	if err != nil {
		return err
	}
	defer f.Close()

	// 获取所有行
	sheetName := "资产记录"
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return err
	}

	// 清空现有记录
	am.records = []AssetRecord{}
	maxID := 0

	// 从第2行开始读取数据（跳过表头）
	for i := 1; i < len(rows); i++ {
		row := rows[i]
		if len(row) < 9 {
			continue // 跳过不完整的行
		}

		// 解析ID
		id, err := strconv.Atoi(row[0])
		if err != nil {
			continue
		}

		// 更新最大ID
		if id > maxID {
			maxID = id
		}

		// 解析金额
		amount, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			continue
		}

		initialAmount, err := strconv.ParseFloat(row[4], 64)
		if err != nil {
			continue
		}

		// 解析日期
		date, err := time.Parse("2006-01-02", row[5])
		if err != nil {
			continue
		}

		// 解析收益
		profit, err := strconv.ParseFloat(row[6], 64)
		if err != nil {
			continue
		}

		// 解析年化收益率
		annualReturn, err := strconv.ParseFloat(row[7], 64)
		if err != nil {
			continue
		}

		// 创建记录
		record := AssetRecord{
			ID:            id,
			Type:          AssetType(row[1]),
			Name:          row[2],
			Amount:        amount,
			InitialAmount: initialAmount,
			Date:          date,
			Profit:        profit,
			AnnualReturn:  annualReturn,
			Notes:         row[8],
		}

		am.records = append(am.records, record)
	}

	// 设置下一个ID
	am.nextID = maxID + 1

	return nil
}

func main() {
	// 创建资产管理器
	assetManager := NewAssetManager()

	// 加载数据
	if err := assetManager.LoadFromFile(); err != nil {
		log.Printf("加载数据失败: %v", err)
	}

	// 创建应用
	a := app.New()
	w := a.NewWindow("个人资产管理")
	w.Resize(fyne.NewSize(1000, 600))

	// 创建资产列表
	assetList := binding.NewUntypedList()

	// 创建统计信息面板
	totalAssetsLabel := widget.NewLabel("")
	totalProfitLabel := widget.NewLabel("")

	// 更新统计信息
	updateStats := func() {
		totalAssets := assetManager.CalculateTotalAssets()
		totalProfit := assetManager.CalculateTotalProfit()

		totalAssetsLabel.SetText(fmt.Sprintf("总资产: %.2f 元", totalAssets))

		if totalProfit > 0 {
			totalProfitLabel.SetText(fmt.Sprintf("总收益: +%.2f 元", totalProfit))
			totalProfitLabel.TextStyle = fyne.TextStyle{Bold: true}
		} else if totalProfit < 0 {
			totalProfitLabel.SetText(fmt.Sprintf("总收益: %.2f 元", totalProfit))
			totalProfitLabel.TextStyle = fyne.TextStyle{Bold: true}
		} else {
			totalProfitLabel.SetText("总收益: 0.00 元")
			totalProfitLabel.TextStyle = fyne.TextStyle{}
		}
		totalProfitLabel.Refresh()
	}

	updateAssetList := func() {
		// 将[]AssetRecord转换为[]any
		records := assetManager.GetAllRecords()
		items := make([]interface{}, len(records))
		for i, r := range records {
			items[i] = r
		}
		assetList.Set(items)

		// 更新完列表后同时更新统计信息
		updateStats()
	}
	updateAssetList()

	// 创建资产类型选择器
	assetTypes := []string{
		string(Stock),
		string(Fund),
		string(Bond),
		string(Cash),
		string(Property),
		string(Other),
	}

	// 创建表格
	table := widget.NewTable(
		func() (int, int) {
			items, _ := assetList.Get()
			return len(items) + 1, 9 // +1 for header row, 9 columns
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Wide Content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			label := o.(*widget.Label)
			label.TextStyle = fyne.TextStyle{}
			label.Alignment = fyne.TextAlignLeading

			if i.Row == 0 {
				// 表头
				headers := []string{"ID", "类型", "名称", "当前金额", "初始金额", "日期", "收益", "年化率", "备注"}
				label.SetText(headers[i.Col])
				label.TextStyle.Bold = true
				// 设置表头背景色
				label.Alignment = fyne.TextAlignCenter
			} else {
				// 数据行
				items, _ := assetList.Get()
				if i.Row-1 < len(items) {
					record := items[i.Row-1].(AssetRecord)

					switch i.Col {
					case 0:
						label.SetText(fmt.Sprintf("%d", record.ID))
						label.Alignment = fyne.TextAlignCenter
					case 1:
						label.SetText(string(record.Type))
						label.Alignment = fyne.TextAlignCenter
					case 2:
						label.SetText(record.Name)
					case 3:
						label.SetText(fmt.Sprintf("%.2f", record.Amount))
						label.Alignment = fyne.TextAlignTrailing
					case 4:
						label.SetText(fmt.Sprintf("%.2f", record.InitialAmount))
						label.Alignment = fyne.TextAlignTrailing
					case 5:
						label.SetText(record.Date.Format("2006-01-02"))
						label.Alignment = fyne.TextAlignCenter
					case 6:
						// 收益显示优化
						if record.Profit > 0 {
							label.SetText(fmt.Sprintf("+%.2f", record.Profit))
							label.TextStyle = fyne.TextStyle{Bold: true}
							// 正收益显示为绿色
							// 注意：Fyne不直接支持文本颜色，这里只是加粗处理
						} else if record.Profit < 0 {
							label.SetText(fmt.Sprintf("%.2f", record.Profit))
						} else {
							label.SetText("0.00")
						}
						label.Alignment = fyne.TextAlignTrailing
					case 7:
						// 优化年化收益率显示
						if record.AnnualReturn > 0 {
							label.SetText(fmt.Sprintf("+%.2f%%", record.AnnualReturn))
							label.TextStyle = fyne.TextStyle{Bold: true}
						} else if record.AnnualReturn < 0 {
							label.SetText(fmt.Sprintf("%.2f%%", record.AnnualReturn))
						} else {
							label.SetText("0.00%")
						}
						label.Alignment = fyne.TextAlignTrailing
					case 8:
						label.SetText(record.Notes)
					}
				}
			}
		})

	// 设置列宽
	table.SetColumnWidth(0, 40)  // ID
	table.SetColumnWidth(1, 60)  // 类型
	table.SetColumnWidth(2, 150) // 名称
	table.SetColumnWidth(3, 90)  // 当前金额
	table.SetColumnWidth(4, 90)  // 初始金额
	table.SetColumnWidth(5, 90)  // 日期
	table.SetColumnWidth(6, 90)  // 累计收益
	table.SetColumnWidth(7, 80)  // 年化收益率
	table.SetColumnWidth(8, 150) // 备注

	// 创建添加资产对话框
	showAddAssetDialog := func() {
		// 创建表单项
		typeSelect := widget.NewSelect(assetTypes, nil)
		nameEntry := widget.NewEntry()
		amountEntry := widget.NewEntry()
		initialAmountEntry := widget.NewEntry()
		dateEntry := widget.NewEntry()
		dateEntry.SetText(time.Now().Format("2006-01-02"))
		notesEntry := widget.NewMultiLineEntry()

		// 设置默认值
		typeSelect.SetSelected(string(Stock))

		// 创建表单
		form := &widget.Form{
			Items: []*widget.FormItem{
				{Text: "资产类型", Widget: typeSelect},
				{Text: "资产名称", Widget: nameEntry},
				{Text: "当前金额", Widget: amountEntry},
				{Text: "初始金额", Widget: initialAmountEntry},
				{Text: "购入日期", Widget: dateEntry},
				{Text: "备注", Widget: notesEntry},
			},
			OnSubmit: func() {
				// 解析输入
				assetType := AssetType(typeSelect.Selected)
				name := nameEntry.Text

				amount, err := strconv.ParseFloat(amountEntry.Text, 64)
				if err != nil {
					dialog.ShowError(fmt.Errorf("当前金额必须是数字"), w)
					return
				}

				initialAmount, err := strconv.ParseFloat(initialAmountEntry.Text, 64)
				if err != nil {
					dialog.ShowError(fmt.Errorf("初始金额必须是数字"), w)
					return
				}

				date, err := time.Parse("2006-01-02", dateEntry.Text)
				if err != nil {
					dialog.ShowError(fmt.Errorf("日期格式无效，请使用YYYY-MM-DD格式"), w)
					return
				}

				// 创建资产记录
				record := AssetRecord{
					Type:          assetType,
					Name:          name,
					Amount:        amount,
					InitialAmount: initialAmount,
					Date:          date,
					Notes:         notesEntry.Text,
				}

				// 添加记录
				assetManager.AddRecord(record)

				// 保存数据
				if err := assetManager.SaveToFile(); err != nil {
					dialog.ShowError(fmt.Errorf("保存数据失败: %v", err), w)
				}

				// 更新列表（已包含更新统计信息）
				updateAssetList()

				// 显示成功消息
				dialog.ShowInformation("添加成功", "资产记录已成功添加", w)
			},
		}

		// 显示对话框
		dialog.ShowCustom("添加资产", "取消", form, w)
	}

	// 创建编辑资产对话框
	showEditAssetDialog := func(id int) {
		// 获取记录
		record, found := assetManager.GetRecordByID(id)
		if !found {
			dialog.ShowError(fmt.Errorf("找不到ID为%d的记录", id), w)
			return
		}

		// 创建表单项
		typeSelect := widget.NewSelect(assetTypes, nil)
		nameEntry := widget.NewEntry()
		amountEntry := widget.NewEntry()
		initialAmountEntry := widget.NewEntry()
		dateEntry := widget.NewEntry()
		notesEntry := widget.NewMultiLineEntry()

		// 设置当前值
		typeSelect.SetSelected(string(record.Type))
		nameEntry.SetText(record.Name)
		amountEntry.SetText(fmt.Sprintf("%.2f", record.Amount))
		initialAmountEntry.SetText(fmt.Sprintf("%.2f", record.InitialAmount))
		dateEntry.SetText(record.Date.Format("2006-01-02"))
		notesEntry.SetText(record.Notes)

		// 创建表单
		form := &widget.Form{
			Items: []*widget.FormItem{
				{Text: "资产类型", Widget: typeSelect},
				{Text: "资产名称", Widget: nameEntry},
				{Text: "当前金额", Widget: amountEntry},
				{Text: "初始金额", Widget: initialAmountEntry},
				{Text: "购入日期", Widget: dateEntry},
				{Text: "备注", Widget: notesEntry},
			},
			OnSubmit: func() {
				// 解析输入
				assetType := AssetType(typeSelect.Selected)
				name := nameEntry.Text

				amount, err := strconv.ParseFloat(amountEntry.Text, 64)
				if err != nil {
					dialog.ShowError(fmt.Errorf("当前金额必须是数字"), w)
					return
				}

				initialAmount, err := strconv.ParseFloat(initialAmountEntry.Text, 64)
				if err != nil {
					dialog.ShowError(fmt.Errorf("初始金额必须是数字"), w)
					return
				}

				date, err := time.Parse("2006-01-02", dateEntry.Text)
				if err != nil {
					dialog.ShowError(fmt.Errorf("日期格式无效，请使用YYYY-MM-DD格式"), w)
					return
				}

				// 更新记录
				updatedRecord := AssetRecord{
					Type:          assetType,
					Name:          name,
					Amount:        amount,
					InitialAmount: initialAmount,
					Date:          date,
					Notes:         notesEntry.Text,
				}

				if assetManager.UpdateRecord(id, updatedRecord) {
					// 保存数据
					if err := assetManager.SaveToFile(); err != nil {
						dialog.ShowError(fmt.Errorf("保存数据失败: %v", err), w)
					}

					// 更新列表（已包含更新统计信息）
					updateAssetList()

					// 显示成功消息
					dialog.ShowInformation("更新成功", "资产记录已成功更新", w)
				} else {
					dialog.ShowError(fmt.Errorf("更新记录失败"), w)
				}
			},
		}

		// 显示对话框
		dialog.ShowCustom("编辑资产", "取消", form, w)
	}

	// 创建删除资产对话框
	showDeleteConfirmDialog := func(id int) {
		// 获取记录
		record, found := assetManager.GetRecordByID(id)
		if !found {
			dialog.ShowError(fmt.Errorf("找不到ID为%d的记录", id), w)
			return
		}

		// 显示确认对话框
		dialog.ShowConfirm(
			"确认删除",
			fmt.Sprintf("确定要删除 %s (ID: %d) 吗？此操作不可撤销。", record.Name, record.ID),
			func(confirmed bool) {
				if confirmed {
					if assetManager.DeleteRecord(id) {
						// 保存数据
						if err := assetManager.SaveToFile(); err != nil {
							dialog.ShowError(fmt.Errorf("保存数据失败: %v", err), w)
						}

						// 更新列表和统计信息
						updateAssetList()

						// 显示成功消息
						dialog.ShowInformation("删除成功", "资产记录已成功删除", w)
					} else {
						dialog.ShowError(fmt.Errorf("删除记录失败"), w)
					}
				}
			},
			w,
		)
	}

	// 创建表格上下文菜单
	table.OnSelected = func(id widget.TableCellID) {
		// 只处理数据行的选择，跳过表头行
		if id.Row > 0 {
			items, _ := assetList.Get()
			if id.Row-1 < len(items) {
				record := items[id.Row-1].(AssetRecord)

				// 创建上下文菜单
				menu := fyne.NewMenu("资产操作",
					fyne.NewMenuItem("编辑", func() {
						showEditAssetDialog(record.ID)
					}),
					fyne.NewMenuItem("删除", func() {
						showDeleteConfirmDialog(record.ID)
					}),
				)

				// 修复：使用正确的API显示上下文菜单
				widget.ShowPopUpMenuAtPosition(menu, fyne.CurrentApp().Driver().CanvasForObject(table),
					fyne.CurrentApp().Driver().AbsolutePositionForObject(table))
			}
		}
	}

	// 删除重复的声明，使用已经定义的变量
	// 初始更新统计信息
	updateStats()

	// 创建按钮
	addButton := widget.NewButtonWithIcon("添加资产", theme.ContentAddIcon(), showAddAssetDialog)
	refreshButton := widget.NewButtonWithIcon("刷新", theme.ViewRefreshIcon(), func() {
		updateAssetList()
		updateStats()
	})

	// 创建过滤器
	filterSelect := widget.NewSelect(append([]string{"全部"}, assetTypes...), func(selected string) {
		// 将[]AssetRecord转换为[]any
		var records []AssetRecord
		if selected == "全部" {
			records = assetManager.GetAllRecords()
		} else {
			records = assetManager.GetRecordsByType(AssetType(selected))
		}

		items := make([]interface{}, len(records))
		for i, r := range records {
			items[i] = r
		}
		assetList.Set(items)

		// 过滤后更新表格
		table.Refresh()

		// 单独更新统计信息，因为过滤后的统计应该基于当前显示的记录
		totalAssets := 0.0
		totalProfit := 0.0
		for _, r := range records {
			totalAssets += r.Amount
			totalProfit += r.Profit
		}

		totalAssetsLabel.SetText(fmt.Sprintf("总资产: %.2f 元", totalAssets))
		totalProfitLabel.SetText(fmt.Sprintf("总收益: %.2f 元", totalProfit))

		if totalProfit > 0 {
			totalProfitLabel.TextStyle = fyne.TextStyle{Bold: true}
		} else {
			totalProfitLabel.TextStyle = fyne.TextStyle{}
		}
		totalProfitLabel.Refresh()
	})
	filterSelect.SetSelected("全部")

	// 创建布局
	toolbar := container.NewHBox(
		addButton,
		refreshButton,
		layout.NewSpacer(),
		widget.NewLabel("过滤:"),
		filterSelect,
	)

	statsBar := container.NewHBox(
		totalAssetsLabel,
		layout.NewSpacer(),
		totalProfitLabel,
	)

	// 创建主布局
	content := container.NewBorder(
		container.NewVBox(
			widget.NewLabel("个人资产管理"),
			toolbar,
		),
		statsBar,
		nil,
		nil,
		container.NewPadded(table),
	)

	// 设置窗口内容
	w.SetContent(content)

	// 设置窗口关闭回调
	w.SetOnClosed(func() {
		// 保存数据
		if err := assetManager.SaveToFile(); err != nil {
			log.Printf("保存数据失败: %v", err)
		}
	})

	// 显示窗口
	w.ShowAndRun()
}
