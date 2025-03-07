package main

import (
	"fmt"
	"image/color"
	"strconv"
	"study/StockCalculator/server"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// 自定义主题
// 自定义主题增强版，支持跟随系统主题和深色/浅色模式
type myTheme struct {
	fyne.Theme
	followSystem bool // 是否跟随系统主题
	darkMode     bool // 是否使用深色模式
}

func newMyTheme(followSystem bool, darkMode bool) fyne.Theme {
	return &myTheme{
		Theme:        theme.DefaultTheme(),
		followSystem: followSystem,
		darkMode:     darkMode,
	}
}

func (m *myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	// 如果设置为跟随系统，则使用系统变体
	if m.followSystem {
		return m.Theme.Color(name, variant)
	}

	// 根据深色/浅色模式选择变体
	selectedVariant := theme.VariantLight
	if m.darkMode {
		selectedVariant = theme.VariantDark
	}

	// 根据深色/浅色模式自定义颜色
	if m.darkMode {
		// 深色模式颜色
		switch name {
		case theme.ColorNameBackground:
			return color.NRGBA{R: 30, G: 30, B: 35, A: 255} // 深色背景
		case theme.ColorNameButton:
			return color.NRGBA{R: 80, G: 120, B: 200, A: 255} // 深色按钮
		case theme.ColorNamePrimary:
			return color.NRGBA{R: 70, G: 130, B: 230, A: 255} // 深色主色调
		case theme.ColorNameForeground:
			return color.NRGBA{R: 220, G: 220, B: 230, A: 255} // 深色前景
		case theme.ColorNameShadow:
			return color.NRGBA{R: 0, G: 0, B: 0, A: 50} // 深色阴影
		}
	} else {
		// 浅色模式颜色
		switch name {
		case theme.ColorNameBackground:
			return color.NRGBA{R: 248, G: 250, B: 252, A: 255} // 更柔和的背景色
		case theme.ColorNameButton:
			return color.NRGBA{R: 66, G: 133, B: 244, A: 255} // 主按钮蓝色
		case theme.ColorNamePrimary:
			return color.NRGBA{R: 50, G: 120, B: 220, A: 255} // 主色调
		case theme.ColorNameForeground:
			return color.NRGBA{R: 40, G: 40, B: 50, A: 255} // 更深的前景色
		case theme.ColorNameShadow:
			return color.NRGBA{R: 0, G: 0, B: 0, A: 30} // 更明显的阴影
		}
	}

	return m.Theme.Color(name, selectedVariant)
}

// 定义按钮颜色常量
var (
	primaryButtonColor   = color.NRGBA{R: 66, G: 133, B: 244, A: 255}  // 蓝色 - 主要操作
	secondaryButtonColor = color.NRGBA{R: 234, G: 67, B: 53, A: 255}   // 红色 - 重置/取消
	successButtonColor   = color.NRGBA{R: 52, G: 168, B: 83, A: 255}   // 绿色 - 成功/确认
	settingsButtonColor  = color.NRGBA{R: 100, G: 100, B: 150, A: 255} // 灰蓝色 - 设置
)

// 添加主题设置功能
func createThemeSettings(a fyne.App, w fyne.Window) fyne.CanvasObject {
	// 获取当前主题设置
	currentTheme := a.Settings().Theme().(*myTheme)

	// 创建主题模式选择单选按钮
	themeGroup := widget.NewRadioGroup([]string{"自定义主题", "跟随系统主题"}, func(selected string) {
		// 获取当前主题
		currentTheme := a.Settings().Theme().(*myTheme)

		if selected == "跟随系统主题" {
			a.Settings().SetTheme(newMyTheme(true, currentTheme.darkMode))
		} else {
			a.Settings().SetTheme(newMyTheme(false, currentTheme.darkMode))
		}
	})

	// 创建深色/浅色模式选择单选按钮
	modeGroup := widget.NewRadioGroup([]string{"浅色模式", "深色模式"}, func(selected string) {
		// 获取当前主题
		currentTheme := a.Settings().Theme().(*myTheme)

		if selected == "深色模式" {
			a.Settings().SetTheme(newMyTheme(currentTheme.followSystem, true))
		} else {
			a.Settings().SetTheme(newMyTheme(currentTheme.followSystem, false))
		}
	})

	// 根据当前主题设置默认选择
	if currentTheme.followSystem {
		themeGroup.Selected = "跟随系统主题"
	} else {
		themeGroup.Selected = "自定义主题"
	}

	if currentTheme.darkMode {
		modeGroup.Selected = "深色模式"
	} else {
		modeGroup.Selected = "浅色模式"
	}

	// 创建设置面板
	settingsForm := container.NewVBox(
		widget.NewLabel("主题设置"),
		themeGroup,
		widget.NewSeparator(),
		widget.NewLabel("显示模式"),
		modeGroup,
	)

	return container.NewPadded(settingsForm)
}

// go run gui/main.go
// go run -tags mobile gui/main.go
// fyne package -os android -appID com.wxw9868.stockcalculator -name "Stock Calculator" -icon stock.png
func main() {
	// 创建Fyne应用并设置主题
	a := app.New()
	a.Settings().SetTheme(newMyTheme(false, false)) // 默认使用自定义浅色主题
	w := a.NewWindow("股票交易计算器")

	// 检测是否在移动设备上运行
	isMobile := fyne.CurrentDevice().IsMobile()

	// 根据设备类型调整窗口大小
	if isMobile {
		// 移动设备使用全屏
		w.SetFullScreen(true)
	} else {
		// 桌面设备使用固定尺寸
		w.Resize(fyne.NewSize(1000, 650))
	}

	// 创建标题和副标题，使用canvas.Text以支持文本大小设置
	titleText := canvas.NewText("股票交易计算器", theme.ForegroundColor())
	titleText.TextSize = 28
	titleText.Alignment = fyne.TextAlignCenter
	titleText.TextStyle = fyne.TextStyle{Bold: true}

	subtitleText := canvas.NewText("快速计算股票交易成本与收益", theme.ForegroundColor())
	subtitleText.TextSize = 18
	subtitleText.Alignment = fyne.TextAlignCenter
	subtitleText.TextStyle = fyne.TextStyle{Italic: true}

	// 创建主题变化监听函数
	updateTextColors := func() {
		titleText.Color = theme.ForegroundColor()
		subtitleText.Color = theme.ForegroundColor()
		titleText.Refresh()
		subtitleText.Refresh()
	}

	// 注册主题变化监听 - 修复方法名
	listener := make(chan fyne.Settings)
	go func() {
		for range listener {
			updateTextColors()
		}
	}()
	a.Settings().AddChangeListener(listener)

	// 初始调用一次确保颜色正确
	updateTextColors()

	// 创建输入控件并美化
	rateEntry := widget.NewEntry()
	rateEntry.SetPlaceHolder("例如: 0.0185")
	rateEntry.SetText("0.0185")
	styleEntry(rateEntry)

	buyPriceEntry := widget.NewEntry()
	buyPriceEntry.SetPlaceHolder("例如: 10.50")
	styleEntry(buyPriceEntry)

	sellPriceEntry := widget.NewEntry()
	sellPriceEntry.SetPlaceHolder("例如: 11.20")
	styleEntry(sellPriceEntry)

	countEntry := widget.NewEntry()
	countEntry.SetPlaceHolder("例如: 1000")
	styleEntry(countEntry)

	// 为移动设备优化输入控件
	if isMobile {
		// 增大输入框高度以便于触控
		rateEntry.SetMinRowsVisible(2)
		buyPriceEntry.SetMinRowsVisible(2)
		sellPriceEntry.SetMinRowsVisible(2)
		countEntry.SetMinRowsVisible(2)
	}

	// 创建结果显示区域
	resultLabel := widget.NewLabel("")
	resultLabel.Wrapping = fyne.TextWrapWord
	resultLabel.Alignment = fyne.TextAlignLeading
	resultLabel.TextStyle = fyne.TextStyle{Monospace: true}

	// 创建状态标签
	statusLabel := widget.NewLabel("")
	statusLabel.Alignment = fyne.TextAlignCenter

	// 创建设置按钮
	settingsButton := newColoredButtonWithIcon("设置", theme.SettingsIcon(), settingsButtonColor, func() {
		// 创建设置对话框
		settingsDialog := dialog.NewCustom("应用设置", "关闭", createThemeSettings(a, w), w)
		settingsDialog.Resize(fyne.NewSize(300, 200))
		settingsDialog.Show()
	})

	// 创建计算按钮 - 使用自定义样式
	calculateButton := newColoredButton("计算交易结果", primaryButtonColor, func() {
		// 获取输入值
		rateStr := rateEntry.Text
		buyPriceStr := buyPriceEntry.Text
		sellPriceStr := sellPriceEntry.Text
		countStr := countEntry.Text

		// 转换为数值
		rate, err1 := strconv.ParseFloat(rateStr, 64)
		buyPrice, err2 := strconv.ParseFloat(buyPriceStr, 64)
		sellPrice, err3 := strconv.ParseFloat(sellPriceStr, 64)
		count, err4 := strconv.ParseFloat(countStr, 64)

		// 检查输入是否有效
		if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
			showStatus(statusLabel, "⚠️ 错误：请输入有效的数字", color.NRGBA{R: 234, G: 67, B: 53, A: 255})
			resultLabel.SetText("")
			return
		}

		if buyPrice <= 0 || sellPrice <= 0 || count <= 0 {
			showStatus(statusLabel, "⚠️ 错误：所有参数都必须是正值！", color.NRGBA{R: 234, G: 67, B: 53, A: 255})
			resultLabel.SetText("")
			return
		}

		// 计算结果
		stock := server.NewStock(rate / 100)
		buyTransaction := server.NewTransaction(stock, buyPrice, count, server.Buy)
		sellTransaction := server.NewTransaction(stock, sellPrice, count, server.Sell)

		// 计算手续费
		handlingFee := buyTransaction.TotalFee + sellTransaction.TotalFee

		// 计算收益率
		yield := (sellPrice - buyPrice) / buyPrice * 100

		// 计算净利润
		netProfit := sellTransaction.TotalAmount - buyTransaction.TotalAmount - handlingFee

		// 设置状态标签
		if netProfit > 0 {
			showStatus(statusLabel, "✅ 计算完成：交易盈利", nil)
		} else {
			showStatus(statusLabel, "❗ 计算完成：交易亏损", nil)
		}

		// 显示结果 - 移动设备上使用更简洁的格式
		var result string
		if isMobile {
			result = fmt.Sprintf("📊 交易详情\n\n"+
				"💰 买入: %.2f元 x %.0f股\n  金额: %.2f元\n  费用: %.2f元\n\n"+
				"💰 卖出: %.2f元 x %.0f股\n  金额: %.2f元\n  费用: %.2f元\n\n"+
				"📝 汇总\n  手续费: %.2f元\n  收益率: %.2f%%\n  净利润: %.2f元",
				buyTransaction.Price, buyTransaction.Count, buyTransaction.TotalAmount, buyTransaction.TotalFee,
				sellTransaction.Price, sellTransaction.Count, sellTransaction.TotalAmount, sellTransaction.TotalFee,
				handlingFee, yield, netProfit)
		} else {
			// 使用新的格式化函数
			result = formatResultWithColor(
				buyTransaction.Price, buyTransaction.Count, buyTransaction.TotalAmount, buyTransaction.TotalFee,
				sellTransaction.Price, sellTransaction.Count, sellTransaction.TotalAmount, sellTransaction.TotalFee,
				handlingFee, yield, netProfit)
		}

		resultLabel.SetText(result)
	})

	// 创建重置按钮
	resetButton := newColoredButtonWithIcon("重置", theme.ContentClearIcon(), secondaryButtonColor, func() {
		buyPriceEntry.SetText("")
		sellPriceEntry.SetText("")
		countEntry.SetText("")
		resultLabel.SetText("")
		showStatus(statusLabel, "🔄 已重置所有输入", nil)
	})

	// 创建示例按钮
	exampleButton := newColoredButton("加载示例", successButtonColor, func() {
		buyPriceEntry.SetText("10.50")
		sellPriceEntry.SetText("11.20")
		countEntry.SetText("1000")
		showStatus(statusLabel, "📋 已加载示例数据", nil)
	})

	// 创建按钮容器容器，添加设置按钮：
	// 创建按钮容器
	buttonContainer := container.NewHBox(
		settingsButton,
		layout.NewSpacer(),
		exampleButton,
		resetButton,
		calculateButton,
		layout.NewSpacer(),
	)

	// 创建分隔线
	divider := canvas.NewLine(color.NRGBA{R: 200, G: 200, B: 200, A: 255})
	divider.StrokeWidth = 1

	// 创建标签样式
	rateLabel := styledLabel("佣金费率 (%)")
	buyPriceLabel := styledLabel("买入价格 (元)")
	sellPriceLabel := styledLabel("卖出价格 (元)")
	countLabel := styledLabel("股票数量 (股)")

	// 创建响应式表单布局 - 使用已定义的标签变量
	form := container.NewVBox(
		container.NewGridWithColumns(2, rateLabel, rateEntry),
		container.NewPadded(layout.NewSpacer()), // 添加间距
		container.NewGridWithColumns(2, buyPriceLabel, buyPriceEntry),
		container.NewPadded(layout.NewSpacer()), // 添加间距
		container.NewGridWithColumns(2, sellPriceLabel, sellPriceEntry),
		container.NewPadded(layout.NewSpacer()), // 添加间距
		container.NewGridWithColumns(2, countLabel, countEntry),
		container.NewPadded(layout.NewSpacer()), // 添加间距
		container.NewPadded(buttonContainer),
	)

	// 创建无边框面板
	leftPanel := createModernPanel("交易参数", "请输入交易相关信息", form)

	// 创建右侧卡片 - 使用滚动容器并设置合适的尺寸
	resultScroll := container.NewScroll(resultLabel)

	// 根据设备类型调整滚动区域大小
	if isMobile {
		resultScroll.SetMinSize(fyne.NewSize(300, 300))
	} else {
		resultScroll.SetMinSize(fyne.NewSize(450, 450))
	}

	rightPanel := createModernPanel("计算结果", "交易详情将显示在这里", resultScroll)

	// 创建底部状态栏 - 使用主题适配的分隔线
	statusBar := container.NewVBox(
		widget.NewSeparator(),
		container.NewPadded(statusLabel),
	)

	// 根据设备类型创建不同的布局
	var content fyne.CanvasObject
	if isMobile {
		// 移动设备使用垂直布局，优化间距和组件大小

		// 调整移动设备上的文本大小
		titleText.TextSize = 22
		subtitleText.TextSize = 14

		// 创建更紧凑的表单布局
		compactForm := container.NewVBox(
			container.NewGridWithColumns(2, rateLabel, rateEntry),
			container.NewGridWithColumns(2, buyPriceLabel, buyPriceEntry),
			container.NewGridWithColumns(2, sellPriceLabel, sellPriceEntry),
			container.NewGridWithColumns(2, countLabel, countEntry),
			container.NewPadded(buttonContainer),
		)

		// 替换左侧面板内容为更紧凑的表单
		leftCompactPanel := createMobilePanel("交易参数", compactForm)

		// 调整结果区域大小
		resultScroll.SetMinSize(fyne.NewSize(280, 250))

		// 创建更紧凑的结果面板
		rightCompactPanel := createMobilePanel("计算结果", resultScroll)

		// 创建更紧凑的移动布局
		content = container.NewVBox(
			container.NewPadded(titleText),
			container.NewPadded(subtitleText),
			container.NewPadded(leftCompactPanel),
			container.NewPadded(rightCompactPanel),
			statusBar,
		)
	} else {
		// 桌面设备使用水平分割布局
		split := container.NewHSplit(leftPanel, rightPanel)
		split.Offset = 0.35

		content = container.NewVBox(
			container.NewPadded(titleText),
			container.NewPadded(subtitleText),
			container.NewPadded(split),
			statusBar,
		)
	}

	// 设置窗口内容 - 移除硬编码的背景，使用主题背景
	w.SetContent(container.NewPadded(content))

	w.ShowAndRun()
}

