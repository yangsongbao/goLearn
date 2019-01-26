package main

import "fmt"

var block = "package"

func main()  {
	block := "function"

	{
		block := "inner"
		fmt.Printf("The book is %s. \n", block)
	}
	fmt.Printf("The book is %s. \n", block)
}
