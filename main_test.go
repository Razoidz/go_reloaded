package main

import (
	"strings"
	"testing"
)

func process(input string) string {
	words := strings.Fields(input)

	words = decimalConversion(words)
	words = caseConversion(words)
	words = caseConversion2(words)
	words = punctuation(words)
	words = quotes(words)
	words = vowelFix(words)

	return strings.Join(words, " ")
}

func TestTextProcessor(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Hex conversion",
			input:    "1E (hex) files were added",
			expected: "30 files were added",
		},
		{
			name:     "Bin conversion",
			input:    "It has been 10 (bin) years",
			expected: "It has been 2 years",
		},
		{
			name:     "Single Up/Low/Cap",
			input:    "Ready, set, go (up) ! I should stop SHOUTING (low) at the bridge (cap)",
			expected: "Ready, set, GO! I should stop shouting at the Bridge",
		},
		{
			name:     "Multiple Up with number",
			input:    "This is so exciting (cap, 2)",
			expected: "This is So Exciting",
		},
		{
			name:     "Standard Punctuation",
			input:    "I was sitting over there ,and then BAMM !!",
			expected: "I was sitting over there, and then BAMM!!",
		},
		{
			name:     "Punctuation Groups",
			input:    "I was thinking ... You were right",
			expected: "I was thinking... You were right",
		},
		{
			name:     "Quotes with spacing",
			input:    "As Elton John said: ' I am the most well-known homosexual in the world '",
			expected: "As Elton John said: 'I am the most well-known homosexual in the world'",
		},
		{
			name:     "A to An (vowel and h)",
			input:    "There it was. A amazing rock! Also a hour ago.",
			expected: "There it was. An amazing rock! Also an hour ago.",
		},
		{
			name:     "AUDIT 1",
			input:    "If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?",
			expected: "If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?",
		},
		{
			name:     "AUDIT 2",
			input:    "I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure",
			expected: "I have to pack 5 outfits. Packed 26 just to be sure",
		},
		{
			name:     "AUDIT 3",
			input:    "Don not be sad ,because sad backwards is das . And das not good",
			expected: "Don not be sad, because sad backwards is das. And das not good",
		},
		{
			name:     "AUDIT 4",
			input:    "harold wilson (cap, 2) : ' I am a hour ,but a optimist who carries a raincoat . '",
			expected: "Harold Wilson: 'I am an hour, but an optimist who carries a raincoat.'",
		},
	}

	for _, tt := range tests {
		result := process(tt.input)

		if result != tt.expected {
			t.Errorf(
				"\nTest: %s\nExpected:\n%q\nGot:\n%q\n",
				tt.name,
				tt.expected,
				result,
			)
		}
	}
}

// package main

// import (
// 	"strings"
// 	"testing"
// )

// func TestFullTextProcessing(t *testing.T) {
// 	input := "If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?. " +
// 		"I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure " +
// 		"Don not be sad ,because sad backwards is das . And das not good " +
// 		"harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '"

// 	words := strings.Fields(input)

// 	// Run all processing steps
// 	words = decimalConversion(words)
// 	words = caseConversion(words)
// 	words = caseConversion2(words)
// 	words = punctuation(words)
// 	words = quotes(words)
// 	words = vowelFix(words)

// 	result := strings.Join(words, " ")

// 	// Full expected output after all transformations
// 	expected := "If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?. " +
// 		"I have to pack 5 outfits. Packed 26 just to be sure Don not be sad, because sad backwards is das. And das not good " +
// 		"Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'"

// 	if result != expected {
// 		t.Errorf("Text processing failed.\nExpected:\n%q\n\nGot:\n%q", expected, result)
// 	}
// }
