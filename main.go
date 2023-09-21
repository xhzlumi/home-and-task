package main

import (
	"fmt"
)

func main() {
	var phonenumber string
	var verification string
	for i := 1; i <= 5; i++ {
		fmt.Print("请输入手机号获取验证码:")
		fmt.Scanln(&phonenumber)
		if phonenumber == "16622225555" {

			fmt.Print("(注意：一分钟内无法再次获取验证码且验证码有效期为5分钟)")
			fmt.Print("请输入验证码:")
		} else {
			if phonenumber != "false" {
				fmt.Println("格式错误")
			}
			fmt.Scanln(&verification)

			if verification == "123456" {
				fmt.Println("登入成功")
				break
			} else {

				if verification != "false" {
					fmt.Println("输入错误")

				}
			}
		}
	}
}
