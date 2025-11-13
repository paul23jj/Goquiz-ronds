package main

import "fmt"

func StartQuizCyber() {
	fmt.Println("=== Quiz Cyber Sécurité ===")
	
	totalQuestions := 3
	score := 0
	// Question 1
	fmt.Println("Quel protocole est utilisé pour sécuriser les communications sur Internet ?")
	fmt.Println("1. HTTP")
	fmt.Println("2. FTP")
	fmt.Println("3. HTTPS")
	if CheckAnswer(3) {
		score++
	}
	// Question 2
	fmt.Println("Qu'est-ce qu'un pare-feu (firewall) ?")
	fmt.Println("1. Un logiciel de traitement de texte")
	fmt.Println("2. Un dispositif de sécurité réseau")
	fmt.Println("3. Un type de virus informatique")
	if CheckAnswer(2) {
		score++
	}
	// Question 3
	fmt.Println("Quelle est la meilleure pratique pour créer un mot de passe sécurisé ?")
	fmt.Println("1. Utiliser des mots courants")
	fmt.Println("2. Utiliser une combinaison de lettres, chiffres et symboles")
	fmt.Println("3. Utiliser son nom ou sa date de naissance")
	if CheckAnswer(2) {
		score++
	}
	CalculateScore(score, totalQuestions)
}


