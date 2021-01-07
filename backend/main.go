package main

import (
	"SecKill_Product/backend/web/controllers"
	"SecKill_Product/common"
	"SecKill_Product/middleware"
	"SecKill_Product/repositories"
	"SecKill_Product/services"
	"context"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/opentracing/opentracing-go/log"
)

func main() {
	// 1. 创建iris实例
	app := iris.New()
	// 2. 设置debug模式
	app.Logger().SetLevel("debug")
	// 3. 设置注册模板
	template := iris.HTML(
		"./web/views", ".html").Layout(
		"shared/layout.html").Reload(true)
	app.RegisterView(template)
    //template2为权限0的layout
	template2 := iris.HTML(
		"./web/views", ".html").Layout(
		"shared/layout2.html").Reload(true)
	app.RegisterView(template2)
	// 4. 设置模板目标
	app.HandleDir("/assets", "./web/assets")
	// 出现异常跳转到指定页面
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message", ctx.Values().GetStringDefault("message", "访问的页面出错！"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})
	db, err := common.NewMysqlConn()
	if err != nil {
		log.Error(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// 5. 注册控制器
	//产品
	productRepository := repositories.NewProductManager("product", db)
	productService := services.NewProductService(productRepository)
	productParty := app.Party("/product")
	productParty.Use(middleware.AuthAdmin)//添加登录验证，product修改所需权限高
	product := mvc.New(productParty)
	product.Register(ctx, productService)
	product.Handle(new(controllers.ProductController))
    //日志
	logRepository := repositories.NewLogManager("log", db)
	logService := services.NewlogService(logRepository)
	logParty := app.Party("/log")
	logParty.Use(middleware.AuthAdmin)//添加登录验证，product修改所需权限高
	log := mvc.New(logParty)
	log.Register(ctx, logService)
	log.Handle(new(controllers.LogController))
    //订单
	orderRepository := repositories.NewOrderManagerRepository("order", db)
	orderService := services.NewOrderService(orderRepository)
	orderParty := app.Party("/order")
	productParty.Use(middleware.AuthAdmin2)//添加登录验证,order查看所需权限低
	order := mvc.New(orderParty)
	order.Register(ctx,orderService)
	order.Handle(new(controllers.OrderController))

	// 6. 启动服务
	app.Run(
		iris.Addr("127.0.0.1:8080"),
		//iris.WithoutVersionChecker,
		// 忽略服务器错误
		iris.WithoutServerError(iris.ErrServerClosed),
		// 尽可能优化
		iris.WithOptimizations,
	)
}

// go mod init
// 编辑文件
// set GO111MODULE=on
// go mod tidy
