package main

import "fmt"

func StartQuizIAData() {
	fmt.Println("=== Quiz IA et Data ===")
	fmt.Println("=== Quiz Intelligence Artificielle et Data Science ===")

	totalQuestions := 10
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
	// Question 4
	fmt.Println("Que signifie 'Machine Learning' ?")
	fmt.Println("1. Apprentissage manuel")
	fmt.Println("2. Apprentissage automatique")
	fmt.Println("3. Apprentissage supervisé uniquement")
	if CheckAnswer(2) {
		score++
	}
	// Question 5
	fmt.Println("Que signiefie SQL ?")
	fmt.Println("1. Structured Query Language")
	fmt.Println("2. Simple Query Language")
	fmt.Println("3. Sequential Query Language")
	if CheckAnswer(1) {
		score++
	}
	// Question 6
	fmt.Println("Que signifie le réseau de neurones ?")
	fmt.Println("1. Un protocole de communication")
	fmt.Println("2. Un type de réseau informatique")
	fmt.Println("3. Un modèle d'apprentissage inspiré du cerveau humain")
	if CheckAnswer(3) {
		score++
	}
	// Question 7
	fmt.Println("Quelle était la tâche principale de deepBlue d'IBM ?")
	fmt.Println("1. Jouer aux échecs")
	fmt.Println("2. Analyser des données")
	fmt.Println("3. Reconnaissance vocale")
	if CheckAnswer(1) {
		score++
	}
	// Question 8
	fmt.Println("Quelle est la différence entre IA faible et IA forte ?")
	fmt.Println("1. IA faible est spécialisée, IA forte est générale")
	fmt.Println("2. IA faible est plus intelligente que IA forte")
	fmt.Println("3. IA faible utilise plus de données que IA forte")
	if CheckAnswer(1) {
		score++
	}
	// Question 9
	fmt.Println("Qui est considéré comme le père de l'intelligence artificielle ?")
	fmt.Println("1. Alan Turing")
	fmt.Println("2. John McCarthy")
	fmt.Println("3. Marvin Minsky")
	if CheckAnswer(2) {
		score++
	}
	// Question 10
	fmt.Println("Quel est l'objectif principal de l'analyse prédictive ?")
	fmt.Println("1. Prévoir les tendances futures")
	fmt.Println("2. Nettoyer les données")
	fmt.Println("3. Stocker les données efficacement")
	if CheckAnswer(1) {
		score++
	}
	CalculateScore(score, totalQuestions)
}
