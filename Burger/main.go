package main

import (
	"fmt"

	"github.com/TwiN/go-color"
)



func main() {
	Logo()
	WhoYu()
	
}













func Logo()  {

	s := ":"

	for i := 0; i < 70; i++ {

		if i == 35 {
			fmt.Println()
			fmt.Println(color.Ize(color.Yellow, ":::   Welocme to Burger King    :::"))
		}
		print(color.Ize(color.Red, s))
	}
fmt.Println()
	
}
