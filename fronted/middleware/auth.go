package middleware

import (
	"SecKill_Product/common"
	"SecKill_Product/encrypt"
	"fmt"
	"github.com/kataras/iris/v12"
	"log"
	"time"
)

func AuthConProduct(ctx iris.Context)  {
	uid := ctx.GetCookie("uid")
	if uid == "" {
		ctx.Application().Logger().Debug("必须先登录")
		ctx.Redirect("/user/login")
		return
	}

	// 获取用户加密串
	signCookie := ctx.GetCookie("sign")

	// 对信息进行解密
	signByte, err := encrypt.DePwdCode(signCookie)
	if err != nil {
		ctx.Application().Logger().Debug("加密内容已篡改")
	}

	if checkInfo(uid, string(signByte)) {
		ctx.Application().Logger().Debug("验证成功")
	}

	ctx.Application().Logger().Debug("用户ID " + uid + " 已经登录")
	ctx.Next()
}

// 自定义逻辑判断
func checkInfo(checkStr string, signStr string) bool {
	if checkStr == signStr {
		return true
	}
	return false
}


//限制访问
func RateMiddleware(ctx iris.Context) {
	limiter, err := common.NewLimiter("redis://localhost:6379")
	if err != nil {
		fmt.Println("main中redis连接")
	}
	// 如果ip请求连接数在两秒内超过5次，返回429并抛出error
	if !limiter.Allow(ctx.Request().RemoteAddr, 1, 5*time.Second) {
		log.Println("too many requests")
		ctx.Redirect("http://127.0.0.1:8082/product/error")
		return
	}
	//log.Println("限制访问执行")
	ctx.Next()
}

