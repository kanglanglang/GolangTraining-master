package main

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"time"

	//_ "github.com/mattn/go-sqlite3"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// 创建新的账户


func NewBCp(host_name string, cpu int32,ram int64, ram_percent float64, ip_addr string) error {
	// 对未存在记录进行插入
	_, err := x.Insert(&Bcp{HostName: host_name, Cpu: cpu, Ram: ram , RamPercent: ram_percent, CreateTime: time.Now(), IpAddr: ip_addr})
	return err
}
// 银行账户
func UpdateBcp(host_name string, cpu int32,ram int64, ram_percent float64, ip_addr string) error {

	info, err := x.In("host_name",host_name).Update(&Bcp{HostName: host_name, Cpu: cpu, Ram: ram , RamPercent: ram_percent, LastUpdateTime: time.Now(), IpAddr: ip_addr})
	if info == 0 {
		if err := NewBCp(host_name, cpu, ram, ram_percent, ip_addr); err != nil {
			fmt.Println("Fail to create new update:", err)

		} else {

			fmt.Println("New account has been  update")
		}
	}

	fmt.Println(err)
	return err
}

type Bcp struct {
	Id      int64
	HostName    string `xorm:"unique"`
	Cpu     int32
	Ram     int64
	RamPercent float64
	LastUpdateTime time.Time
	CreateTime time.Time
	IpAddr string
}

var bcpinfo []Bcp

var isexit bool

func FindBcp()  bool {
	isexit = false
	isexit, err := x.IsTableExist(&Bcp{})
	if err != nil {
		log.Fatalf("Chenck : %v\n", err)
	}
	return isexit
}





// ORM 引擎

var x *xorm.Engine


func init() {
	// 创建 ORM 引擎与数据库
	var err error
	x, err = xorm.NewEngine("mysql", "root:2wsxCDE#@tcp(192.168.200.12:3306)/simpledbt?charset=utf8")
	if err != nil {
		log.Fatalf("Fail to create engine: %v\n", err)
	}
	// 同步结构体与数据表
	if err = x.Sync(new(Bcp)); err != nil {
		log.Fatalf("Fail to sync database: %v\n", err)
	}
}
