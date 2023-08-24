package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

const (
	Addition int = iota
	Subtraction
	Multiplication
	Modulo
)

const questionTextTemplate = "%d%s%d="

func generateTasks(options quizOptions) []quizTask {
	tasks := make([]quizTask, *options.quantity)
	min := int(math.Pow(10, float64(*options.level-1)))
	max := int(math.Pow(10, float64(*options.level)))

	for i := range tasks {
		first := rand.Intn(max-min) + min
		second := rand.Intn(max-min) + min

		task := generateTask(first, second)

		tasks[i] = task
	}
	return tasks
}

func generateTask(first, second int) quizTask {
	var question, answer string
	switch operation := rand.Intn(4-0) + 0; operation {

	case Addition:
		question = fmt.Sprintf(questionTextTemplate, first, "+", second)
		answer = strconv.Itoa(first + second)
	case Subtraction:
		question = fmt.Sprintf(questionTextTemplate, first, "-", second)
		answer = strconv.Itoa(first - second)
	case Multiplication:
		question = fmt.Sprintf(questionTextTemplate, first, "*", second)
		answer = strconv.Itoa(first * second)
	case Modulo:
		question = fmt.Sprintf(questionTextTemplate, first, "%", second)
		answer = strconv.Itoa(first % second)
	default:
		panic(errors.New("Critical error generating tasks"))
	}

	return quizTask{
		question,
		answer,
	}
}
