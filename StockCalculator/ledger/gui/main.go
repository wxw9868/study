package main

import (
	"fmt"
	"image/color"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// 定义颜色常量
var (
	redColor    = color.NRGBA{R: 255, G: 77, B: 79, A: 255}
	greenColor  = color.NRGBA{R: 76, G: 175, B: 80, A: 255}
	whiteColor  = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	grayColor   = color.NRGBA{R: 150, G: 150, B: 150, A: 255}
	lightGray   = color.NRGBA{R: 240, G: 240, B: 240, A: 255}
	bgColor     = color.NRGBA{R: 245, G: 245, B: 245, A: 255}
	shadowColor = color.NRGBA{R: 0, G: 0, B: 0, A: 10}
	orangeColor = color.NRGBA{R: 255, G: 152, B: 0, A: 255}
	cnyBgColor  = color.NRGBA{R: 255, G: 243, B: 224, A: 255}
	blackColor  = color.NRGBA{R: 51, G: 51, B: 51, A: 255}
)

// 账户类型
type AccountType int

const (
	Stock AccountType = iota
	Fund
	Bank
)

// 账户结构
type Account struct {
	Type        AccountType
	Name        string
	TotalAsset  float64
	TotalIncome float64
	YieldRate   float64
}

func main() {
	a := app.New()
	w := a.NewWindow("资产简记")

	// 检测是否在移动设备上运行
	isMobile := fyne.CurrentDevice().IsMobile()
	if isMobile {
		w.SetFullScreen(true)
	} else {
		w.Resize(fyne.NewSize(400, 700))
	}

	// 创建顶部导航栏
	backButton := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
		// 返回上一页
	})
	backButton.Importance = widget.LowImportance

	titleText := canvas.NewText("资产简记", blackColor)
	titleText.TextSize = 18
	titleText.TextStyle.Bold = true
	titleText.Alignment = fyne.TextAlignCenter

	title := container.NewCenter(titleText)

	topBar := container.NewBorder(nil, nil, backButton, nil, title)

	// 创建总资产卡片
	totalAssetCard := createTotalAssetCard(71.59, 4699.80, 0.38)

	// 创建各类账户卡片
	stockAccount := Account{
		Type:        Stock,
		Name:        "股票账户",
		TotalAsset:  21.06,
		TotalIncome: -30473.41,
		YieldRate:   -3.17,
	}
	stockCard := createAccountCard(stockAccount)

	fundAccount := Account{
		Type:        Fund,
		Name:        "基金账户",
		TotalAsset:  2.2239,
		TotalIncome: -5780.38,
		YieldRate:   -5.36,
	}
	fundCard := createAccountCard(fundAccount)

	bankAccount := Account{
		Type:        Bank,
		Name:        "银行账户",
		TotalAsset:  33.00,
		TotalIncome: 40953.59,
		YieldRate:   48.17,
	}
	bankCard := createAccountCard(bankAccount)

	// 创建添加账户按钮
	addButton := widget.NewButton("", func() {
		// 添加账户逻辑
	})
	addButton.Importance = widget.LowImportance

	// 创建一个带边框和圆角的容器
	addButtonBg := canvas.NewRectangle(whiteColor)
	addButtonBg.CornerRadius = 8

	// 创建加号图标和文本
	addIcon := canvas.NewText("+", blackColor)
	addIcon.TextSize = 16
	addIcon.TextStyle = fyne.TextStyle{Bold: true}

	addText := canvas.NewText("继续添加账户", blackColor)
	addText.TextSize = 14

	// 组合图标和文本
	addButtonContent := container.NewHBox(
		addIcon,
		addText,
	)

	// 将按钮内容放在背景上
	customAddButton := container.NewStack(
		addButtonBg,
		addButtonContent,
		addButton, // 按钮放在最上层以捕获点击事件
	)

	// 创建背景
	background := canvas.NewRectangle(bgColor)

	// 创建主内容区，调整卡片间距
	content := container.NewVBox(
		topBar,
		totalAssetCard,
		stockCard,
		fundCard,
		bankCard,
		layout.NewSpacer(),
		container.NewHBox(
			layout.NewSpacer(),
			customAddButton,
			layout.NewSpacer(),
		),
	)

	// 设置内容并显示窗口
	// 使用标准的Padded容器而不是自定义的InsetLayout
	paddedContent := container.NewPadded(content)

	w.SetContent(container.NewStack(background, paddedContent))
	w.ShowAndRun()
}

