package game

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/Miloshinjo/gophersizes-1-quiz-game/pkg/problems"
	"github.com/Miloshinjo/gophersizes-1-quiz-game/pkg/userinput"
)

// parseAnswer parses the answer for spaces and newlines
func parseAnswer(a string) string {
	t := strings.TrimSpace(a)
	t = strings.ReplaceAll(t, "\n", " ")
	t = strings.ToLower(t)

	return t
}

// printResults prints the game feedback for the user when game ends
func printResult(score int, problemsLen int) {
	fmt.Printf("\nGame over! Here is your score:\n")
	fmt.Printf("Congrats! %d out of %d answers are correct!", score, problemsLen)
}

// prompt prints the problem to user for answering
func prompt(questionNumber int, question string) {
	fmt.Printf("Question %d: %s: ", questionNumber, question)
}

func checkAnswer(answer string, correctAnswer string) bool {
	return answer == correctAnswer
}

func New() {
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")

	problems := problems.Load()

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	score := 0

	for index, problem := range problems {
		// Prompt the question to user
		prompt(index+1, problem.Q)

		// Create answers channel to track
		// if user typed in something.
		answerCh := make(chan string)

		// Create a goroutine to listen for
		// user answers
		go func() {
			answer := userinput.Get()

			// If user typed in an answer, insert it into this channe;
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			printResult(score, len(problems))
			return
		case answer := <-answerCh:
			if checkAnswer(parseAnswer(answer), problem.A) {
				score++
			}
		}
	}

	printResult(score, len(problems))

}
