package main

import (
	"flag"
	"fmt"
)

func main()  {
	//var name string
	//flag.StringVar(&name, "name", "everyone", "The greeting object.")

	//var name = *flag.String("name", "everyone", "The greeting object.")

	var name = getTheFlag();
	flag.Parse()
	fmt.Printf("Hello, %v!\n", name)
}

func getTheFlag() *string  {
	return flag.String("name", "everyone", "The greeting object.");
}
