package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func ToCap(text string) string {
	return strings.ToUpper(text[:1]) + strings.ToLower(text[1:])
}

func binToDec(text string) (int64, error) {
	bin, err := strconv.ParseInt(text, 2, 64)
	if err != nil {
		return 0, err
	}
	return bin, nil
}

func ToUpper(text string) string {
	return strings.ToUpper(text)
}

func ToLower(text string) string {
	return strings.ToLower(text)
}

func hexToDec(text string) (int64, error) {
	hex, err := strconv.ParseInt(text, 16, 64)
	if err != nil {
		return 0, err
	}
	return hex, nil
}

func AorAn(text string) string {
	vowels := "aeiouAEIOU"
	if strings.ContainsAny(string(text[0]), vowels) {
		return "an " + text
	}
	return "a " + text
}

func Punctuation(texts []string) string {
	word := strings.Join(texts, " ")
	word2 := regexp.MustCompile(`\s+([.,!?;:])`)
	return word2.ReplaceAllString(word, "$1")
}

func fixQuote(text []string) string {
	word := strings.Join(text, " ")
	quote := regexp.MustCompile(`'\s+(.*?)\s+'`)
	return quote.ReplaceAllString(word, "'$1'")
}

func main() {
	input6 := []string{"As Elton John said: ' I am the most well-known homosexual in the world '"}
	fixedQuote := fixQuote(input6)
	fmt.Println("Fixed Quote:", fixedQuote)
	input5 := []string{"I", "am", "me", ",", "and", "u", "are", "u", "!", "!", "!"}
	input4 := "apple"
	input3 := "1E2A"
	input2 := "1010"
	hexe, err := hexToDec(input3)
	if err != nil {
		fmt.Println("Invalid:", err)
		return
	}
	bini, err := binToDec(input2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	input := "hello world"
	d := Punctuation(input5)
	s := ToLower(input)
	article := AorAn(input4)
	upper := ToUpper(input)
	capitalized := ToCap(input)

	fmt.Println("Punctuation: ", d)
	fmt.Println("English: ", article)
	fmt.Println("capitalized: ", capitalized)
	fmt.Println("lower: ", s)
	fmt.Println("upper: ", upper)
	fmt.Println("binary: ", bini)
	fmt.Println("hexadecimal: ", hexe)

	input7 := "this is so exciting"
	texts := strings.Fields(input7)
	result := ToUpperN(texts, 1)
	fmt.Println(strings.Join(result, " "))
	// input7 := "this is so exciting"
	// texts := strings.Fields(input7)
	// result := ToUpperN(, 2)
	// fmt.Println(strings.Join(result, " "))
	// word := strings.Fields(input7)
	// result := ToUpperN(word, 2)
	// fmt.Println(strings.Join(result, " ")) 
}


func ToUpperN(word []string, n int) []string {
	if n > len(word) {
		n = len(word)
	}
	result := make([]string, len(word))
	copy(result, word)

	for i := n + 1; i < len(word); i++ { 
		result[i] = strings.ToUpper(word[i])
	}

	return result
}