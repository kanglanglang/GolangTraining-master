package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"

)

func main() {

	v, _ := mem.VirtualMemory()
	fmt.Println(v)
	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", (v.Total / 1024 / 1024 / 1024), v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	fmt.Println(v)
	c, _ := cpu.Info()
	//p, _ := cpu.Percent(time.Millisecond, true)
	fmt.Println(c)
	//fmt.Println(p)

	println(v)

	}
