package main

import (
	"SecKill_Product/common"
	"SecKill_Product/fronted/web/controllers"
	"SecKill_Product/rabbitmq"
	"SecKill_Product/repositories"
	"SecKill_Product/services"
	"context"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"github.com/opentracing/opentracing-go/log"
	"time"
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
	// 4. 设置模板目标
	app.HandleDir("/public",  "./web/public")
	app.HandleDir("/html", "./web/htmlProductShow")
	//app.HandleDir("/js",  "./public/js")
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
	session := sessions.New(sessions.Config{Cookie: "helloworld", Expires: 60 * time.Minute,})

	// 5. 注册控制器
	// 注册user控制器
	userRepository := repositories.NewUserRepository("user", db)
	userService := services.NewUserService(userRepository)
	userPro := mvc.New(app.Party("/user"))
	userPro.Register(userService, ctx, session.Start)
	userPro.Register(userService, ctx)
	userPro.Handle(new(controllers.UserController))

	rabbitmq := rabbitmq.NewRabbitMQSimple("seckillProduct")



	// 注册product控制器
	productRepository := repositories.NewProductManager("product", db)
	orderRepository := repositories.NewOrderManagerRepository("order", db)
	productService := services.NewProductService(productRepository)
	orderService := services.NewOrderService(orderRepository)
	proProduct := app.Party("/product")
	//proProduct.Use(middleware.AuthConProduct) // 添加登录验证
	//proProduct.Use(middleware.RateMiddleware)//添加IP限制功能，实现限流
	product := mvc.New(proProduct)
	product.Register(productService, orderService, ctx, session.Start)
	product.Register(productService, orderService, rabbitmq, ctx)
	product.Handle(new(controllers.ProductController))

	// 6. 启动服务
	app.Run(
		iris.Addr("127.0.0.1:8082"),
		//iris.WithoutVersionChecker,
		// 忽略服务器错误
		iris.WithoutServerError(iris.ErrServerClosed),
		// 尽可能优化
		iris.WithOptimizations,
	)
}
