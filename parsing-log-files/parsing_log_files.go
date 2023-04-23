package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<(=|~|-|\*)*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	var count int
	re := regexp.MustCompile(`(?i)".*password.*"`)

	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d*`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	result := []string{}
	re := regexp.MustCompile(`User\s+(\w+)`)

	for _, line := range lines {
		userline := re.FindStringSubmatch(line)
		switch {
		case userline != nil:
			result = append(result, fmt.Sprintf("[USR] %s %s", userline[1], line))
		default:
			result = append(result, line)
		}
	}
	return result
}
