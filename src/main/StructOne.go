package main

import "fmt"

type Books struct {
	title string

	author string

	subject string

	book_id int
}

func main() {

	fmt.Println(Books{"GO 语言", "www.baidu.com", "Go 语言", 1234})

	fmt.Println("--------------------------------------------------------------------")

	fmt.Println(Books{title: "Go语言", author: "www.baidu.com", subject: "Go 语言教程", book_id: 123})

	fmt.Println("--------------------------------------------------------------------")

	fmt.Println(Books{title: "Go 语言", author: "www.google.com"})

}
