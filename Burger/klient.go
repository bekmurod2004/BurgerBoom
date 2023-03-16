package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/TwiN/go-color"
)

type BorOvqatlar struct {
	Id    int
	Name  string
	Price int
}

type Ketadigan struct {
	Id int
	Foods  []string
	Summ   int
	Active int
}

func Clienter() {


	s := ":"

	for i := 0; i < 70; i++ {

		if i == 35 {
			fmt.Println()
			fmt.Println(color.Ize(color.Green, ":::       Welcome Client        :::"))
		}
		print(color.Ize(color.Red, s))
	}
fmt.Println()












	date, hato := ioutil.ReadFile("foods.json")
	if hato != nil {
		fmt.Printf("::  %v ::", hato)
		return
	}

	var hammaOvq []BorOvqatlar
	hato = json.Unmarshal(date, &hammaOvq)
	if hato != nil {
		fmt.Printf("::  %v ::", hato)
		return
	}

	for i := 0; i < len(hammaOvq); i++ {
		fmt.Println(hammaOvq[i].Id, hammaOvq[i].Name, hammaOvq[i].Price)
	}

	var tanlash string

	fmt.Print(" zakaz || yoq : ")
	fmt.Scanln(&tanlash)

	act := true

	var ovqNomi string
	var ovqSoni int
	var yegindi []string

	for act {

		if tanlash == "zakaz" {

			fmt.Print("Ovqat nomini tanlang : ")
			fmt.Scanln(&ovqNomi)
			fmt.Print("Ovqat sonini tanlang : ")
			fmt.Scanln(&ovqSoni)

			var tekshirish = false

			for i := 0; i < len(hammaOvq); i++ {
				if hammaOvq[i].Name == ovqNomi {
					tekshirish = true

				}
			}

			if tekshirish {

				
				

				for i := 0; i < ovqSoni; i++ {

					yegindi = append(yegindi, ovqNomi)

				}

			}

			if ovqNomi == "yoq" {
				Summer(yegindi, hammaOvq)

				// fmt.Println(yegindi)
				act = false
			}

			// fmt.Println(yegindi)
			

		} else if tanlash == "yoq" {
			fmt.Println(yegindi)
			act = false

		}

	}

}

func Summer(yegindi []string, hammaOvq []BorOvqatlar) {
	var summ int = 0

	for i := 0; i < len(yegindi); i++ {
		for j := 0; j < len(hammaOvq); j++ {
			if yegindi[i] == hammaOvq[j].Name {
				summ += hammaOvq[j].Price

			}
		}

	}

	Tashab(summ, yegindi)

}

func Tashab(summ int, yegindi []string) {

	if summ == 0 {
		return
	}

	dat, err := ioutil.ReadFile("order.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	var borishL []Ketadigan
	err = json.Unmarshal(dat, &borishL)
	if err != nil {
		fmt.Println( err)
		return
	}

	idi := 0

	if len(borishL) == 0 {
		idi = 1
	}else if len(borishL) != 0 {
		idi = len(borishL) +1
		
	}

	lol := Ketadigan{
		Id: idi,
		Foods:  yegindi,
		Summ:   summ,
		Active: 0,
	}


	borishL = append(borishL, lol)

	dat, err = json.Marshal(borishL)
	if err != nil {
		fmt.Println( err)
		return
	}

	err = ioutil.WriteFile("order.json",dat, 0644)
	if err != nil {
		fmt.Println( err)
		return
	}


	// fmt.Println(lol)

}


