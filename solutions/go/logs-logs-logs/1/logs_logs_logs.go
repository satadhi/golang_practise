package logs


import "fmt"


func Application(log string) string {

	fmt.Printf("%s\n", log)
    charMap := map[rune]string{
		'❗': "recommendation",
		'🔍': "search",
		'☀': "weather",
	}

	for _, char := range log {
		if app, ok := charMap[char]; ok {
			return app
		}
	}

	return "default"
}
// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var result []rune 

    for _, val := range log {
        if val == oldRune {
            result = append(result, newRune)
        } else {
            result = append(result, val)
        }
    }

    return string(result)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
return len([]rune(log)) <= limit
}
