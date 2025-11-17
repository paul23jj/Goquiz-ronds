package main

import "fmt"

func StartQuizInfo() {
	fmt.Println("=== Quiz Informatique ===")
	fmt.Println("=== Quiz Informatique Générale ===")

	totalQuestions := 10
	score := 0
	// Question 1
	fmt.Println("Quel est le système d’exploitation libre le plus utilisé ?")
	fmt.Println("1. Windows")
	fmt.Println("2. macOS")
	fmt.Println("3. Linux")
	if CheckAnswer(3) {
		score++
	}
	// Question 2
	fmt.Println("Quel langage est compilé ?")
	fmt.Println("1. Python")
	fmt.Println("2. Go")
	fmt.Println("3. JavaScript")
	if CheckAnswer(2) {
		score++
	}
	// Question 3
	fmt.Println("Que signifie HTML ?")
	fmt.Println("1. HyperText Markup Language")
	fmt.Println("2. HighText Machine Learning")
	fmt.Println("3. Hyper Transfer Main Line")
	if CheckAnswer(1) {
		score++
	}
	//Question 4
	fmt.Println("Qu'est ce qu'une adresse IP ?")
	fmt.Println("1. Un identifiant unique pour un appareil sur un réseau")
	fmt.Println("2. Un protocole de communication")
	fmt.Println("3. Un type de logiciel malveillant")
	if CheckAnswer(1) {
		score++
	}
	//Question 5
	fmt.Println("Quel composant est considéré comme le 'cerveau' de l'ordinateur ?")
	fmt.Println("1. La RAM")
	fmt.Println("2. Le processeur (CPU)")
	fmt.Println("3. Le disque dur")
	if CheckAnswer(2) {
		score++
	}
	//Question 6
	fmt.Println("Que signifie le terme 'open source' ?")
	fmt.Println("1. Logiciel dont le code source est accessible et modifiable par tous")
	fmt.Println("2. Logiciel payant")
	fmt.Println("3. Logiciel développé uniquement par des entreprises")
	if CheckAnswer(1) {
		score++
	}
	//Question 7
	fmt.Println("Qu'est ce qu'un SSD ?")
	fmt.Println("1. Un type de mémoire vive")
	fmt.Println("2. Un type de disque de stockage rapide")
	fmt.Println("3. Un protocole réseau")
	if CheckAnswer(2) {
		score++
	}
	//Question 8
	fmt.Println("Qu'est ce qu'une API ?")
	fmt.Println("1. Une Interface de Programmation d'Applications")
	fmt.Println("2. Un type de virus informatique")
	fmt.Println("3. Un langage de programmation")
	if CheckAnswer(1) {
		score++
	}
	//Question 9
	fmt.Println("A quoi sert le BIOS dans un ordinateur ?")
	fmt.Println("1. Initialiser le matériel au démarrage")
	fmt.Println("2. Gérer les fichiers")
	fmt.Println("3. Protéger contre les virus")
	if CheckAnswer(1) {
		score++
	}
	//Question 10
	fmt.Println("Quel est le rôle de la RAM ?")
	fmt.Println("1. Stocker temporairement les données en cours d'utilisation")
	fmt.Println("2. Stocker les données de manière permanente")
	fmt.Println("3. Gérer les connexions réseau")
	if CheckAnswer(1) {
		score++
	}
	CalculateScore(score, totalQuestions)
}

func CheckAnswer(correctAnswer int) bool {
	var userAnswer int
	fmt.Print("Ta réponse (numéro) : ")
	fmt.Scan(&userAnswer)

	if userAnswer == correctAnswer {
		fmt.Println("✅ Correct !")
		return true
	} else {
		fmt.Println("❌ Mauvaise réponse.")
		return false
	}
}

func CalculateScore(score int, total int) {
	fmt.Printf("Tu as %d/%d bonnes réponses.\n", score, total)

	if score <= 3 {
		fmt.Println("🔰 Niveau : Apprenti codeur")
	} else if score > 3 && score <= 7 {
		fmt.Println("💻 Niveau : Développeur en progression")
	} else {
		fmt.Println("🤯 Niveau : Cyber Mastermind")
	}
}
