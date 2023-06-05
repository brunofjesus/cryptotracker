package mapper

import (
	"fmt"
	"strconv"
)

func sumStringAndFloat(n1 float64, n2 string) float64 {
	converted, err := strconv.ParseFloat(n2, 64)
	if err != nil {
		fmt.Printf("Cannot convert %s to float64: %v", n2, err)
		return 0
	}

	return n1 + converted
}

func subStrings(s1 string, s2 string) string {
	f1, f2, err := stringsToFloat(s1, s2)
	if err != nil {
		return "ERR"
	}
	return float64ToString(f1 - f2)
}

func mulStrings(s1 string, s2 string) string {
	f1, f2, err := stringsToFloat(s1, s2)
	if err != nil {
		return "ERR"
	}

	return float64ToString(f1 * f2)
}

func divStrings(s1 string, s2 string) string {
	f1, f2, err := stringsToFloat(s1, s2)
	if err != nil {
		return "ERR"
	}

	return float64ToString(f1 / f2)
}

func stringsToFloat(s1 string, s2 string) (float64, float64, error) {
	f1, err := strconv.ParseFloat(s1, 64)
	if err != nil {
		fmt.Printf("Cannot convert %s to float64: %v", s1, err)
		return -1, -1, err
	}

	f2, err := strconv.ParseFloat(s1, 64)
	if err != nil {
		fmt.Printf("Cannot convert %s to float64: %v", s2, err)
		return f1, -1, err
	}

	return f1, f2, nil
}

func float64ToString(n float64) string {
	return fmt.Sprintf("%f", n)
}
