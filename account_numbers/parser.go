package account_numbers

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var dictionary = make(map[int]string, 10)

func ReadFile(fileName string) string {
	fmt.Println("Reading", fileName)
	content, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error reading", fileName, err.Error())
	}

	return strings.Trim(string(content), "\n\n\n")
}

func ParseEntry(entry string) string {
	var digits string

	for _, digit := range GetDigits(entry, 27) {
    if value := ParseDigit(digit) ; value != "" {
      digits += value
    } else {
      digits += "?"
    }
	}

	return digits
}

func ParseDigit(digit string) string {
	var value string

	for number, representation := range dictionary {
		if representation == digit {
			value = strconv.Itoa(number)
		}
	}

	return value
}

func GetDigits(entry string, lineLength int) []string {
	var digits []string

	for index := 0; index < lineLength; index += 3 {
		digit := ""

		for _, line := range strings.Split(entry, "\n") {
			if line != "" {
				digit += line[index:index+3] + "\n"
			} else {
				digit += "   \n"
			}
		}
		digits = append(digits, digit)
	}
	return digits
}

func GetEntries(content string) []string {
	return regexp.MustCompile(`\s{27}\n`).Split(content, -1)
}

func LoadDictionary() {
	contents := ReadFile("digits.txt")
	entry := GetEntries(contents)[0]
	digits := GetDigits(entry, 30)
	for i, v := range digits {
		dictionary[i] = v
	}
}

func Parse(fileName string) (accountNumbers []string) {
	LoadDictionary()
	contents := ReadFile(fileName)
	entries := GetEntries(contents)
	fmt.Println("Found", len(entries), "entries in", fileName)

	for _, entry := range entries {
		accountNumbers = append(accountNumbers, ParseEntry(entry))
	}

	return accountNumbers
}
