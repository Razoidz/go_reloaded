package main

import (
	"strings"
	"testing"
)

func TestFullTextProcessing(t *testing.T) {
	input := "If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?. " +
		"I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure " +
		"Don not be sad ,because sad backwards is das . And das not good " +
		"harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '"

	words := strings.Fields(input)

	// Run all processing steps
	words = decimalConversion(words)
	words = caseConversion(words)
	words = caseConversion2(words)
	words = punctuation(words)
	words = quotes(words)
	words = vowelFix(words)

	result := strings.Join(words, " ")

	// Full expected output after all transformations
	expected := "If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?. " +
		"I have to pack 5 outfits. Packed 26 just to be sure Don not be sad, because sad backwards is das. And das not good " +
		"Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'"

	if result != expected {
		t.Errorf("Text processing failed.\nExpected:\n%q\n\nGot:\n%q", expected, result)
	}
}
