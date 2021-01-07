package controllers

import (
	"SecKill_Product/datamodels"
	"SecKill_Product/rabbitmq"
	"SecKill_Product/services"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"html/template"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
)

type ProductController struct {
	Ctx            iris.Context
	ProductService services.IProductService
	OrderService   services.IOrderService
	RabbitMQ	   *rabbitmq.RabbitMQ
	Session        *sessions.Session
}

var (
	htmlOutPath  = "./web/htmlProductShow" // 生成的Html保存目录
	templatePath = "./web/views/template"  // 静态文件模板目录
)

func (p *ProductController)GetGenerate()  {
	// 1. 获取模板文件地址
	idString := p.Ctx.URLParam("productID")
	productID,err := strconv.Atoi(idString)
	if err!= nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	contentTpl, err := template.ParseFiles(filepath.Join(templatePath, "product.html"))
	if err != nil{
		p.Ctx.Application().Logger().Debug(err)
	}
	// 2. 获取html生成路径
	fileName := filepath.Join(htmlOutPath, "htmlProduct.html")
	// 3. 获取模板渲染数据
	product, err := p.ProductService.GetProductByID(int64(productID))
	if err != nil{
		p.Ctx.Application().Logger().Debug(err)
	}
	// 4.生成静态文件
	generateStaticHtml(p.Ctx, contentTpl, fileName, product)
}

// 生成html静态文件
func generateStaticHtml(ctx iris.Context, template *template.Template, filename string, product *datamodels.Product) {
	if exist(filename) {
		err := os.Remove(filename)
		if err != nil {
			ctx.Application().Logger().Debug(err)
		}
	}
	// 2. 生成静态文件
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		ctx.Application().Logger().Debug(err)
	}
	defer file.Close()
	template.Execute(file, &product)
}

func exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func (p *ProductController) GetDetail() mvc.View {
	idString := p.Ctx.URLParam("productID")
	id, errID := strconv.ParseInt(idString, 10, 64)
	if errID != nil{
		p.Ctx.Application().Logger().Debug(errID)
	}
	product, err := p.ProductService.GetProductByID(id)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	return mvc.View{
		Layout: "shared/productLayout.html",
		Name:   "product/view.html",
		Data: iris.Map{
			"product": product,
		},
	}
}

func (p *ProductController) GetCheck() mvc.View {

	userIDString := p.Ctx.GetCookie("uid")
	id, errID := strconv.ParseInt(userIDString, 10, 64)
	if errID != nil{
		p.Ctx.Application().Logger().Debug(errID)
	}
	//userID, err :=  strconv.Atoi(userIDString)
	//if err != nil {
	//	p.Ctx.Application().Logger().Debug(err)
	//}
	orderMap, _ := p.OrderService.GetResult(id)

	return mvc.View{
		Layout: "shared/productLayout.html",
		Name:"/product/check.html",
		Data:iris.Map{
			"order": orderMap,
		},
	}
}

func (p *ProductController) GetError() mvc.View {

	return mvc.View{
		Layout: "shared/productLayout.html",
		Name:   "product/error.html",

	}
}

func (p *ProductController) GetOrder() mvc.View {
	productIDString := p.Ctx.URLParam("productID")
	userIDString := p.Ctx.GetCookie("uid")
	productID, err := strconv.ParseInt(productIDString, 10, 64)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	userID, err :=  strconv.ParseInt(userIDString, 10, 64)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}

	//实现redis
	conn,err := redis.Dial("tcp","127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect redis error :",err)
	}
	fmt.Println("connect redis success")
	defer conn.Close()

	var showMessage string
	res, err := redis.String(conn.Do("LPOP", "product"))
	if err != nil {
		fmt.Println("redis POP error:", err)
		showMessage = "抢购失败"
		return mvc.View{
			Layout: "shared/productLayout.html",
			Name:   "product/result.html",
			Data: iris.Map{
				//"orderID":     orderID,
				"showMessage": showMessage,
			},
		}

	} else {
		res_type := reflect.TypeOf(res)
		fmt.Printf("res type : %s \n", res_type)
		fmt.Printf("res  : %s \n", res)
		showMessage = "操作成功，请稍后查询抢购结果"
	}
	//showMessage := "操作成功，请稍后查询抢购结果"
	// 创建消息体
	message := datamodels.NewMessage(userID, productID)
	byteMessage, err := json.Marshal(message)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	err = p.RabbitMQ.PublishSimple(string(byteMessage))
	if err!=nil {
		p.Ctx.Application().Logger().Debug(err)
	}

	return mvc.View{
		Layout: "shared/productLayout.html",
		Name:   "product/result.html",
		Data: iris.Map{
			//"orderID":     orderID,
			"showMessage": showMessage,
		},
	}

	//return []byte("true")
	//product, err := p.ProductService.GetProductByID(int64(productID))
	//if err != nil {
	//	p.Ctx.Application().Logger().Debug(err)
	//}
	//var orderID int64
	//showMessage := "抢购失败！"
	//// 判断商品数量是否满足需求
	//if product.ProductNum > 0 {
	//	// 扣除商品数量
	//	product.ProductNum -= 1
	//	err := p.ProductService.UpdateProduct(product)
	//	if err != nil {
	//		p.Ctx.Application().Logger().Debug(err)
	//	}
	//	// 创建订单
	//	userID, err := strconv.Atoi(userIDString)
	//	order := &datamodels.Order{
	//		UserId:      int64(userID),
	//		ProductId:   int64(productID),
	//		OrderStatus: datamodels.OrderSuccess,
	//	}
	//	orderID, err = p.OrderService.InsertOrder(order)
	//	if err != nil {
	//		p.Ctx.Application().Logger().Debug(err)
	//	} else {
	//		showMessage = "抢购成功！"
	//	}
	//}
	//showMessage := "请稍后查询抢购结果"
	//return mvc.View{
	//	Layout: "shared/productLayout.html",
	//	Name:   "product/result.html",
	//	Data: iris.Map{
	//		//"orderID":     orderID,
	//		"showMessage": showMessage,
	//	},
	//}

}
