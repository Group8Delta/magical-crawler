package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func ParsePersianDate(phrase string) (time.Time, error) {
	now := time.Now()
	var t time.Time

	// Normalize Persian numerals to standard numerals
	numeralReplacements := map[string]string{
		"۰": "0", "۱": "1", "۲": "2", "۳": "3", "۴": "4",
		"۵": "5", "۶": "6", "۷": "7", "۸": "8", "۹": "9",
	}
	for k, v := range numeralReplacements {
		phrase = strings.ReplaceAll(phrase, k, v)
	}

	// Define exact phrases without numbers
	dayPatterns := map[string]int{
		"امروز":  0,
		"دیروز":  -1,
		"پریروز": -2,
	}
	for pattern, days := range dayPatterns {
		if phrase == pattern {
			t = now.AddDate(0, 0, days)
			return t, nil
		}
	}

	// Regex patterns to capture phrases with numbers for days, weeks, months, and years
	dayRegex := regexp.MustCompile(`(\d+)?\s?روز\sپیش`)
	weekRegex := regexp.MustCompile(`(\d+)?\s?هفته\sپیش`)
	monthRegex := regexp.MustCompile(`(\d+)?\s?ماه\sپیش`)
	yearRegex := regexp.MustCompile(`(\d+)?\s?سال\sپیش`)

	switch {
	case dayRegex.MatchString(phrase):
		t = parseTimeAgo(dayRegex, phrase, now, "day")
	case weekRegex.MatchString(phrase):
		t = parseTimeAgo(weekRegex, phrase, now, "week")
	case monthRegex.MatchString(phrase):
		t = parseTimeAgo(monthRegex, phrase, now, "month")
	case yearRegex.MatchString(phrase):
		t = parseTimeAgo(yearRegex, phrase, now, "year")
	default:
		return now, fmt.Errorf("unrecognized phrase")
	}

	return t, nil
}

// Helper function to parse time differences based on unit (day, week, month, year)
func parseTimeAgo(regex *regexp.Regexp, phrase string, now time.Time, unit string) time.Time {
	match := regex.FindStringSubmatch(phrase)
	quantity := 1 // Default to 1 if no number is specified
	if match[1] != "" {
		quantity, _ = strconv.Atoi(match[1])
	}

	switch unit {
	case "day":
		return now.AddDate(0, 0, -quantity)
	case "week":
		return now.AddDate(0, 0, -7*quantity)
	case "month":
		return now.AddDate(0, -quantity, 0)
	case "year":
		return now.AddDate(-quantity, 0, 0)
	}

	return now
}
