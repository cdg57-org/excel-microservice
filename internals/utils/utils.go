package utils

import (
	"fmt"
	"strconv"
)

func conv(i int) string {

	if i <= 0 {

		return ""

	}

	j := (i - 1) % 26

	l := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	s := conv((i - j) / 26)

	return s + string(l[j])
}

func GetAxis(row, col int) (string, error) {
	if col <= 0 {
		return "", fmt.Errorf("column value was invalid")
	}

	if row <= 0 {
		return "", fmt.Errorf("row value was invalid")

	}
	cstr := conv(col)
	rstr := strconv.Itoa(row)
	return cstr + rstr, nil
}
