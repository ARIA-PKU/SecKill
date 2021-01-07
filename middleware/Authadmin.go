package middleware

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"SecKill_Product/common"
)

func AuthAdmin(ctx iris.Context)  {

	token := ctx.GetCookie("token")
	if token == "" {
		ctx.Application().Logger().Debug("必须先登录")
		ctx.Redirect("http://127.0.0.1:8081/login")
		return
	}
	_, claims, _ := common.ParseToken(token)
	fmt.Println("用户权限为"+claims.Authority)
	fmt.Println("用户电话为"+claims.Telephone)

	//DB := common.GetDB()
	//
	//var user datamodels.Admin
	//DB.Where("telephone = ?", claims.Telephone).First(&user)
	if claims.Authority !="1"{
		ctx.Application().Logger().Debug("必须先登录")
		ctx.Redirect("http://127.0.0.1:8081/login")
		return
	}
	ctx.Next()
}
