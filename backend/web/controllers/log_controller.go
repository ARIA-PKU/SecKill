package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"SecKill_Product/services"
)

type LogController struct {
	Ctx iris.Context
	ILogService services.ILogService
}

func (o *LogController) GetLog() mvc.View {

	logMap, err := o.ILogService.GetAllLogInfo()
	if err!= nil {
		o.Ctx.Application().Logger().Debug("查询订单信息失败")
		o.Ctx.Application().Logger().Debug(err)
	}



	return mvc.View{
		Name:"/log/view.html",
		Data:iris.Map{
			"log": logMap,
		},
	}
}