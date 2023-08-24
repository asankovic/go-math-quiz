package main

import "fmt"

//TODO add timer, shuffle and replay options

func runQuiz(tasks []quizTask, options quizOptions) {
	correct := 0
	numQuestions := *options.quantity
	if numQuestions > len(tasks) {
		numQuestions = len(tasks)
	}

	fmt.Printf("Starting level %d quiz with %d questions, good luck!\n", *options.level, numQuestions)
	for i, task := range tasks[:numQuestions] {
		fmt.Println("===========================================")
		fmt.Printf("Question #%d: %s", i+1, task.question)
		if evaluateUser(task) {
			correct++
		}
	}
	fmt.Println("===========================================")
	fmt.Printf("Quiz finished, your score was %d/%d!\n", correct, len(tasks))
}

func evaluateUser(task quizTask) bool {
	var userAnswer string
	fmt.Scan(&userAnswer)
	if userAnswer == task.answer {
		fmt.Println("✅ Correct!")
		return true
	} else {
		fmt.Printf("❌ Wrong, correct answer was %s!\n", task.answer)
		return false
	}
}
