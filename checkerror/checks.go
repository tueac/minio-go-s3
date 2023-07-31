package checkerror

import (
	"fmt"
	"os"
)

func CheckCodeError(runner error) {
	if runner != nil {
		fmt.Println("Error: ", runner)
		os.Exit(-1)
	}
}
