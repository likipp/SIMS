package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func StringConvUint(UuidS string) (UuidU uint64) {
	UuidI, _ := strconv.Atoi(UuidS)
	UuidU = uint64(UuidI)
	return
}

func StringConvInt(str string) (int int) {
	int, _ = strconv.Atoi(str)
	return
}

func StringConvJoin(f, l string) (s string) {
	s = strings.Join([]string{f, l}, "/")
	return s
}

func IntConvJoin(s string) string {
	var newNumStr string
	lastNumber := StringConvInt(s)
	newNumber := lastNumber + 1
	if newNumber < 10 {
		newNumStr = fmt.Sprintf("%s%d", "0", newNumber)
		return newNumStr
	}
	newNumStr = strconv.Itoa(newNumber)
	return newNumStr
}
