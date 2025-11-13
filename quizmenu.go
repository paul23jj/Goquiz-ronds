package main

import "fmt"

func ShowMenu() {
	var choice int
	fmt.Println("==== Quiz Menu ====")
	fmt.Println("1. Quiz Info")
	fmt.Println("2. Quiz Cyber")
	fmt.Println("3. Quiz IA & Data")
	fmt.Println("0. Quitter\n")
	fmt.Print("Choix : ")
	fmt.Scan(&choice)

	if choice == 1 {
		StartQuizInfo()
	} else if choice == 2 {
		StartQuizCyber()
	} else if choice == 3 {
		StartQuizIAData()
	} else {
		fmt.Println("Merci d'avoir joué 👋")
	}
}