// 创建自定义无边框面板 - 现代风格
// 创建按钮的辅助函数
func newColoredButton(text string, bgColor color.Color, tapped func()) *widget.Button {
	button := widget.NewButton(text, tapped)
	button.Importance = widget.HighImportance
	return button
}

// 创建带图标的按钮
func newColoredButtonWithIcon(text string, icon fyne.Resource, bgColor color.Color, tapped func()) *widget.Button {
	button := widget.NewButtonWithIcon(text, icon, tapped)
	button.Importance = widget.HighImportance
	return button
}

// 设置状态标签的文本
func showStatus(label *widget.Label, text string, textColor color.Color) {
	label.SetText(text)
}

// 美化标签
func styledLabel(text string) *widget.Label {
	label := widget.NewLabel(text)
	label.TextStyle = fyne.TextStyle{Bold: true}
	return label
}

// 美化输入框
func styleEntry(entry *widget.Entry) {
	entry.TextStyle = fyne.TextStyle{Monospace: true}
}

// 格式化结果，添加颜色标记
func formatResultWithColor(buyPrice, buyCount, buyAmount, buyFee, sellPrice, sellCount, sellAmount, sellFee, handlingFee, yield, netProfit float64) string {
	// 使用emoji和格式化来增强可读性
	profitStatus := "📈 盈利"
	if netProfit < 0 {
		profitStatus = "📉 亏损"
	}

	return fmt.Sprintf("📊 交易详情\n\n"+
		"💰 【买入】\n"+
		"  价格: %.2f 元\n"+
		"  数量: %.0f 股\n"+
		"  金额: %.2f 元\n"+
		"  费用: %.2f 元\n\n"+
		"💰 【卖出】\n"+
		"  价格: %.2f 元\n"+
		"  数量: %.0f 股\n"+
		"  金额: %.2f 元\n"+
		"  费用: %.2f 元\n\n"+
		"📝 汇总信息\n"+
		"  手续费: %.2f 元\n"+
		"  收益率: %.2f%%\n"+
		"  净利润: %.2f 元\n"+
		"  状态: %s",
		buyPrice, buyCount, buyAmount, buyFee,
		sellPrice, sellCount, sellAmount, sellFee,
		handlingFee, yield, netProfit, profitStatus)
}

