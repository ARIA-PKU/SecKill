package controllers

import (
	"encoding/json"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"SecKill_Product/common"
	"SecKill_Product/datamodels"
	"SecKill_Product/services"
	"strconv"
	"time"
)

type OrderController struct {
	Ctx iris.Context
	OrderService services.IOrderService
}

func (o *OrderController) GetAll() mvc.View {

	orderMap, err := o.OrderService.GetAllOrderInfo()
	if err!= nil {
		o.Ctx.Application().Logger().Debug("查询订单信息失败")
		o.Ctx.Application().Logger().Debug(err)
	}

	//IP
	localIp, err := common.GetIntranceIp()
	//idString := o.Ctx.PostValue("id")
	//productId, err := strconv.ParseInt(idString, 10, 64)
	token := o.Ctx.GetCookie("token")

	_, claims, _ := common.ParseToken(token)

	// 24小时制
	timeObj := time.Now()
	var NowTime = timeObj.Format("2006/01/02 15:04:05")

	log := datamodels.Log{
		Name :claims.Name,
		ProductID:0,
		IP :localIp,
		Operation: "login",
		Time:NowTime,
	}
	_ = o.OrderService.InsertLog(log)
	o.Ctx.Application().Logger().Debug("ip地址为"+localIp+"登录成功")

	return mvc.View{
		Name:"/order/view.html",
		Data:iris.Map{
			"order": orderMap,
		},
	}
}

//面向权限较低的用户的界面
func (o *OrderController) GetAllorder() mvc.View {

	orderMap, err := o.OrderService.GetAllOrderInfo()
	if err!= nil {
		o.Ctx.Application().Logger().Debug("查询订单信息失败")
		o.Ctx.Application().Logger().Debug(err)
	}

	//IP
	localIp, err := common.GetIntranceIp()
	//idString := o.Ctx.PostValue("id")
	//productId, err := strconv.ParseInt(idString, 10, 64)
	token := o.Ctx.GetCookie("token")

	_, claims, _ := common.ParseToken(token)

	// 24小时制
	timeObj := time.Now()
	var NowTime = timeObj.Format("2006/01/02 15:04:05")

	log := datamodels.Log{
		Name :claims.Name,
		ProductID:0,
		IP :localIp,
		Operation: "login",
		Time:NowTime,
	}
	_ = o.OrderService.InsertLog(log)
	o.Ctx.Application().Logger().Debug("ip地址为"+localIp+"登录成功")

	return mvc.View{
		Layout: "shared/layout2.html",
		Name:"/order/view.html",
		Data:iris.Map{
			"order": orderMap,
		},
	}
}

func (o *OrderController) GetManager() mvc.View  {
	idString := o.Ctx.URLParam("id")
	id, err := strconv.ParseInt(idString, 10, 16)
	if err != nil {
		o.Ctx.Application().Logger().Debug(err)
	}
	order, err := o.OrderService.GetOrderByID(id)
	if err!= nil {
		o.Ctx.Application().Logger().Debug(err)
	}

	//IP
	localIp, err := common.GetIntranceIp()
	//idString := o.Ctx.PostValue("id")
	//productId, err := strconv.ParseInt(idString, 10, 64)
	token := o.Ctx.GetCookie("token")

	_, claims, _ := common.ParseToken(token)

	// 24小时制
	timeObj := time.Now()
	var NowTime = timeObj.Format("2006/01/02 15:04:05")

	log := datamodels.Log{
		Name :claims.Name,
		ProductID:order.ID,
		IP :localIp,
		Operation: "check order",
		Time:NowTime,
	}
	_ = o.OrderService.InsertLog(log)
	o.Ctx.Application().Logger().Debug("ip地址为"+localIp+"查看订单")

	return mvc.View{
		Name:"order/manager.html",
		Data:iris.Map{
			"order": order,
		},
	}
}

