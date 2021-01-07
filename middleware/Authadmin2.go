package middleware

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"SecKill_Product/common"
)

func AuthAdmin2(ctx iris.Context)  {

	//storetoken := ctx.GetCookie("token")
	//_, claims, _ := common.ParseToken(storetoken)
	//DB := common.GetDB()
	//var user datamodels.Admin
	//userId := claims.UserId
	//DB.First(&user, userId)
	//if user.ID == 0 {
	//	ctx.Application().Logger().Debug("必须先登录")
	//	ctx.Redirect("/login")
	//	return
	//}
	//ctx.Application().Logger().Debug("用户ID " + strconv.FormatUint(uint64(userId), 10) + " 已经登录")
	//ctx.Next()
	token := ctx.GetCookie("token")
	if token == "" {
		ctx.Application().Logger().Debug("必须先登录")
		ctx.Redirect("http://127.0.0.1:8081/login")
		return
	}
	_, claims, _ := common.ParseToken(token)
	fmt.Println("用户权限为"+claims.Authority)
	if claims.Authority !="0"||claims.Authority !="1"{
		ctx.Application().Logger().Debug("必须先登录")
		ctx.Redirect("http://127.0.0.1:8081/login")
		return
	}
	//ctx.Application().Logger().Debug("用户ID " + claims.Authority + " 已经登录")
	ctx.Next()
}
