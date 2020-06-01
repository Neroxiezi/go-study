package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"log"
)

func main() {
	tmw := new(TabMainWindow)
	if err := (MainWindow{
		Title:      "淘宝客微博管理工具",
		AssignTo:   &tmw.MainWindow,
		Background: SolidColorBrush{Color: walk.RGB(22, 119, 179)},
		MenuItems: []MenuItem{
			Menu{
				Text: "&微博",
				Items: []MenuItem{
					Action{
						Text:        "微博账号管理",
						OnTriggered: func() { tmw.NewWeiBo() },
					},
				},
			},
			Menu{
				Text: "&优惠券",
				Items: []MenuItem{
					Action{
						Text: "优惠券列表",
						OnTriggered: func() { tmw.CouponsList() },
					},
					Action{
						Text:        "更新优惠券",
						OnTriggered: func() { tmw.UpdateCoupons() },
					},
				},
			},
			Menu{
				Text: "&操作",
				Items: []MenuItem{
					Separator{},
					Action{
						Text:        "退出",
						OnTriggered: func() { tmw.Close() },
					},
				},
			},
			Menu{
				Text: "&帮助",
				Items: []MenuItem{
					Action{
						Text:        "关于",
						OnTriggered: tmw.about,
					},
				},
			},
		},
		Size:   Size{1100, 680},
		Layout: VBox{},
		Children: []Widget{
			WebView{
				AssignTo: &tmw.wv,
			},
		},
	}.Create()); err != nil {
		log.Fatal(err)
	}
	InitVehicleTypePage(tmw)
	tmw.Run()
}

type TabMainWindow struct {
	*walk.MainWindow
	TabWidget *walk.TabWidget
	wv        *walk.WebView
}

func (mw *TabMainWindow) about() {
	walk.MsgBox(mw, "关于", "淘宝客微博管理工具\r\n作者：PFinal南丞\r\n邮箱：lampxiezi@163.com", walk.MsgBoxIconInformation)
}

func InitVehicleTypePage(tmw *TabMainWindow) {
	tmw.wv.SetURL("file:///" + getCurrentDirectory() + "/tpl/init.html")
}

func (tmw *TabMainWindow) NewWeiBo() {
	tmw.wv.SetURL("file:///" + getCurrentDirectory() + "/tpl/weibo.html")
}

func (tmw *TabMainWindow)  CouponsList() {
    tmw.wv.SetURL("file:///" + getCurrentDirectory() + "/tpl/coupons.html")
}

func (tmw *TabMainWindow) UpdateCoupons() {
	tmw.wv.SetURL("file:///" + getCurrentDirectory() + "/tpl/load.html")
}