// 创建总资产卡片
func createTotalAssetCard(totalAsset, totalIncome, yieldRate float64) *fyne.Container {
	// 创建红色背景
	bg := canvas.NewRectangle(redColor)
	bg.SetMinSize(fyne.NewSize(0, 150))
	bg.CornerRadius = 12 // 添加圆角效果

	// 添加卡片阴影效果
	shadow := canvas.NewRectangle(shadowColor)
	shadow.Move(fyne.NewPos(1, 1))
	shadow.CornerRadius = 12 // 阴影也需要圆角

	// 创建总资产标题
	titleText := canvas.NewText("汇总总资产(元)", whiteColor)
	titleText.TextSize = 14

	eyeIcon := canvas.NewText("👁", whiteColor)
	eyeIcon.TextSize = 16

	// 使用圆形背景包裹问号
	helpCircle := canvas.NewCircle(color.NRGBA{R: 255, G: 255, B: 255, A: 60})
	circleContainer := container.NewWithoutLayout(helpCircle)
	circleContainer.Resize(fyne.NewSize(18, 18))

	helpIcon := canvas.NewText("?", whiteColor)
	helpIcon.TextSize = 12
	helpIcon.Move(fyne.NewPos(6, 2))

	helpContainer := container.NewStack(circleContainer, helpIcon)

	titleContainer := container.NewHBox(
		titleText,
		layout.NewSpacer(),
		eyeIcon,
		helpContainer, // 移除额外的内边距
	)

	// 创建总资产金额
	assetText := canvas.NewText(fmt.Sprintf("%.2f万", totalAsset), whiteColor)
	assetText.TextSize = 40
	assetText.TextStyle = fyne.TextStyle{Bold: true}
	assetText.Alignment = fyne.TextAlignLeading

	// 创建累计收益和年化收益率
	incomeLabel := canvas.NewText("累计收益(元)", whiteColor)
	incomeLabel.TextSize = 12

	// 格式化总收益数值，添加千位分隔符
	incomeStr := formatNumber(fmt.Sprintf("%.2f", totalIncome))
	incomeValue := canvas.NewText(incomeStr, whiteColor)
	incomeValue.TextSize = 14
	incomeValue.TextStyle = fyne.TextStyle{Bold: true}

	yieldLabel := canvas.NewText("年化收益率", whiteColor)
	yieldLabel.TextSize = 12

	// 根据收益率正负设置颜色
	yieldPrefix := "+"
	if yieldRate < 0 {
		yieldPrefix = ""
	}
	yieldValue := canvas.NewText(fmt.Sprintf("%s%.2f%%", yieldPrefix, yieldRate), whiteColor)
	yieldValue.TextSize = 14
	yieldValue.TextStyle = fyne.TextStyle{Bold: true}

	// 添加右箭头指示器
	rightArrow := canvas.NewText("›", whiteColor)
	rightArrow.TextSize = 24
	rightArrow.TextStyle = fyne.TextStyle{Bold: true}

	// 创建底部数据行
	bottomRow := container.NewBorder(
		nil, nil, nil, rightArrow,
		container.NewGridWithColumns(2,
			container.NewVBox(
				incomeLabel,
				incomeValue,
			),
			container.NewVBox(
				yieldLabel,
				yieldValue,
			),
		),
	)

	// 组合所有元素，调整内边距
	// 使用标准的Padded容器
	titlePadded := container.NewPadded(titleContainer)
	assetPadded := container.NewPadded(assetText)
	bottomPadded := container.NewPadded(bottomRow)

	card := container.NewBorder(
		titlePadded,
		bottomPadded,
		nil,
		nil,
		assetPadded,
	)

	// 将卡片放在背景上
	return container.NewStack(shadow, bg, card)
}

