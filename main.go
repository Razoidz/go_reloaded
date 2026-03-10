package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"os"
	
)

func main() {
		if len(os.Args) != 3 {
		fmt.Println("Usage: program input.txt result.txt")
		return
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file: ", err)
	}
	words := strings.Fields(string(data))
	words = decimalConversion(words)
	words = caseConversion(words)
	words = caseConversion2(words)
	words = punctuation(words)
	words = quotes(words)
	words = vowelFix(words)

	resultText := strings.Join(words, " ")

	err = os.WriteFile(outputFile, []byte(resultText), 0644)
	if err != nil {
		fmt.Println("Error writing file: ", err)
	}

	// text:= "There is no greater agony than bearing a untold story inside you."
	// words := strings.Fields(text)
	// words = decimalConversion(words)
	// words = caseConversion(words)
	// words = caseConversion2(words)
	// words = quotes(words)
	// words = vowelFix(words)
	// fmt.Println(strings.Join(words, " "))
}

func decimalConversion(words []string) []string{
	for i:= 0; i< len(words); i++{
		if words[i] == "(hex)" && i > 0{
			// Convert previous word from string → int
			number, _:= strconv.ParseInt(words[i-1], 16, 64)
				// Replace previous word with decimal string
			words[i-1] = strconv.FormatInt(number, 10)
			
			
			// Remove the word "hex"
			words = append(words[:i], words[i+1:]...)
			// Adjust index because slice changed
			i--
		}
		if words[i] == "(bin)" && i > 0{
			// Convert previous word from string → int
			number, _:= strconv.ParseInt(words[i-1], 2, 64)
			
				// Replace previous word with decimal string
				words[i-1] = strconv.FormatInt(number, 10)
			
			
			
			// Remove the word "hex"
			words = append(words[:i], words[i+1:]...)
			// Adjust index because slice changed
			i--
		}
	}
	return words
}

func caseConversion(words []string)[]string{
	for i:= 0; i< len(words); i++{
		if words[i] == "(up)" && i > 0{
			words[i-1] = strings.ToUpper(words[i-1])
			words = append(words[:i], words[i+1:]...)
			i--
		}
		if words[i] == "(low)" && i > 0{
			words[i-1] = strings.ToLower(words[i-1])
			words = append(words[:i], words[i+1:]...)
			i--
		}
		if words[i] == "(cap)" && i > 0{
			words[i-1] = strings.ToUpper(words[i-1][0:1]) + strings.ToLower(words[i-1][1:])
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}
	return words
}

func caseConversion2(words []string)[]string{
	for i:= 0; i< len(words); i++{
		if words[i] == "(up," && i+1 < len(words){
			cleanNumber := strings.Trim(words[i+1], ")")
			number, _:= strconv.Atoi(cleanNumber)
			for j := 1; j <= number && i-j >= 0; j++{
				words[i-j] = strings.ToUpper(words[i-j])
			}
			words = append(words[:i], words[i+2:]...)
			i--
		}
		if words[i] == "(low," && i+1 < len(words){
			cleanNumber := strings.Trim(words[i+1], ")")
			number, _:= strconv.Atoi(cleanNumber)
			for j := 1; j <= number && i-j >= 0; j++{
				words[i-j] = strings.ToLower(words[i-j])
			}
			words = append(words[:i], words[i+2:]...)
			i--
		}
		if words[i] == "(cap," && i+1 < len(words){
			cleanNumber := strings.Trim(words[i+1], ")")
			number, _:= strconv.Atoi(cleanNumber)
			
			for j := 1; j <= number && i-j >= 0; j++{
			words[i-j] = strings.ToUpper(words[i-j][0:1]) + strings.ToLower(words[i-j][1:])
			}
			words = append(words[:i], words[i+2:]...)
			i--
		}
	}
	return words
}

func punctuation(words []string) []string {
	for i := 1; i < len(words); i++ {
		for len(words[i]) > 0 && unicode.IsPunct(rune(words[i][0])) {
			words[i-1] += string(words[i][0])
			words[i] = words[i][1:]
		}
		if words[i] == "" {
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}
	return words
}

func quotes(words []string) []string{
	quoteCount := 0

	for i := 0; i < len(words); i++{
			
			if words[i] == "'"{
				
				if quoteCount % 2 == 0 && i+1 < len(words){
					words[i+1] = words[i] + words[i+1]
					
				}else if quoteCount % 2 != 0 && i>0{
					words[i-1] = words[i-1] + words[i]
					
			}
				words = append(words[:i], words[i+1:]...)
				i--
				quoteCount ++
			}	
	}
	return words
}

func isVowel(r rune) bool{

	switch r {
	case 'a','e','i','o', 'u':
		return true
	}
	return false
}
func vowelFix(words []string) []string{
	for i:= 0; i< len(words); i++{
		if strings.ToLower(words[i]) == "a" && i+1 < len(words){
			firstcharacter := unicode.ToLower(rune(words[i+1][0]))
			if isVowel(firstcharacter){
				if words[i] == "A"{
					words[i] = "An"
				}else {
					words[i] = "an"
				}
			} 
		}
	}
	return words
}