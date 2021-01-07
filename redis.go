package main


import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)


func main()  {
	conn,err := redis.Dial("tcp","127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect redis error :",err)
		return
	}
	fmt.Println("connect redis success")
	defer conn.Close()
	//放入库存数量的产品
	productNum := 100
	for i:=0;i< productNum;i++{
		_, err = conn.Do("LPUSH", "product", "1")
		if err != nil {
			fmt.Println("redis mset error:", err)
	}
	}
}
