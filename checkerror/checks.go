package checkerror

import (
	"fmt"
	"os"
)

func CheckCodeError(codeerror error) {
	if codeerror != nil {
		fmt.Println("Error: ", codeerror)
		os.Exit(-1)
	}
}
