package main

import "fmt"

func main() {

	/*var a int = 20

	fmt.Printf("变量的地址：%x\n", &a)
	*/

	var a int = 20

	var ip *int

	ip = &a

	fmt.Printf("a 变量的地址是: %x\n", &a)

	fmt.Printf("ip 变量储存的指针地址:%x\n", ip)

	fmt.Printf("*ip 变量的值: %d\n", *ip)
}
