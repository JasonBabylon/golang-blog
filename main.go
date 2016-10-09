package main

import (
	"bee_app/controllers"
	"bee_app/models"
	_ "bee_app/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"os"
	"github.com/beego/i18n"
)

func init() {
	models.RegisterDB()
}

func main() {
	i18n.SetMessage("en-US", "conf/locale_en-US.ini")
	i18n.SetMessage("zh-CN", "conf/locale_zh-CN.ini")

	// 注册国际化模板函数
	beego.AddFuncMap("i18n", i18n.Tr)

	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/topic", &controllers.TopicController{})
	beego.AutoRouter(&controllers.TopicController{})
	beego.Router("/reply", &controllers.ReplyController{})
	beego.AutoRouter(&controllers.ReplyController{})

	// 创建附件目录
	os.Mkdir("attachment", os.ModePerm)
	// 作为静态文件
	//beego.SetStaticPath("/attachment", "attachment")
	// 作为单独一个控制器来处理
	beego.Router("/attachment/:all", &controllers.AttachController{})
	// 启动 beego
	beego.Run()
}
