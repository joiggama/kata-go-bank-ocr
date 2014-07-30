package account_numbers

import (
	"math"
	"strconv"
	"strings"
)

func IsValid(accountNumber string) (result bool) {
	return CheckSum(accountNumber) == 0
}

func CheckSum(accountNumber string) (res int) {
	d := [9]int{}

	for index, digit := range strings.Split(accountNumber, "") {
		value, _ := strconv.Atoi(digit)
		d[(len(d)-1)-index] = value
	}

	res = (d[0] + 2*d[1] + 3*d[2] + 4*d[3] + 5*d[4] + 6*d[5] + 7*d[6] + 8*d[7] + 9*d[8])

	return int(math.Mod(float64(res), 11))
}
