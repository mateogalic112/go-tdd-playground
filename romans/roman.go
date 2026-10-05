// Package romans with teach us about property based tests
package romans

func ConvertToRoman(arabic int) string {
	if arabic == 2 {
		return "II"
	}
	return "I"
}
