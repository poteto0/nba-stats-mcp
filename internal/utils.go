package internal

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// ex "2023-24" -> 2023
func ParseSeasonBeginYear(seasonYear string) (int, error) {
	parts := strings.Split(seasonYear, "-")
	if len(parts) != 2 {
		return 0, errors.New("invalid season year format")
	}

	return strconv.Atoi(parts[0])
}

func FormatSeasonYear(year int) string {
	nextYear := (year + 1) % 100
	return fmt.Sprintf("%d-%02d", year, nextYear)
}
