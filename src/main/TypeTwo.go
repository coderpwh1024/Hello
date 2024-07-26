package main

import "fmt"

func main() {

	/*	str := "123"

		num, err := strconv.Atoi(str)

		if err != nil {
			fmt.Println("转换错误:", err)
		} else {
			fmt.Printf("字符串 '%s' 转换为整数为: %d\n", str, num)
		}
	*/

	/*	num := 3.14
		str := strconv.FormatFloat(num, 'f', 2, 64)

		fmt.Printf("浮点数 %f 转为字符串为: '%s'\n", num, str)*/

	var i interface{} = "Hello,World"

	str, ok := i.(string)
	if ok {
		fmt.Println("%s is  a string \n", str)
	} else {
		fmt.Println("conversion failed")
	}

}
