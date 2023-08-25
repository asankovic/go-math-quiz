package main

import (
	"fmt"
	"math/rand"
	"time"
)

func runQuiz(tasks []quizTask, options quizOptions) {
	var (
		timer    *time.Timer
		timerCh  <-chan time.Time
		timedOut bool
	)

	correct := 0
	quizCh := make(chan bool)
	stopCh := make(chan bool)

	numQuestions := *options.quantity
	if numQuestions > len(tasks) {
		numQuestions = len(tasks)
	}

	selectedTasks := tasks[:numQuestions]
	if *options.shuffle {
		rand.Shuffle(len(selectedTasks), func(i, j int) { selectedTasks[i], selectedTasks[j] = selectedTasks[j], selectedTasks[i] })
	}

	timerMessage := "no time limit"
	if *options.timeLimit > 0 {
		timerMessage = fmt.Sprintf("%d seconds", *options.timeLimit)
	}

	fmt.Printf("Press any key when you are ready to start, you have %s to answer the questions. Good luck!", timerMessage)
	fmt.Scanln()
	fmt.Printf("Starting level %d quiz with %d questions!\n", *options.level, numQuestions)

	if *options.timeLimit > 0 {
		timer = time.NewTimer(time.Duration(*options.timeLimit) * time.Second)
		timerCh = timer.C
	}

quizLoop:
	for i, task := range selectedTasks {
		fmt.Println("===========================================")
		fmt.Printf("Question #%d: %s", i+1, task.question)

		go correctUserAnswer(task, quizCh, stopCh)

		select {
		case <-timerCh:
			fmt.Printf("\nWhoops, time's up, write your last words! 😆\n")
			stopCh <- true
			timedOut = true
			break quizLoop
		case validAnswer := <-quizCh:
			if validAnswer {
				correct++
			}
		}

	}

	if timedOut {
		<-quizCh
	}
	fmt.Println("===========================================")
	fmt.Printf("Quiz finished, your score was %d/%d!\n", correct, numQuestions)
}

func correctUserAnswer(task quizTask, quizCh chan<- bool, stopCh <-chan bool) {
	var userAnswer string
	fmt.Scan(&userAnswer)

	select {
	case <-stopCh:
		quizCh <- false
	default:
		if userAnswer == task.answer {
			fmt.Println("✅ Correct!")
			quizCh <- true
		} else {
			fmt.Printf("❌ Wrong, correct answer was %s!\n", task.answer)
			quizCh <- false
		}
	}

}
