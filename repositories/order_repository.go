package repositories

import (
	"SecKill_Product/common"
	"SecKill_Product/datamodels"
	"database/sql"
	"fmt"
	"strconv"
)

type IOrderRepository interface {
	Conn() error
	Insert(*datamodels.Order) (int64, error)
	Delete(int64) bool
	Update(*datamodels.Order) error
	SelectByKey(int64) (*datamodels.Order, error)
	SelectAll() ([]*datamodels.Order, error)
	SelectResult(int64) ([]*datamodels.Order, error)
	SelectAllWithInfo() (map[int]map[string]string, error)
	InsertLogger(log datamodels.Log)error
}

func NewOrderManagerRepository(table string, db *sql.DB) IOrderRepository {
	return &OrderManagerRepository{table: table, mysqlConn: db}
}

type OrderManagerRepository struct {
	table     string
	mysqlConn *sql.DB
}

func (o *OrderManagerRepository) Conn() error {
	if o.mysqlConn == nil {
		mysql, err := common.NewMysqlConn()
		if err != nil {
			return err
		}
		o.mysqlConn = mysql
	}
	if o.table == "" {
		o.table = "order"
	}
	return nil
}

//查看自己的抢购结果
func (o *OrderManagerRepository) SelectResult(userId int64) (orderArray []*datamodels.Order, err error) {
	//fmt.Printf("执行查询抢购结果")
	if errConn := o.Conn(); errConn != nil {
		//fmt.Println("连接错误")
		return nil, errConn
	}
	//fmt.Printf("1")
	sql := "select * from homework." + o.table +" where userID="+ strconv.FormatInt(userId, 10)
	rows, errRows := o.mysqlConn.Query(sql)
	if errRows != nil {
		//fmt.Printf("errRows错误")
		return nil, errRows
	}
	result := common.GetResultRows(rows)
	if len(result) == 0 {
		//fmt.Printf("查询为空")
		return nil, err
	}
	//fmt.Printf("查询不为空")
	for _, v := range result {
		order := &datamodels.Order{}
		common.DataToStructByTagSql(v, order)
		orderArray = append(orderArray, order)
	}
	return
}

//这里是插入日志的函数
func (o *OrderManagerRepository) InsertLogger(log datamodels.Log)(err error) {
	if err = o.Conn(); err != nil {
		return
	}
	sql := "INSERT homework." + "Log" + " set Name=?,IP=?,Operation=?,Time=?,ProductID=?"
	stmt, errStmt := o.mysqlConn.Prepare(sql)
	if errStmt != nil {
		fmt.Println("errstmt error")
		return
	}
	_, errResult := stmt.Exec(log.Name, log.IP, log.Operation,log.Time,log.ProductID)
	if errResult != nil {
		fmt.Println("errResult error")
		return
	}
	return
}

func (o *OrderManagerRepository) Insert(order *datamodels.Order) (orderID int64, err error) {
	if err = o.Conn(); err != nil {
		return
	}
	sql := "INSERT homework." + o.table + " set UserID=?,productID=?,orderStatus=?"
	stmt, errStmt := o.mysqlConn.Prepare(sql)
	if errStmt != nil {
		return orderID, errStmt
	}
	result, errResult := stmt.Exec(order.UserId, order.ProductId, order.OrderStatus)
	if errResult != nil {
		return orderID, errResult
	}
	return result.LastInsertId()
}

func (o *OrderManagerRepository) Delete(orderID int64) (isOK bool) {
	if err := o.Conn(); err != nil {
		return
	}
	sql := "delete from homework." + o.table + " where ID=?"
	stmt, errStmt := o.mysqlConn.Prepare(sql)
	if errStmt != nil {
		return
	}
	_, err := stmt.Exec(orderID)
	if err != nil {
		return
	}
	return true
}

func (o *OrderManagerRepository) Update(order *datamodels.Order) (err error) {
	if errConn := o.Conn(); errConn != nil {
		return errConn
	}
	sql := "UPDATE homework." + o.table + " set userID=?,productID=?,orderStatus=? WHERE ID=" + strconv.FormatInt(order.ID, 10)
	stmt, errStmt := o.mysqlConn.Prepare(sql)
	if errStmt != nil {
		return errStmt
	}
	_, err = stmt.Exec(order.UserId, order.ProductId, order.OrderStatus)
	return
}

func (o *OrderManagerRepository) SelectByKey(orderID int64) (order *datamodels.Order, err error) {
	if errConn := o.Conn(); errConn != nil {
		return &datamodels.Order{}, errConn
	}
	sql := "select * from homework." + o.table + " where ID=" + strconv.FormatInt(orderID, 10)
	row, errRow := o.mysqlConn.Query(sql)
	if errRow != nil {
		return &datamodels.Order{}, errRow
	}
	result := common.GetResultRow(row)
	if len(result) == 0 {
		return &datamodels.Order{}, err
	}
	order = &datamodels.Order{}
	common.DataToStructByTagSql(result, order)
	return
}

func (o *OrderManagerRepository) SelectAll() (orderArray []*datamodels.Order, err error) {
	if errConn := o.Conn(); errConn != nil {
		return nil, errConn
	}
	sql := "select * from " + o.table
	rows, errRows := o.mysqlConn.Query(sql)
	if errRows != nil {
		return nil, errRows
	}
	result := common.GetResultRows(rows)
	if len(result) == 0 {
		return nil, err
	}
	for _, v := range result {
		order := &datamodels.Order{}
		common.DataToStructByTagSql(v, order)
		orderArray = append(orderArray, order)
	}
	return
}

func (o *OrderManagerRepository) SelectAllWithInfo() (orderMap map[int]map[string]string, err error) {
	if errConn := o.Conn(); errConn != nil {
		return nil, errConn
	}
	sql := "select o.ID,u.userName,p.productName,o.orderStatus from homework.order as o left join product as p on o.productID=p.ID left join user as u on o.userID=u.ID;"
	rows, errRows := o.mysqlConn.Query(sql)
	if errRows != nil {
		return nil, errRows
	}
	return common.GetResultRows(rows), err
}
