package bot

import "github.com/BrandokVargas/dukibot/internal/domain"

func getMostSimilarCommand(input string, commands map[string]domain.Command) (string, float64) {
	highestScore := 0.0
	var bestMatch string

	for cmd := range commands {
		similarity := stringSimilarity(input, cmd)
		if similarity > highestScore {
			highestScore = similarity
			bestMatch = cmd
		}
	}
	return bestMatch, highestScore
}

func stringSimilarity(a, b string) float64 {
	matches := 0
	minLen := len(a)
	if len(b) < minLen {
		minLen = len(b)
	}
	for i := 0; i < minLen; i++ {
		if a[i] == b[i] {
			matches++
		}
	}
	return float64(matches) / float64(len(b))
}
