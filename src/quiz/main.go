package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	quiz "quiz/controller"
)

var inputFile string
var timeoutSeconds int
var shuffle bool

func init() {
	flag.StringVar(&inputFile, "in", "questions.csv", "file conaining questions and answers")
	flag.IntVar(&timeoutSeconds, "t", 30, "time limit in seconds")
	flag.BoolVar(&shuffle, "s", false, "shuffle questions")
	flag.Parse()
}

func main() {
	bytes, err := ioutil.ReadFile(inputFile)
	check(err)

	in := string(bytes)

	fmt.Printf("Press <Enter> to start. Time to finish: %d seconds\n", timeoutSeconds)
	fmt.Scanln()
	quiz.ClearScreen()
	quiz := quiz.NewQuiz(timeoutSeconds)
	quiz.Load(in, shuffle)

	for e := quiz.NextQuestion(); e != nil; e = quiz.NextQuestion() {
		resp := e.AskQuestion()
		isCorrect, isTimeout := e.CheckAnswer(resp)
		if isTimeout {
			break
		}
		if isCorrect {
			quiz.IncrementCorrect()
		}
	}

	fmt.Printf("%d questions, %d correct\n", quiz.NumQuestions(), quiz.NumCorrect())
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
