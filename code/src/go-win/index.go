package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"log"
	//"os"
)

func main() {
	tmw := new(TabMainWindow)
	if err := (MainWindow{
		Icon:       "icon/logo.ico",
		Title:      "客户管理工具",
		AssignTo:   &tmw.MainWindow,
		Background: SolidColorBrush{Color: walk.RGB(22, 119, 179)},
		MenuItems: []MenuItem{
			Menu{
				Text: "&账号",
				Items: []MenuItem{
					Action{
						Text:        "账号管理",
						OnTriggered: func() { tmw.NewWeiBo() },
					},
					Action{
						Text:        "更新账号",
						OnTriggered: func() { tmw.UpWeiBo() },
					},
				},
			},
			Menu{
				Text: "&任务",
				Items: []MenuItem{
					Action{
						Text:        "任务列表",
						OnTriggered: func() { tmw.CouponsList() },
					},
					Action{
						Text:        "更新任务",
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
	//login_status, err := check_login()
	//if err != nil {
	//	tmw.wv.SetURL("file:///" + getCurrentDirectory() + "/tpl/login.html")
	//}
	//if login_status != true {
	//	tmw.wv.SetURL("file:///" + getCurrentDirectory() + "/tpl/login.html")
	//} else {
		tmw.wv.SetURL("file:///" + getCurrentDirectory() + "/tpl/weibo.html")
	//}

}

func (tmw *TabMainWindow) CouponsList() {
    //login_status, err := check_login()
	//if err != nil {
	//	tmw.wv.SetURL("file:///" + getCurrentDirectory() + "/tpl/login.html")
	//}
	//if login_status != true {
	//	tmw.wv.SetURL("file:///" + getCurrentDirectory() + "/tpl/login.html")
	//} else {
		tmw.wv.SetURL("file:///" + getCurrentDirectory() + "/tpl/coupons.html")
	//}
}

func (tmw *TabMainWindow) UpdateCoupons() {
	tmw.CouponsList()
}

func (tmw *TabMainWindow) UpWeiBo() {
	update_weibo_data()
	tmw.NewWeiBo()
}

//func check_login() (bool, error) {
//	str, err := os.Getwd()
//	if err != nil {
//		log.Fatal(err)
//	}
//	status, err := PathExists(str + "/tokenLock")
//	if err != nil {
//		return status, err
//	}
//	return status, nil
//}