package services

import (
	"SecKill_Product/datamodels"
	"SecKill_Product/repositories"
)

type IOrderService interface {
	GetOrderByID(int64)(*datamodels.Order, error)
	DeleteOrderByID(int64)bool
	UpdateOrder(*datamodels.Order)error
	InsertOrder(*datamodels.Order) (int64, error)
	GetAllOrder()([]*datamodels.Order, error)

	GetResult(int64)([]*datamodels.Order, error)//获取用户的抢购结果

	GetAllOrderInfo()(map[int]map[string]string, error)
	InsertOrderByMessage(message *datamodels.Message)(int64, error)
	InsertLog(log datamodels.Log)error
}

func NewOrderService(repository repositories.IOrderRepository)IOrderService{
	return &OrderService{repository}
}

type OrderService struct {
	OrderRepository repositories.IOrderRepository
}

func (o *OrderService) GetResult(userId int64) ([]*datamodels.Order, error) {
	return o.OrderRepository.SelectResult(userId)
}

func (o *OrderService) InsertLog(log datamodels.Log)(err error) {
	return o.OrderRepository.InsertLogger(log)
}

func (o *OrderService)GetOrderByID(orderID int64)(order *datamodels.Order, err error)  {
	return o.OrderRepository.SelectByKey(orderID)
}

func (o *OrderService) DeleteOrderByID(orderID int64) bool{
	return o.OrderRepository.Delete(orderID)
}

func (o *OrderService) UpdateOrder (order *datamodels.Order) (err error) {
	return o.OrderRepository.Update(order)
}

func (o *OrderService) InsertOrder(order *datamodels.Order)(orderID int64, err error){
	return o.OrderRepository.Insert(order)
}
func (o *OrderService) GetAllOrder()(orderArray []*datamodels.Order,err error){
	return o.OrderRepository.SelectAll()
}

func (o *OrderService) GetAllOrderInfo()(orderMap map[int]map[string]string, err error){
	return o.OrderRepository.SelectAllWithInfo()
}

func (o *OrderService) InsertOrderByMessage(message *datamodels.Message)(int64, error){
	order := &datamodels.Order{
		UserId:message.UserID,
		ProductId:message.ProductID,
		OrderStatus:datamodels.OrderSuccess,
	}
	return o.InsertOrder(order)
}