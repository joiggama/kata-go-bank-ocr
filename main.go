package main

import (
	"fmt"
	"github.com/joiggama/kata-go-bank-ocr/account_numbers"
)

func main() {
	numbers := account_numbers.Parse("user_story_3.txt")

	for _, number := range numbers {
		_, state := account_numbers.IsValid(number)
		fmt.Println(number, state)
	}
}
