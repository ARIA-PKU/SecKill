package main

import (
	"fmt"
	"SecKill_Product/common"
	"SecKill_Product/rabbitmq"
	"SecKill_Product/repositories"
	"SecKill_Product/services"
)

func main() {
	db, err := common.NewMysqlConn()
	if err!=nil {
		fmt.Println(err)
	}
	// 创建product数据库操作实例
	productRepository := repositories.NewProductManager("product", db)
	productService := services.NewProductService(productRepository)
	orderRepository := repositories.NewOrderManagerRepository("order", db)
	orderService := services.NewOrderService(orderRepository)
	rabbitmqConsumerSimple := rabbitmq.NewRabbitMQSimple("seckillProduct")
	rabbitmqConsumerSimple.ConsumeSimple(orderService, productService)
}
