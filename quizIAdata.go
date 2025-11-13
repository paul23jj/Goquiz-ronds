package main

import "fmt"

func StartQuizIAData() {
	fmt.Println("=== Quiz IA et Data ===")
	fmt.Println("=== Quiz Intelligence Artificielle et Data Science ===")

	totalQuestions := 3
	score := 0

	// Question 1
	fmt.Println("Quel algorithme est couramment utilisé pour les tâches de classification en apprentissage automatique ?")
	fmt.Println("1. Régression linéaire")
	fmt.Println("2. Forêt aléatoire")
	fmt.Println("3. K-means")
	if CheckAnswer(2) {
		score++
	}

	// Question 2
	fmt.Println("Quel langage de programmation est le plus populaire pour l'analyse de données ?")
	fmt.Println("1. R")
	fmt.Println("2. Java")
	fmt.Println("3. C++")
	if CheckAnswer(1) {
		score++
	}
	// Question 3
	fmt.Println("Que signifie le terme 'Big Data' ?")
	fmt.Println("1. Données volumineuses et complexes")
	fmt.Println("2. Données de petite taille")
	fmt.Println("3. Données non structurées uniquement")
	if CheckAnswer(1) {
		score++
	}
	CalculateScore(score, totalQuestions)
}
