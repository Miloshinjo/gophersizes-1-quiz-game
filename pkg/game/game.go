package game

import (
	"1-quiz-game/pkg/questions"
	"1-quiz-game/pkg/userinput"
	"fmt"
	"strings"
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
	score := Score{
		Correct:   0,
		Incorrect: 0,
	}

	for index, question := range questions {
		answer := userinput.Get(fmt.Sprintf("Question %d: %s: ", index+1, question.Q))

		parsedAnswer := parseAnswer(answer)

		isCorrect := parsedAnswer == question.A

		if isCorrect {
			score.Correct = score.Correct + 1
		} else {
			score.Incorrect = score.Incorrect + 1
		}
	}

	fmt.Println("Game over! Here is your score:")
	fmt.Println("Correct answers: ", score.Correct)
	fmt.Println("Incorrect answers: ", score.Incorrect)
}