//面向权限较低的用户的界面
func (o *OrderController) GetManagerorder() mvc.View  {
	idString := o.Ctx.URLParam("id")
	id, err := strconv.ParseInt(idString, 10, 16)
	if err != nil {
		o.Ctx.Application().Logger().Debug(err)
	}
	order, err := o.OrderService.GetOrderByID(id)
	if err!= nil {
		o.Ctx.Application().Logger().Debug(err)
	}

	//IP
	localIp, err := common.GetIntranceIp()
	//idString := o.Ctx.PostValue("id")
	//productId, err := strconv.ParseInt(idString, 10, 64)
	token := o.Ctx.GetCookie("token")

	_, claims, _ := common.ParseToken(token)

	// 24小时制
	timeObj := time.Now()
	var NowTime = timeObj.Format("2006/01/02 15:04:05")

	log := datamodels.Log{
		Name :claims.Name,
		ProductID:order.ID,
		IP :localIp,
		Operation: "check order",
		Time:NowTime,
	}
	_ = o.OrderService.InsertLog(log)
	o.Ctx.Application().Logger().Debug("ip地址为"+localIp+"查看订单成功")

	return mvc.View{
		Layout: "shared/layout2.html",
		Name:"order/manager.html",
		Data:iris.Map{
			"order": order,
		},
	}
}

func (o *OrderController) PostUpdate()  {
	order := &datamodels.Order{}
	o.Ctx.Request().ParseForm()
	dec := common.NewDecoder(&common.DecoderOptions{"homework", false, false})
	if err := dec.Decode(o.Ctx.Request().Form, order);err!= nil{
		o.Ctx.Application().Logger().Debug(err)
	}
	err := o.OrderService.UpdateOrder(order)
	if err!= nil{
		o.Ctx.Application().Logger().Debug(err)
	}

    //IP
	localIp, err := common.GetIntranceIp()
	//idString := o.Ctx.PostValue("id")
	//productId, err := strconv.ParseInt(idString, 10, 64)
	token := o.Ctx.GetCookie("token")

	_, claims, _ := common.ParseToken(token)

	// 24小时制
	timeObj := time.Now()
	var NowTime = timeObj.Format("2006/01/02 15:04:05")

	log := datamodels.Log{
		Name :claims.Name,
		ProductID:order.ID,
		IP :localIp,
		Operation: "update order",
		Time:NowTime,
	}
	_ = o.OrderService.InsertLog(log)
	o.Ctx.Application().Logger().Debug("ip地址为"+localIp+"更新订单成功，ID为"+ strconv.FormatInt(order.ID, 10))
	o.Ctx.Redirect("/order/allorder")
}

func (o *OrderController) PostDelete() []byte{
	//idString := o.Ctx.URLParam("id")
	localIp, err := common.GetIntranceIp()
	idString := o.Ctx.PostValue("id")
	productId, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		o.Ctx.Application().Logger().Debug(err)
	}
	isOk := o.OrderService.DeleteOrderByID(productId)
	if isOk {
		o.Ctx.Application().Logger().Debug("ip地址为"+localIp+"删除订单成功，ID为"+idString)

		token := o.Ctx.GetCookie("token")
		//if token == "" {
		//	o.Ctx.Application().Logger().Debug("必须先登录")
		//	o.Ctx.Redirect("http://127.0.0.1:8081/login")
		//}
		_, claims, _ := common.ParseToken(token)
		//fmt.Println("用户权限为"+claims.Authority)
		//fmt.Println("用户电话为"+claims.Telephone)
		//fmt.Println("用户名为"+ claims.Name)


		// 24小时制
		timeObj := time.Now()
		var NowTime = timeObj.Format("2006/01/02 15:04:05")
		//fmt.Println(NowTime) // 2020/04/26 17:48:53

		log := datamodels.Log{
			Name :claims.Name,
			ProductID:productId,
			IP :localIp,
			Operation: "delete order",
			Time:NowTime,
		}
		_ = o.OrderService.InsertLog(log)
		//if err!=nil{
		//	fmt.Println("修改失败")
		//}
		response, _ := json.Marshal(iris.Map{"code": 200,"msg": "删除订单成功"})

		return response
	}else {
		o.Ctx.Application().Logger().Debug("删除订单失败，ID为"+idString)
		response, _ := json.Marshal(iris.Map{"code": 201,"msg": "删除订单失败"})
		return response
	}
	//o.Ctx.Redirect("/order/all")
}