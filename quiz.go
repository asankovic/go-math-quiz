package main

import (
	"fmt"
	"math/rand"
)

//TODO add timer
//TODO add regex check for custom csv questions

func runQuiz(tasks []quizTask, options quizOptions) {
	correct := 0

	numQuestions := *options.quantity
	if numQuestions > len(tasks) {
		numQuestions = len(tasks)
	}

	selectedTasks := tasks[:numQuestions]
	if *options.shuffle {
		rand.Shuffle(len(selectedTasks), func(i, j int) { selectedTasks[i], selectedTasks[j] = selectedTasks[j], selectedTasks[i] })
	}

	fmt.Printf("Starting level %d quiz with %d questions, good luck!\n", *options.level, numQuestions)
	for i, task := range selectedTasks {
		fmt.Println("===========================================")
		fmt.Printf("Question #%d: %s", i+1, task.question)
		if correctUserAnswer(task) {
			correct++
		}
	}
	fmt.Println("===========================================")
	fmt.Printf("Quiz finished, your score was %d/%d!\n", correct, numQuestions)
}

func correctUserAnswer(task quizTask) bool {
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
