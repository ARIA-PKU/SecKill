package repositories

import (
	"database/sql"
	"SecKill_Product/common"
)

// 第一步，先开发对应的接口
// 第二步，实现定义的接口

type ILogRepository interface {
	// 连接数据
	ConL() (error)
	SelectAllLog() (map[int]map[string]string, error)

}

func NewLogManager(table string, db *sql.DB) ILogRepository {
	return &LogManager{table, db}
}

type LogManager struct {
	table     string
	mysqlConn *sql.DB
}


// 数据库连接
func (p *LogManager) ConL() (err error) {
	if p.mysqlConn == nil {
		mysql, err := common.NewMysqlConn()
		if err != nil {
			return err
		}
		p.mysqlConn = mysql
	}
	if p.table == "" {
		p.table = "log"
	}
	return
}

func (p *LogManager) SelectAllLog() (orderMap map[int]map[string]string, err error) {
	if errConn := p.ConL(); errConn != nil {
		return nil, errConn
	}
	sql := "select Name,ProductID,IP,Operation,Time from homework.log ;"
	rows, errRows := p.mysqlConn.Query(sql)
	if errRows != nil {
		return nil, errRows
	}
	return common.GetResultRows(rows), err
}