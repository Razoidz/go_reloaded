package main

import (
	"reflect"
	"testing"
)

func TestDecimalConversion(t *testing.T) {
	tests := []struct {
		input  []string
		output []string
	}{
		{
			[]string{"1A", "(hex)"},
			[]string{"26"},
		},
		{
			[]string{"1010", "(bin)"},
			[]string{"10"},
		},
	}

	for _, test := range tests {
		result := decimalConversion(test.input)

		if !reflect.DeepEqual(result, test.output) {
			t.Errorf("expected %v got %v", test.output, result)
		}
	}
}

func TestCaseConversion(t *testing.T) {
	tests := []struct {
		input  []string
		output []string
	}{
		{
			[]string{"hello", "(up)"},
			[]string{"HELLO"},
		},
		{
			[]string{"HELLO", "(low)"},
			[]string{"hello"},
		},
		{
			[]string{"hello", "(cap)"},
			[]string{"Hello"},
		},
	}

	for _, test := range tests {
		result := caseConversion(test.input)

		if !reflect.DeepEqual(result, test.output) {
			t.Errorf("expected %v got %v", test.output, result)
		}
	}
}

func TestCaseConversion2(t *testing.T) {
	tests := []struct {
		input  []string
		output []string
	}{
		{
			[]string{"this", "is", "test", "(up,", "2)"},
			[]string{"this", "IS", "TEST"},
		},
		{
			[]string{"HELLO", "WORLD", "(low,", "2)"},
			[]string{"hello", "world"},
		},
		{
			[]string{"hello", "world", "(cap,", "2)"},
			[]string{"Hello", "World"},
		},
	}

	for _, test := range tests {
		result := caseConversion2(test.input)

		if !reflect.DeepEqual(result, test.output) {
			t.Errorf("expected %v got %v", test.output, result)
		}
	}
}

func TestPunctuation(t *testing.T) {
	input := []string{"Hello", ",", "world", "!"}
	expected := []string{"Hello,", "world!"}

	result := punctuation(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v got %v", expected, result)
	}
}

func TestQuotes(t *testing.T) {
	input := []string{"'", "hello", "world", "'"}
	expected := []string{"'hello", "world'"}

	result := quotes(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v got %v", expected, result)
	}
}

func TestVowelFix(t *testing.T) {
	tests := []struct {
		input  []string
		output []string
	}{
		{
			[]string{"a", "apple"},
			[]string{"an", "apple"},
		},
		{
			[]string{"A", "orange"},
			[]string{"An", "orange"},
		},
		{
			[]string{"a", "banana"},
			[]string{"a", "banana"},
		},
	}

	for _, test := range tests {
		result := vowelFix(test.input)

		if !reflect.DeepEqual(result, test.output) {
			t.Errorf("expected %v got %v", test.output, result)
		}
	}
}
