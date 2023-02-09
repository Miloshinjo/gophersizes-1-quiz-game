package game

import (
	"fmt"
	"strings"

	"github.com/Miloshinjo/gophersizes-1-quiz-game/pkg/problems"
	"github.com/Miloshinjo/gophersizes-1-quiz-game/pkg/userinput"
)

func parseAnswer(a string) string {
	t := strings.TrimSpace(a)

	t = strings.ReplaceAll(t, "\n", " ")

	t = strings.ToLower(t)

	return t
}

func New() {
	problems := problems.Load()
	score := 0

	for index, problem := range problems {
		answer := userinput.Get(fmt.Sprintf("Question %d: %s: ", index+1, problem.Q))

		parsedAnswer := parseAnswer(answer)

		isCorrect := parsedAnswer == problem.A

		if isCorrect {
			score = score + 1
		}
	}

	fmt.Println("Game over! Here is your score:")
	fmt.Printf("Congrats! %d out of %d answers are correct!", score, len(problems))
}
