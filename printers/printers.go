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
	if (os.Getenv("DEBUG") == "1") {
		fmt.Printf("%s%s%s\n", Yellow+Bold, message, Reset)
	}
}

func Success(message string) {
	fmt.Printf("%s%s%s\n", Green+Bold, message, Reset)
}

func Error(message string) {
	fmt.Fprintf(os.Stderr, "%sError: %s%s\n", Red+Bold, message, Reset)
}

func Errorf(format string, args ...interface{}) {
	formattedMessage := fmt.Sprintf(format, args...)
	fmt.Fprintf(os.Stderr, "%sError: %s%s\n", Red+Bold, formattedMessage, Reset)
}

func NewError(message string) error {
	fmt.Fprintf(os.Stderr, "%sError: %s%s\n", Red+Bold, message, Reset)
	return fmt.Errorf(message)
}

func NewErrorf(format string, args ...interface{}) error {
	formattedMessage := fmt.Sprintf(format, args...)
	fmt.Fprintf(os.Stderr, "%sError: %s%s\n", Red+Bold, formattedMessage, Reset)
	return fmt.Errorf(formattedMessage)
}

func Logf(format string, args ...interface{}) {
	if (os.Getenv("DEBUG") == "1") {
		formattedMessage := fmt.Sprintf(format, args...)
		fmt.Printf("%s%s%s\n", Yellow+Bold, formattedMessage, Reset)
	}
}

func Successf(format string, args ...interface{}) {
	if (os.Getenv("DEBUG") == "1") {
		formattedMessage := fmt.Sprintf(format, args...)
		fmt.Printf("%s%s%s\n", Green+Bold, formattedMessage, Reset)
	}
}