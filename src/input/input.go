package input

import (
	"bufio"
	"os"
	"strconv"
)

func InputInt() (int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strconv.Atoi(scanner.Text())
}

func InputString() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
