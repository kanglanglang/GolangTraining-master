package main

import (
"fmt"
"os"
"gopkg.in/ini.v1"
)

var (
	DBPort string
	DBName string
	DBUser string
	DBPasswd string
	DBHost string
)

func main() {
	cfg, err := ini.Load("/etc/bcpconfig.cfg")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	// 典型读取操作，默认分区可以使用空字符串表示

	DBHost=cfg.Section("Database").Key("DBHost").String()
	DBPort=cfg.Section("Database").Key("DBPort").String()
	DBName=cfg.Section("Database").Key("DBName").String()
	DBUser=cfg.Section("Database").Key("DBUser").String()
	DBPasswd=cfg.Section("Database").Key("DBPasswd").String()

	fmt.Println(DBHost,DBPort,DBName,DBUser,DBPasswd)
}
