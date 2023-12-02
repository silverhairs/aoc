package aoc2023

import (
	"bufio"
	"os"
	"strconv"
	"unicode"
)

func Trebuchet(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	defer file.Close()

	scnr := bufio.NewScanner(file)
	var code int
	for scnr.Scan() {
		line := scnr.Text()
		code += getCalibrationValue(line)
	}
	println(code)
}

func getCalibrationValue(line string) int {
	digits := ""
	for _, char := range line {
		if unicode.Is(unicode.Nd, char) {
			digits += string(char)
		}
	}
	if digits != "" {
		digits = string(digits[0]) + string(digits[len(digits)-1])
		if num, err := strconv.Atoi(digits); err == nil {
			return num
		}
	}
	return 0
}
