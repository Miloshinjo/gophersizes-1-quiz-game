package userinput

import (
	"bufio"
	"fmt"
	"os"
)

func Get(prompt string) string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println(prompt)

	userinput, _ := reader.ReadString('\n')

	return userinput
}
