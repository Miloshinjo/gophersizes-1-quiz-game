package problems

import (
	"encoding/csv"
	"flag"
	"os"
)

type Problem struct {
	Q string
	A string
}

func Load() []Problem {
	// get CSV filename
	csvFilename := flag.String("csv", "problems.csv", "a csv file in format of question/answer")
	flag.Parse()

	// Load CSV file
	file, err := os.Open(*csvFilename)
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

	var problems []Problem

	for _, record := range records {
		question := Problem{
			Q: record[0],
			A: record[1],
		}

		problems = append(problems, question)
	}

	return problems
}
