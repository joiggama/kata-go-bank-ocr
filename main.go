package main

import (
	"fmt"
	"github.com/joiggama/kata-go-bank-ocr/account_numbers"
)

func main() {
	numbers := account_numbers.Parse("numbers.txt")

	for _, number := range numbers {
		res := account_numbers.IsValid(number)
		fmt.Println(number, res)
	}
}
