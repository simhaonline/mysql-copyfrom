package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"time"
)

var db = &sql.DB{}

func init(){
	db,_ = sql.Open("mysql", "root:root@/book")
}

func main() {
	insert("")
}



func insert(sql string) {

	var start = time.Now()
	//Begin函数内部会去获取连接
	tx,_ := db.Begin()
	//每次循环用的都是tx内部的连接，没有新建连接，效率高
	tx.Exec(sql)
	//最后释放tx内部的连接
	tx.Commit()
	var end = time.Now()
	fmt.Println("insert total time:",end.Sub(start).Seconds())


	//方式1 insert total time: 3.7952171
	//方式2 insert total time: 4.3162468
	//方式3 insert total time: 4.3392482
	//方式4 insert total time: 0.3970227
	//方式5 insert total time: 7.3894226
}