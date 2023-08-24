package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func saveTasks(tasks []quizTask) (file string) {
	csvFile, err := os.Create(fmt.Sprintf("generated_tasks_%d.csv", time.Now().UnixMilli()))
	defer csvFile.Close()
	checkErr(err, "Failed creating new CSV file")

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	writer.WriteAll(createCsvRows(tasks))

	filePath, err := filepath.Abs(csvFile.Name())
	checkFatalErr(err, "Cannot retrieve path for saved CSV")
	return filePath
}

func createCsvRows(tasks []quizTask) [][]string {
	data := make([][]string, len(tasks))
	for i, task := range tasks {
		data[i] = []string{task.question, task.answer}
	}
	return data
}

func readTasks(fileName string) []quizTask {
	csvFile, err := os.Open(fileName)
	checkFatalErr(err, fmt.Sprintf("Failed opening file '%s'", fileName))

	reader := csv.NewReader(csvFile)

	data, err := reader.ReadAll()
	checkFatalErr(err, fmt.Sprintf("Failed reading tasks from file '%s'", fileName))

	return createTasks(data)
}

func createTasks(data [][]string) []quizTask {
	tasks := make([]quizTask, len(data))
	for i, row := range data {
		tasks[i] = quizTask{
			question: row[0],
			answer:   row[1],
		}
	}
	return tasks
}
