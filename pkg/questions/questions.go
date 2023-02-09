package questions

import (
	"encoding/csv"
	"os"
)

type Question struct {
	Q string
	A string
}

func Load() []Question {
	// Load CSV file
	file, err := os.Open("problems.csv")
	if err != nil {
		panic("could not open CSV file.")
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read all the records
	records, err := reader.ReadAll()
	if err != nil {
		panic("could not read records.")
	}

	var questions []Question

	for _, record := range records {
		question := Question{
			Q: record[0],
			A: record[1],
		}

		questions = append(questions, question)
	}

	return questions
}
