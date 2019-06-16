package main

import (
	"fmt"
	"xusi-projects/xusi-framework/xdoc"
)

func main() {

	xdoc.Run()

	for _, value := range xdoc.Docs {
		fmt.Println("package ：", value.Name, value.Describe)
		fmt.Println("\t常量 ：", len(value.Const), "个")
		for _, cons := range value.Const {
			fmt.Println("\t\t", cons.Name, cons.Type, cons.Value, cons.Describe)
		}
		fmt.Println("\t结构体 ：", len(value.Struct), "个")
		for _, structer := range value.Struct {
			fmt.Println("\t\t", structer.Name)
			for _, attr := range structer.Attr {
				fmt.Println("\t\t\t", attr.Name, attr.Type, attr.Describe, attr.Addition)
			}
		}
		fmt.Println("\t函数 ：", len(value.Func), "个")
		for _, fun := range value.Func {
			fmt.Println("\t\t", fun.Name, fun.Describe, "参数：", len(fun.Params), "个")
			for _, param := range fun.Params {
				fmt.Println("\t\t\t", param.Name, param.Type, param.Describe)

			}
		}

		fmt.Println("")
	}
}
