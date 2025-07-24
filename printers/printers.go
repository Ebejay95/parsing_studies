package printers

import (
	"fmt"
	"os"
)

const (
    Reset  = "\033[0m"
    Red    = "\033[31m"
    Green  = "\033[32m"
    Yellow = "\033[33m"
    Blue   = "\033[34m"
    Purple = "\033[35m"
    Cyan   = "\033[36m"
    White  = "\033[37m"
    Bold   = "\033[1m"
)

func Log(message string) {
	fmt.Printf("%s%s%s\n", Yellow+Bold, message, Reset)
}

func Success(message string) {
	fmt.Printf("%s%s%s\n", Green+Bold, message, Reset)
}

func Error(message string) {
	fmt.Fprintf(os.Stderr, "%sError: %s%s\n", Red+Bold, message, Reset)
}
