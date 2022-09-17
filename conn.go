package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type XzAutoServerConf struct {
	GroupZone  string `gorm:"column:group_zone"`
	ServerId   int
	OpenTime   string
	ServerName string
	Status     int
}

type ConnInfo struct {
	MyUser   string
	Password string
	Host     string
	Port     int
	Db       string
}

type Person struct {
	ID   int    `gorm:"primary_key"`
	Name string `gorm:"not_null"`
}

func add(db *gorm.DB) {

	user := &Person{Name: "zhangsan"}
	db.Create(user)

}

func DbConn(MyUser, Password, Host, Db string, Port int) *gorm.DB {

	connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", MyUser, Password, Host, Port, Db)
	db, err := gorm.Open("mysql", connArgs)
	if err != nil {
		log.Fatal(err)
	}

	db.SingularTable(true)
	return db
}

func main() {
	cn := ConnInfo{
		"root",
		"123456",
		"127.0.0.1",
		3306,
		"mysql",
	}
	db := DbConn(cn.MyUser, cn.Password, cn.Host, cn.Db, cn.Port)
	defer db.Close() // 关闭数据库链接，defer会在函数结束时关闭数据库连接
	var rows []XzAutoServerConf
	//select
	db.Where("status=?", 0).Select([]string{"group_zone", "server_id", "open_time", "server_name"}).Find(&rows)
	//update
	err := db.Model(&rows).Where("server_id=?", 80).Update("status", 1).Error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rows)

	//创建表
	db = db.CreateTable(Person{})
	//增加数据
	add(db)
	user := &Person{ID: 1}

	//更新数据
	db.Model(user).Update("Name", "lisi")

	//获取一条数据
	user1 := new(Person)
	db.First(user1, 1)
	fmt.Println(user1)

	//删除指定数据
	user2 := &Person{ID: 1}
	db.Delete(user2)
}