// 创建现代风格面板
func createModernPanel(title, subtitle string, content fyne.CanvasObject) fyne.CanvasObject {
	// 创建标题和副标题，使用主题颜色
	titleText := widget.NewLabel(title)
	titleText.TextStyle = fyne.TextStyle{Bold: true}
	titleText.Alignment = fyne.TextAlignLeading

	subtitleText := widget.NewLabel(subtitle)
	subtitleText.TextStyle = fyne.TextStyle{Italic: true}
	subtitleText.Alignment = fyne.TextAlignLeading

	// 创建内容容器
	contentWithPadding := container.NewPadded(content)

	// 创建分隔线，使用主题分隔线
	divider := widget.NewSeparator()

	// 组合面板内容
	panelContent := container.NewVBox(
		container.NewPadded(titleText),
		container.NewPadded(subtitleText),
		container.NewPadded(divider),
		contentWithPadding,
	)

	// 使用Card容器，它会自动适应主题
	return widget.NewCard("", "", panelContent)
}

// 创建移动设备专用的紧凑面板
func createMobilePanel(title string, content fyne.CanvasObject) fyne.CanvasObject {
	// 创建标题，使用主题颜色
	titleText := widget.NewLabel(title)
	titleText.TextStyle = fyne.TextStyle{Bold: true}
	titleText.Alignment = fyne.TextAlignLeading

	// 创建内容容器，减少内边距
	contentWithPadding := container.New(layout.NewPaddedLayout(), content)

	// 创建分隔线，使用主题分隔线
	divider := widget.NewSeparator()

	// 组合面板内容，减少内边距
	panelContent := container.NewVBox(
		container.NewPadded(titleText),
		container.NewPadded(divider),
		contentWithPadding,
	)

	// 使用Card容器，它会自动适应主题
	return widget.NewCard("", "", panelContent)
}
