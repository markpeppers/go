package quiz

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Quizzer interface {
	NumQuestions() int
	NumCorrect() int
	Load(string, bool)
	NextQuestion() *Query
	IncrementCorrect()
}

type Query interface {
	AskQuestion() string
	CheckAnswer(string) (bool, bool)
}

type Entry struct {
	question string
	answer   string
	timeout  chan bool
}

type Quiz struct {
	Questions       []Entry
	currentQuestion int
	numCorrect      int
	timeout         chan bool
}

func (q Quiz) NumQuestions() int {
	return len(q.Questions)
}

func (q Quiz) NumCorrect() int {
	return q.numCorrect
}

func (q *Quiz) Load(questions string, shuffle bool) {
	r := csv.NewReader(strings.NewReader(questions))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		check(err)

		q.Questions = append(q.Questions, Entry{
			question: record[0],
			answer:   record[1],
		})
	}
	if shuffle {
		rand.Shuffle(q.NumQuestions(),
			func(i, j int) {
				q.Questions[i], q.Questions[j] = q.Questions[j], q.Questions[i]
			})
	}
}

func (q *Quiz) IncrementCorrect() {
	q.numCorrect++
}

func (q *Quiz) NextQuestion() *Entry {
	if q.currentQuestion >= q.NumQuestions() {
		return nil
	}
	entry := &q.Questions[q.currentQuestion]
	entry.timeout = q.timeout
	q.currentQuestion++
	return entry
}

func NewQuiz(timeoutSeconds int) Quiz {
	timeoutch := make(chan bool)
	go func() {
		time.Sleep(time.Duration(timeoutSeconds) * time.Second)
		timeoutch <- true
	}()
	return Quiz{timeout: timeoutch,
		currentQuestion: 0,
	}
}

func (e Entry) AskQuestion() string {
	fmt.Printf("%s?\n", e.question)
	ch := make(chan string)
	go queryTask(ch)
	select {
	case ans := <-ch:
		ClearScreen()
		close(ch)
		return ans
	case <-e.timeout:
		fmt.Println("TIME")
		return "__TIME__"
	}
}

func queryTask(ch chan<- string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	ch <- scanner.Text()
}

func (e Entry) CheckAnswer(resp string) (bool, bool) {
	if resp == "__TIME__" {
		return false, true
	}
	resp = strings.ToLower(strings.Trim(resp, " "))
	if resp == strings.ToLower(strings.Trim(e.answer, " ")) {
		fmt.Println("Correct")
		return true, false
	}
	fmt.Println("WRONG!!")
	return false, false
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
	return
}
