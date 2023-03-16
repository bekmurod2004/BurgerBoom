package main

import "fmt"

func  WhoYu() {
	fmt.Print("Hello choise are you Client or Cooker : ")	
	var choser string

	fmt.Scanln(&choser)

	if choser == "Client" || choser == "client"  {
		Clienter()
		
	}else if choser == "Cooker" || choser == "cooker" {
		Oshpaz()
		
	}
}