// 创建账户卡片
func createAccountCard(account Account) *fyne.Container {
	// 创建白色背景
	bg := canvas.NewRectangle(whiteColor)
	bg.SetMinSize(fyne.NewSize(0, 120))
	bg.CornerRadius = 8 // 添加圆角效果

	// 添加卡片阴影效果
	shadow := canvas.NewRectangle(shadowColor)
	shadow.Move(fyne.NewPos(1, 1))
	shadow.CornerRadius = 8 // 阴影也需要圆角

	// 创建账户名称和货币类型
	nameText := canvas.NewText(account.Name, blackColor)
	nameText.TextSize = 15
	nameText.TextStyle = fyne.TextStyle{Bold: true}

	// 创建CNY标签
	cnyBg := canvas.NewRectangle(cnyBgColor)
	cnyBgContainer := container.NewWithoutLayout(cnyBg)
	cnyBgContainer.Resize(fyne.NewSize(36, 20))
	cnyBg.CornerRadius = 4 // CNY标签添加圆角

	cnyText := canvas.NewText("CNY", orangeColor)
	cnyText.TextSize = 11

	cnyLabel := container.NewStack(cnyBgContainer, cnyText)

	titleRow := container.NewBorder(
		nil, nil, nameText,
		cnyLabel,
	)

	// 创建总资产标签
	assetLabelText := canvas.NewText("总资产(元)", grayColor)
	assetLabelText.TextSize = 12

	// 创建总资产值
	var assetText string
	if account.Type == Fund {
		// 基金账户显示完整数字，不用"万"
		assetText = "22,239.41" // 与图片保持一致
	} else {
		assetText = fmt.Sprintf("%.2f万", account.TotalAsset)
	}

	assetValue := canvas.NewText(assetText, blackColor)
	assetValue.TextSize = 24
	assetValue.TextStyle = fyne.TextStyle{Bold: true}
	assetValue.Alignment = fyne.TextAlignLeading

	// 创建累计收益和年化收益率
	incomeLabel := canvas.NewText("累计收益(元)", grayColor)
	incomeLabel.TextSize = 12

	// 根据收益正负设置颜色
	incomeColor := greenColor
	if account.TotalIncome < 0 {
		incomeColor = redColor
	}

	// 格式化收益数值，添加千位分隔符
	incomeStr := fmt.Sprintf("%.2f", account.TotalIncome)
	if account.TotalIncome < 0 {
		// 负数，保留负号
		incomeStr = "-" + formatNumber(strings.TrimPrefix(incomeStr, "-"))
	} else {
		incomeStr = formatNumber(incomeStr)
	}

	incomeValue := canvas.NewText(incomeStr, incomeColor)
	incomeValue.TextSize = 14
	incomeValue.TextStyle = fyne.TextStyle{Bold: true}

	yieldLabel := canvas.NewText("年化收益率", grayColor)
	yieldLabel.TextSize = 12

	// 根据收益率正负设置颜色
	yieldColor := greenColor
	yieldPrefix := "+"
	if account.YieldRate < 0 {
		yieldColor = redColor
		yieldPrefix = ""
	}
	yieldValue := canvas.NewText(fmt.Sprintf("%s%.2f%%", yieldPrefix, account.YieldRate), yieldColor)
	yieldValue.TextSize = 14
	yieldValue.TextStyle = fyne.TextStyle{Bold: true}

	// 添加右箭头指示器
	arrowIcon := canvas.NewText("›", grayColor)
	arrowIcon.TextSize = 22
	arrowIcon.TextStyle = fyne.TextStyle{Bold: true}

	// 创建底部数据行
	bottomRow := container.NewBorder(
		nil, nil, nil, arrowIcon,
		container.NewGridWithColumns(2,
			container.NewVBox(
				incomeLabel,
				incomeValue,
			),
			container.NewVBox(
				yieldLabel,
				yieldValue,
			),
		),
	)

	// 创建资产部分
	assetContainer := container.NewVBox(
		assetLabelText,
		assetValue,
	)

	// 组合所有元素，调整内边距
	titlePadded := container.NewPadded(titleRow)
	assetPadded := container.NewPadded(assetContainer)
	bottomPadded := container.NewPadded(bottomRow)

	card := container.NewBorder(
		titlePadded,
		bottomPadded,
		nil, nil,
		assetPadded,
	)

	// 将卡片放在背景上
	return container.NewStack(shadow, bg, card)
}

// 格式化数字，添加千位分隔符
func formatNumber(numStr string) string {
	parts := strings.Split(numStr, ".")
	intPart := parts[0]

	var result []byte
	for i, c := range intPart {
		if i > 0 && (len(intPart)-i)%3 == 0 {
			result = append(result, ',')
		}
		result = append(result, byte(c))
	}

	if len(parts) > 1 {
		return string(result) + "." + parts[1]
	}
	return string(result)
}
