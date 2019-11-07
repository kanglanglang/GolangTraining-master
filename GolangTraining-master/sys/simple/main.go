
package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"net"
	"os"
	"time"
)

func main() {

	var ram int64
	var ram_percent float64
	var cpu_core int32
	var ip_addr string

	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				if ipnet.IP.String()[:9] == "192.168.2"{
					ip_addr = ipnet.IP.String()
					fmt.Println(ip_addr)
				}

			}

		}
	}
	t:=time.NewTicker(3*time.Second)

	for {
		select {
		case <-t.C:

			c, _ := cpu.Info()

			cpu_core = c[0].Cores

			v, _ := mem.VirtualMemory()
			ram_percent= v.UsedPercent

			ram = int64(v.Total / 1024 / 1024 / 1024)
			host_name, err := os.Hostname()
			if err != nil {
				fmt.Printf("get host%s", err)
			} else {
				fmt.Printf("get host %s", host_name)
			}

			if err := UpdateBcp(host_name, cpu_core, ram, ram_percent, ip_addr); err != nil {
				fmt.Println("Fail to create new update: v%\n", err)

			} else {

				fmt.Println("New account has been  update")
			}

		}
	}
}

