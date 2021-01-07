package controllers

import (
	"SecKill_Product/common"
	"SecKill_Product/datamodels"
	"SecKill_Product/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iris-contrib/jade/testdata/imp/model"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

//注册处理
func Register(ctx *gin.Context) {
	DB := common.GetDB()
	//使用map获取请求参数
	var requestUser = datamodels.Admin{}
	_ = ctx.Bind(&requestUser)

	////获取参数
		name := ctx.PostForm("userName")
		telephone := ctx.PostForm("telephone")
		password  := ctx.PostForm("passWord")
	    Authority := ctx.PostForm("Authority")

	//创建用户
	hasePassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("加密失败")
		return
	}

	newUser := datamodels.Admin{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasePassword),
		Authority: Authority,
	}
	DB.Create(&newUser)
	//返回结果
	//发放token
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "系统异常")
		log.Printf("token generate error: %v", err)
		return
	}
	//返回结果
	//response.Success(ctx, gin.H{"token": token}, "注册成功")
	fmt.Printf(token)
	ctx.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:8081/login")
}

func Login(ctx *gin.Context) {
	DB := common.GetDB()
	//获取数据
	//使用map获取请求参数
	var requestUser = model.User{}
	ctx.Bind(&requestUser)

	//获取参数
	telephone := ctx.PostForm("telephone")
	password  := ctx.PostForm("passWord")
	//数据验证
	fmt.Println(telephone, "手机号码长度", len(telephone))
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}

	//判断手机号是否存在
	var user datamodels.Admin
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 400, nil, "用户不存在")
		return
	}

	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}

	//发放token
	token, err := common.ReleaseToken(user)
	fmt.Printf(user.Name)
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "系统异常")
		log.Printf("token generate error: %v", err)
		return
	}

	//将token写入本地cookie中，有效期2小时
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(7200 * time.Second),
	})
	//这里把用户信息写入上下文
	ctx.Set("user",user)



	if user.Authority=="1" {
		ctx.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:8080/product")
	}

	if user.Authority=="0"{
		ctx.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:8080/order/allorder")
	}


	//返回结果
	//response.Success(ctx, gin.H{"token": token}, "登录成功")

}

//func Info(ctx *gin.Context) {
//	user, _ := ctx.Get("user")
//	ctx.JSON(http.StatusOK, gin.H{
//		"code": 200,
//		"data": gin.H{"user": dto.ToUserDto(user.(datamodels.Admin))},
//	})
//}

//判断是否已经存在此电话注册的用户
//func isTelephoneExists(db *gorm.DB, telephone string) bool {
//	var user model.User
//	db.Where("telephone = ?", telephone).First(&user)
//	if user.ID != 0 {
//		return true
//	}
//	return false
//}