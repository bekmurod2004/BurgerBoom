package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/TwiN/go-color"
)

type Kegan struct {
	Id int
	Foods []string
	Summ int
	Active int
}



func Oshpaz() {


	s := ":"

	for i := 0; i < 70; i++ {

		if i == 35 {
			fmt.Println()
			fmt.Println(color.Ize(color.Red, ":::        Welcome Cooker       :::"))
		}
		print(color.Ize(color.Blue, s))
	}
fmt.Println()





	data, err := ioutil.ReadFile("order.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	var zakazlar []Kegan
	err = json.Unmarshal(data, &zakazlar)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	if len(zakazlar) == 0 {
		fmt.Println("Zakaz hali kemagam")
		return
	}


	for _, v := range zakazlar {
		
		fmt.Println("Ochered : ",v.Id)

		fmt.Print("Ovqatlar : ")
		for i := 0; i < len(v.Foods); i++ {
			fmt.Print(v.Foods[i], " ")
		}
		fmt.Println()
		


		fmt.Println("Obshiy narh :",v.Summ)

		if v.Active == 0  {
			fmt.Println("qibolinmagan")
		}else {
			fmt.Println("qibolingan")
		}


		fmt.Println()
		
	}





	

	
}


