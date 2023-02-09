package userinput

import (
	"bufio"
	"os"
)

func Get() string {
	reader := bufio.NewReader(os.Stdin)

	userinput, _ := reader.ReadString('\n')

	return userinput
}
