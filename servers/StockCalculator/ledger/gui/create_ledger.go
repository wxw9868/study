package main

import (
	"fmt"
	"image/color"
	"time"

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
	redColor       = color.NRGBA{R: 255, G: 77, B: 79, A: 255}
	lightRedColor  = color.NRGBA{R: 255, G: 102, B: 102, A: 255}
	pinkColor      = color.NRGBA{R: 255, G: 182, B: 182, A: 255}
	whiteColor     = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	grayColor      = color.NRGBA{R: 150, G: 150, B: 150, A: 255}
	lightGrayColor = color.NRGBA{R: 240, G: 240, B: 240, A: 255}
	blackColor     = color.NRGBA{R: 51, G: 51, B: 51, A: 255}
	iconBgColor    = color.NRGBA{R: 255, G: 240, B: 240, A: 255}
)

func createAddAccountPage() fyne.Window {
	a := app.New()
	w := a.NewWindow("添加账户")

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

	titleText := canvas.NewText("添加账户", whiteColor)
	titleText.TextSize = 18
	titleText.TextStyle = fyne.TextStyle{Bold: true}
	titleText.Alignment = fyne.TextAlignCenter

	title := container.NewCenter(titleText)

	// 创建红色背景
	topBg := canvas.NewRectangle(redColor)
	topBg.SetMinSize(fyne.NewSize(0, 60))

	// 组合顶部导航栏
	topBar := container.NewStack(
		topBg,
		container.NewBorder(nil, nil, backButton, nil, title),
	)

	// 创建账本图标 - 优化为更接近图片的效果
	iconBg := canvas.NewCircle(iconBgColor)
	iconBgContainer := container.NewWithoutLayout(iconBg)
	iconBgContainer.Resize(fyne.NewSize(100, 100))

	// 创建账本图标 - 使用更复杂的组合来模拟图片中的账本图标
	ledgerBase := canvas.NewRectangle(redColor)
	ledgerBase.SetMinSize(fyne.NewSize(40, 40))

	// 添加白色小圆点模拟图标中的细节
	dot1 := canvas.NewCircle(whiteColor)
	dot1Container := container.NewWithoutLayout(dot1)
	dot1Container.Resize(fyne.NewSize(6, 6))
	dot1Container.Move(fyne.NewPos(10, 15))

	dot2 := canvas.NewCircle(whiteColor)
	dot2Container := container.NewWithoutLayout(dot2)
	dot2Container.Resize(fyne.NewSize(6, 6))
	dot2Container.Move(fyne.NewPos(10, 25))

	// 创建一个容器来组合账本图标的各个部分
	ledgerIconContainer := container.NewWithoutLayout(
		ledgerBase,
		dot1Container,
		dot2Container,
	)
	ledgerIconContainer.Resize(fyne.NewSize(40, 40))

	// 创建一个容器来居中显示图标
	iconContainer := container.NewCenter(container.NewStack(
		iconBgContainer,
		container.NewCenter(ledgerIconContainer),
	))

	// 添加一些垂直空间，使图标位置更接近图片
	// 使用空白矩形替代Spacer
	iconSpacer := canvas.NewRectangle(color.Transparent)
	iconSpacer.SetMinSize(fyne.NewSize(0, 20))

	// 创建输入字段 - 使用真实的输入框替代静态文本
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("请输入账户名称")
	// 设置样式
	nameEntry.TextStyle = fyne.TextStyle{Italic: true}

	pencilIcon := canvas.NewText("✎", grayColor)
	pencilIcon.TextSize = 16

	// 创建自定义输入行
	nameRow := container.NewBorder(
		nil, nil,
		container.NewHBox(pencilIcon),
		nil,
		nameEntry,
	)

	// 创建账户资产行 - 使用真实的输入框
	assetLabel := canvas.NewText("账户资产", blackColor)
	assetLabel.TextSize = 16
	assetLabel.TextStyle = fyne.TextStyle{Bold: true}

	// 创建金额输入框
	assetEntry := widget.NewEntry()
	assetEntry.SetPlaceHolder("请输入金额")
	assetEntry.TextStyle = fyne.TextStyle{Italic: true}
	// 设置右对齐
	assetEntry.Alignment = fyne.TextAlignTrailing

	pencilIcon2 := canvas.NewText("✎", grayColor)
	pencilIcon2.TextSize = 16

	// 组合资产输入行
	assetRow := container.NewBorder(
		nil, nil,
		assetLabel,
		container.NewHBox(pencilIcon2),
		assetEntry,
	)

	// 创建开始时间行 - 使用可点击的按钮
	startTimeLabel := canvas.NewText("开始时间", blackColor)
	startTimeLabel.TextSize = 16
	startTimeLabel.TextStyle = fyne.TextStyle{Bold: true}

	// 获取当前日期
	now := time.Now()
	dateStr := fmt.Sprintf("%d-%02d-%02d (今天)", now.Year(), now.Month(), now.Day())

	// 创建日期选择按钮
	dateSelectButton := widget.NewButton(dateStr, func() {
		// 日期选择逻辑
	})
	dateSelectButton.Alignment = widget.ButtonAlignTrailing
	dateSelectButton.Importance = widget.LowImportance

	arrowIcon := canvas.NewText("›", grayColor)
	arrowIcon.TextSize = 22
	arrowIcon.TextStyle = fyne.TextStyle{Bold: true}

	// 组合日期选择行
	startTimeRow := container.NewBorder(
		nil, nil,
		startTimeLabel,
		container.NewHBox(arrowIcon),
		dateSelectButton,
	)

	// 创建币种行 - 使用可点击的按钮
	currencyLabel := canvas.NewText("币种", blackColor)
	currencyLabel.TextSize = 16
	currencyLabel.TextStyle = fyne.TextStyle{Bold: true}

	// 创建币种选择按钮
	currencySelectButton := widget.NewButton("人民币 (CNY)", func() {
		// 币种选择逻辑
	})
	currencySelectButton.Alignment = widget.ButtonAlignTrailing
	currencySelectButton.Importance = widget.LowImportance

	arrowIcon2 := canvas.NewText("›", grayColor)
	arrowIcon2.TextSize = 22
	arrowIcon2.TextStyle = fyne.TextStyle{Bold: true}

	// 组合币种选择行
	currencyRow := container.NewBorder(
		nil, nil,
		currencyLabel,
		container.NewHBox(arrowIcon2),
		currencySelectButton,
	)

	// 添加更多垂直空间
	// 使用空白矩形替代Spacer
	nameSpacer := canvas.NewRectangle(color.Transparent)
	nameSpacer.SetMinSize(fyne.NewSize(0, 10))

	nameDivider := canvas.NewLine(lightGrayColor)
	nameDivider.StrokeWidth = 1

	// 创建账户资产行
	assetLabel := canvas.NewText("账户资产", blackColor)
	assetLabel.TextSize = 16
	assetLabel.TextStyle = fyne.TextStyle{Bold: true}

	assetPlaceholder := canvas.NewText("请输入金额", grayColor)
	assetPlaceholder.TextSize = 16
	assetPlaceholder.Alignment = fyne.TextAlignTrailing

	pencilIcon2 := canvas.NewText("✎", grayColor)
	pencilIcon2.TextSize = 16

	assetRow := container.NewBorder(
		nil, nil,
		assetLabel,
		container.NewHBox(assetPlaceholder, pencilIcon2),
	)

	assetDivider := canvas.NewLine(lightGrayColor)
	assetDivider.StrokeWidth = 1

	// 创建开始时间行
	startTimeLabel := canvas.NewText("开始时间", blackColor)
	startTimeLabel.TextSize = 16
	startTimeLabel.TextStyle = fyne.TextStyle{Bold: true}

	// 获取当前日期
	now := time.Now()
	dateStr := fmt.Sprintf("%d-%02d-%02d (今天)", now.Year(), now.Month(), now.Day())

	dateValue := canvas.NewText(dateStr, blackColor)
	dateValue.TextSize = 16
	dateValue.Alignment = fyne.TextAlignTrailing

	arrowIcon := canvas.NewText("›", grayColor)
	arrowIcon.TextSize = 22
	arrowIcon.TextStyle = fyne.TextStyle{Bold: true}

	startTimeRow := container.NewBorder(
		nil, nil,
		startTimeLabel,
		container.NewHBox(dateValue, arrowIcon),
	)

	startTimeDivider := canvas.NewLine(lightGrayColor)
	startTimeDivider.StrokeWidth = 1

	// 创建币种行
	currencyLabel := canvas.NewText("币种", blackColor)
	currencyLabel.TextSize = 16
	currencyLabel.TextStyle = fyne.TextStyle{Bold: true}

	currencyValue := canvas.NewText("人民币 (CNY)", blackColor)
	currencyValue.TextSize = 16
	currencyValue.Alignment = fyne.TextAlignTrailing

	arrowIcon2 := canvas.NewText("›", grayColor)
	arrowIcon2.TextSize = 22
	arrowIcon2.TextStyle = fyne.TextStyle{Bold: true}

	currencyRow := container.NewBorder(
		nil, nil,
		currencyLabel,
		container.NewHBox(currencyValue, arrowIcon2),
	)

	currencyDivider := canvas.NewLine(lightGrayColor)
	currencyDivider.StrokeWidth = 1

	// 创建确认添加按钮 - 优化为更接近图片的效果
	confirmBg := canvas.NewRectangle(pinkColor)
	confirmBg.SetMinSize(fyne.NewSize(0, 50))
	confirmBg.CornerRadius = 25

	confirmText := canvas.NewText("确认添加", whiteColor)
	confirmText.TextSize = 16
	confirmText.TextStyle = fyne.TextStyle{Bold: true}

	confirmButton := widget.NewButton("", func() {
		// 添加账户逻辑
	})
	confirmButton.Importance = widget.LowImportance

	confirmContainer := container.NewStack(
		confirmBg,
		container.NewCenter(confirmText),
		confirmButton,
	)

	// 添加底部空间
	// 使用空白矩形替代Spacer
	bottomSpacer := canvas.NewRectangle(color.Transparent)
	bottomSpacer.SetMinSize(fyne.NewSize(0, 20))

	// 创建主内容区 - 优化布局和间距
	content := container.NewVBox(
		topBar,
		iconSpacer,
		iconContainer,
		nameSpacer,
		container.NewPadded(nameRow),
		nameDivider,
		container.NewPadded(assetRow),
		assetDivider,
		container.NewPadded(startTimeRow),
		startTimeDivider,
		container.NewPadded(currencyRow),
		currencyDivider,
		layout.NewSpacer(),
		container.NewPadded(confirmContainer),
		bottomSpacer,
	)

	// 设置内容并显示窗口
	w.SetContent(content)
	return w
}

func main() {
	w := createAddAccountPage()
	w.ShowAndRun()
}
