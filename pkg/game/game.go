package game

import (
	"fmt"
	"strings"

	"github.com/Miloshinjo/gophersizes-1-quiz-game/pkg/questions"
	"github.com/Miloshinjo/gophersizes-1-quiz-game/pkg/userinput"
)

func parseAnswer(a string) string {
	t := strings.TrimSpace(a)

	t = strings.ReplaceAll(t, "\n", " ")

	t = strings.ToLower(t)

	return t
}

type Score struct {
	Correct   uint
	Incorrect uint
}

func New() {
	questions := questions.Load()
	score := 0

	for index, question := range questions {
		answer := userinput.Get(fmt.Sprintf("Question %d: %s: ", index+1, question.Q))

		parsedAnswer := parseAnswer(answer)

		isCorrect := parsedAnswer == question.A

		if isCorrect {
			score = score + 1
		}
	}

	fmt.Println("Game over! Here is your score:")
	fmt.Printf("Congrats! %d out of %d answers are correct!", score, len(questions))
}
