package generate

import "testing"

func TestRandomWordsUniqueness(t *testing.T) {
	words := make(map[string]struct{})
	for _, word := range randomWords() {
		if _, exists := words[word]; exists {
			t.Errorf("word %q already exists", word)
		}
		words[word] = struct{}{}
	}
}
