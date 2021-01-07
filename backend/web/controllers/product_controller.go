package controllers

import (
	"SecKill_Product/common"
	"SecKill_Product/datamodels"
	"SecKill_Product/services"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"strconv"
	"time"
)

type ProductController struct {
	Ctx iris.Context
	ProductService services.IProductService
}


func (p *ProductController)GetAll() mvc.View  {
	productArray, _ := p.ProductService.GetAllProduct()

	return mvc.View{
		Layout: "shared/layout.html",
		Name:"product/view.html",
		Data:iris.Map{
			"productArray": productArray,
		},
	}
}

// 修改商品
func (p *ProductController) PostUpdate ()  {
	product := &datamodels.Product{}
	p.Ctx.Request().ParseForm()
	dec := common.NewDecoder(&common.DecoderOptions{"homework", false, false})
	if err := dec.Decode(p.Ctx.Request().Form, product);err!= nil{
		p.Ctx.Application().Logger().Debug(err)
	}
	err := p.ProductService.UpdateProduct(product)
	if err!= nil{
		p.Ctx.Application().Logger().Debug(err)
	}

	//IP
	localIp, _ := common.GetIntranceIp()

	token := p.Ctx.GetCookie("token")

	_, claims, _ := common.ParseToken(token)

	// 24小时制
	timeObj := time.Now()
	var NowTime = timeObj.Format("2006/01/02 15:04:05")

	log := datamodels.Log{
		Name :claims.Name,
		ProductID:product.ID,
		IP :localIp,
		Operation: "change product",
		Time:NowTime,
	}
	_ = p.ProductService.InsertLogProduct(log)
	p.Ctx.Application().Logger().Debug("ip地址为"+localIp+"修改产品信息成功")

	p.Ctx.Redirect("/product/all")
}

// 添加商品Get
func (p *ProductController)GetAdd()mvc.View  {
	return mvc.View{
		Name:"product/add.html",
	}
}

// 添加商品Post
func (p *ProductController)PostAdd()  {
	product := &datamodels.Product{}
	p.Ctx.Request().ParseForm()
	dec := common.NewDecoder(&common.DecoderOptions{"homework", false, false})
	if err := dec.Decode(p.Ctx.Request().Form, product);err!= nil{
		p.Ctx.Application().Logger().Debug(err)
	}
	_, err := p.ProductService.InsertProduct(product)
	if err!= nil{
		p.Ctx.Application().Logger().Debug(err)
	}
	//IP
	localIp, _ := common.GetIntranceIp()

	token := p.Ctx.GetCookie("token")

	_, claims, _ := common.ParseToken(token)

	// 24小时制
	timeObj := time.Now()
	var NowTime = timeObj.Format("2006/01/02 15:04:05")

	log := datamodels.Log{
		Name :claims.Name,
		ProductID:product.ID,
		IP :localIp,
		Operation: "add product",
		Time:NowTime,
	}
	_ = p.ProductService.InsertLogProduct(log)
	p.Ctx.Application().Logger().Debug("ip地址为"+localIp+"添加产品信息成功")

	p.Ctx.Redirect("/product/all")
}

// 根据ID查询商品
func (p *ProductController)GetManager() mvc.View  {
	idString := p.Ctx.URLParam("id")
	id, err := strconv.ParseInt(idString, 10, 16)
	if err != nil{
		p.Ctx.Application().Logger().Debug(err)
	}
	product, err := p.ProductService.GetProductByID(id)
	if err != nil{
		p.Ctx.Application().Logger().Debug(err)
	}

	//IP
	localIp, _ := common.GetIntranceIp()

	token := p.Ctx.GetCookie("token")

	_, claims, _ := common.ParseToken(token)

	// 24小时制
	timeObj := time.Now()
	var NowTime = timeObj.Format("2006/01/02 15:04:05")

	log := datamodels.Log{
		Name :claims.Name,
		ProductID:product.ID,
		IP :localIp,
		Operation: "check product",
		Time:NowTime,
	}
	_ = p.ProductService.InsertLogProduct(log)
	p.Ctx.Application().Logger().Debug("ip地址为"+localIp+"查看产品信息成功")

	return mvc.View{
		Name:"product/manager.html",
		Data:iris.Map{
			"product": product,
		},
	}
}

// 删除商品
func (p *ProductController)GetDelete()  {
	idString := p.Ctx.URLParam("id")
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	isOk := p.ProductService.DeleteProductByID(id)
	if isOk {
		p.Ctx.Application().Logger().Debug("删除商品成功，ID为"+idString)

		//IP
		localIp, _ := common.GetIntranceIp()

		token := p.Ctx.GetCookie("token")

		_, claims, _ := common.ParseToken(token)

		// 24小时制
		timeObj := time.Now()
		var NowTime = timeObj.Format("2006/01/02 15:04:05")

		log := datamodels.Log{
			Name :claims.Name,
			ProductID:id,
			IP :localIp,
			Operation: "delete product",
			Time:NowTime,
		}
		_ = p.ProductService.InsertLogProduct(log)
		p.Ctx.Application().Logger().Debug("ip地址为"+localIp+"删除产品信息成功")

	}else {
		p.Ctx.Application().Logger().Debug("删除商品失败，ID为"+idString)
	}

	p.Ctx.Redirect("/product/all")
}