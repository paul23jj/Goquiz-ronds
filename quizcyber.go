package main

import "fmt"

func StartQuizCyber() {
	fmt.Println("=== Quiz Cyber Sécurité ===")

	totalQuestions := 10
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
	// Question 4
	fmt.Println("Qu'est-ce que le phishing ?")
	fmt.Println("1. Une technique de pêche en mer")
	fmt.Println("2. Une attaque visant à obtenir des informations sensibles")
	fmt.Println("3. Un type de logiciel antivirus")
	if CheckAnswer(2) {
		score++
	}
	// Question 5
	fmt.Println("Quel est le rôle d'un antivirus ?")
	fmt.Println("1. Protéger contre les logiciels malveillants")
	fmt.Println("2. Analyser les performances du système")
	fmt.Println("3. Gérer les connexions réseau")
	if CheckAnswer(1) {
		score++
	}
	// Question 6
	fmt.Println("Qu'est-ce qu'un VPN ?")
	fmt.Println("1. Un type de virus")
	fmt.Println("2. Un réseau privé virtuel")
	fmt.Println("3. Un protocole de transfert de fichiers")
	if CheckAnswer(2) {
		score++
	}
	// Question 7
	fmt.Println("Quelle est la principale fonction de la cryptographie ?")
	fmt.Println("1. Accélérer les connexions Internet")
	fmt.Println("2. Protéger les informations en les rendant illisibles sans clé")
	fmt.Println("3. Gérer les adresses IP")
	if CheckAnswer(2) {
		score++
	}
	// Question 8
	fmt.Println("Qu'est-ce qu'une attaque par déni de service (DdoS) ?")
	fmt.Println("1. Une attaque visant à rendre un service indisponible")
	fmt.Println("2. Une méthode de cryptage des données")
	fmt.Println("3. Un type de pare-feu")
	if CheckAnswer(1) {
		score++
	}
	// Question 9
	fmt.Println("Quel est le but principal des mises à jour de sécurité ?")
	fmt.Println("1. Ajouter de nouvelles fonctionnalités")
	fmt.Println("2. Améliorer les performances graphiques")
	fmt.Println("3. Corriger les vulnérabilités et améliorer la sécurité")
	if CheckAnswer(3) {
		score++
	}
	// Question 10
	fmt.Println("Quel est le rôle principal de l'authentification à deux facteurs (2FA) ?")
	fmt.Println("1. Augmenter la vitesse de connexion")
	fmt.Println("2. Ajouter une couche de sécurité supplémentaire")
	fmt.Println("3. Simplifier la gestion des mots de passe")
	if CheckAnswer(2) {
		score++
	}
	CalculateScore(score, totalQuestions)
}
