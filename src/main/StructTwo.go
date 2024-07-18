package main

import "fmt"

type Book struct {
	title string

	author string

	subject string

	book_id int
}

func main() {

	var Book1 Book

	var Book2 Book

	Book1.title = "Go语言"
	Book1.author = "www.google.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 123456

	Book2.title = "Java 语言"
	Book2.author = "www.google.com"
	Book2.subject = "Java 语言教程"
	Book2.book_id = 23456

	/*	fmt.Printf("Book 1 title %s\n", Book1.title)
		fmt.Printf("Book 1 author %s\n", Book1.author)
		fmt.Printf("Book 1 subject %s\n", Book1.subject)
		fmt.Printf("Book 1 book_id %d\n", Book1.book_id)

		fmt.Println("----------------------------------------------------------------")

		fmt.Printf("Book 2 title %s\n", Book2.title)
		fmt.Printf("Book 2 author %s\n", Book2.author)
		fmt.Printf("Book 2 subject %s\n", Book2.subject)
		fmt.Printf("Book 2 book_id %d\n", Book2.book_id)*/

	printBook(Book1)
	fmt.Println("----------------------------------------------------------------")
	printBook(Book2)

}

func printBook(book Book) {

	fmt.Printf("Book title: %s\n", book.title)
	fmt.Printf("Book author：%s\n", book.author)
	fmt.Printf("Book subject:%s\n", book.subject)
	fmt.Printf("Booke book_id:%d\n", book.book_id)

}
