package main

import (
	"assignment2/pkg"
	"fmt"
)

var (
	base = pkg.BasePc{}
	home = pkg.HomePc{
		Cpu:          4,
		GraphicsCard: 1,
		Wrapper:      base,
	}
	server = pkg.ServerPc{
		Cpu:     16,
		Memory:  256,
		Wrapper: base,
	}
)

func main() {
	fmt.Printf("Base price:[%f]\n", base.GetPrice())
	fmt.Printf("Home Cpu:[%d] Cards:[%d] Price:[%f]\n", home.Cpu, home.GraphicsCard, home.GetPrice())
	fmt.Printf("Server Cpu:[%d] Memory:[%d] Price:[%f]\n", server.Cpu, server.Memory, server.GetPrice())
}
