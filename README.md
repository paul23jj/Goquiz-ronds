GoQuiz - README
📋 Objectif
GoQuiz est une application interactive de quiz en ligne de commande développée en Go. Elle permet aux utilisateurs de tester leurs connaissances sur trois domaines différents :

  - Informatique générale
  - Cybersécurité
  - Intelligence Artificielle et Data Science
L'application propose une expérience gamifiée avec un système de scoring et des niveaux de progression.
🏗️ Structure du Projet
```
GoQuiz/
├── main.go              # Point d'entrée de l'application
├── quizmenu.go          # Gestion du menu principal
├── quizinfo.go          # Quiz Informatique Générale
├── quizcyber.go         # Quiz Cybersécurité
├── quizIAdata.go        # Quiz IA & Data Science
├── go.mod               # Dépendances du projet Go
└── README.md            # Documentation (ce fichier)
```
🔄 Logique du Programme
Flux Global
```
Démarrage (main.go)
        ↓
    ShowMenu() (quizmenu.go)
        ↓
    Choix utilisateur (1, 2, 3, ou 0)
        ↓
    ┌───────────────────────────────┐
    │  StartQuizInfo()              │
    │  StartQuizCyber()             │  ← Exécution du quiz choisi
    │  StartQuizIAData()            │
    └───────────────────────────────┘
        ↓
    Fin / Retour au menu
```
📚 Flux Interne d'un Quiz (Exemple: quizinfo.go)
Structure Détaillée
```
StartQuizInfo()
    ├─ Initialisation
    │  ├─ totalQuestions = 10
    │  └─ score = 0
    │
    ├─ Boucle de 10 questions
    │  └─ Pour chaque question:
    │     ├─ Afficher la question (fmt.Println)
    │     ├─ Afficher 3 options numérotées (1, 2, 3)
    │     ├─ Appeler CheckAnswer(correctAnswer)
    │     │  ├─ Lire la réponse utilisateur (fmt.Scan)
    │     │  ├─ Comparer avec la réponse correcte
    │     │  ├─ Afficher ✅ Correct ! ou ❌ Mauvaise réponse
    │     │  └─ Retourner true/false
    │     │
    │     └─ Si réponse correcte → score++
    │
    └─ Affichage final
       └─ CalculateScore(score, totalQuestions)
          ├─ Afficher le score: "X/10 bonnes réponses"
          └─ Afficher le niveau:
             ├─ score ≤ 1  → 🔰 Apprenti codeur
             ├─ score = 2  → 💻 Développeur en progression
             └─ score ≥ 3  → 🤯 Cyber Mastermind
```
Exemple de Flux Détaillé (Question 1)
```
Question: "Quel est le système d'exploitation libre le plus utilisé ?"
Options affichées:
  1. Windows
  2. macOS
  3. Linux
  
Réponse correcte attendue: 3 (Linux)

Appel: CheckAnswer(3)
  └─ Prompt: "Ta réponse (numéro) : "
  └─ Utilisateur tape: 3
  └─ Vérification: userAnswer (3) == correctAnswer (3) ✓
  └─ Affichage: "✅ Correct !"
  └─ Retour: true
  
Résultat: score++ (score passe de 0 à 1)
```
🎮 Flux d'Utilisation Complet
```
1. Lancement: go run main.go
                ↓
2. Menu affiché:
   ==== GoQuiz ====
   1. Quiz Info
   2. Quiz Cyber
   3. Quiz IA & Data
   0. Quitter
   Choix: 
                ↓
3. Utilisateur choisit 1 (Quiz Info)
                ↓
4. 10 questions posées successivement
   (Chaque question attend une réponse)
                ↓
5. Score calculé et niveau attribué
                ↓
6. Retour au menu (implicite)
```
🚀 Utilisation
Lancer l'application
Naviguer dans l'application
1. Sélectionnez un quiz (1, 2, ou 3)
2. Répondez à chaque question en entrant le numéro de votre choix
3. Consultez votre score et votre niveau à la fin
4. Appuyez sur 0 pour quitter
📊 Système de Scoring
Score	Niveau
0-3	🔰 Apprenti codeur
3-7	💻 Développeur en progression
7-10	🤯 Cyber Mastermind
💡 Points Clés du Code
 - Pas de boucles explicitées : Chaque question est posée individuellement et séquentiellement
 - Fonctions réutilisables : CheckAnswer() et CalculateScore() sont partagées entre tous les quiz
 - Format cohérent : Chaque quiz suit le même pattern (initialisation → questions → scoring)
 - Interaction utilisateur : Via fmt.Scan() pour lire les inputs
