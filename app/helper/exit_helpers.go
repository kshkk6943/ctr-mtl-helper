package helper

import (
	"bufio"
	"fmt"
	"os"
)

func CloseDueToError(err error) int {
	fmt.Println(err.Error())
	fmt.Println("Closing the application due to an error...")
	fmt.Println("Press anything to close the application")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return 0
}

func CloseApplication() int {
	fmt.Println("Done.")
	fmt.Println("Press anything to close the application")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return 0
}
