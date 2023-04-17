package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return strings.Join([]string{"Welcome to the Tech Palace,", strings.ToUpper(customer)}, " ")
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	starLine := strings.Repeat("*", numStarsPerLine)
	return strings.Join([]string{starLine, welcomeMsg, starLine}, "\n")
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	return strings.TrimSpace(strings.ReplaceAll(oldMsg, "*", ""))
}
