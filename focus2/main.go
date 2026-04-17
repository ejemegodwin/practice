package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Wrong number of arguments")
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input file:", inputFile)
		return
	}

	text := string(data)
	fmt.Println("File read successfully")

	text = transformation(text)
	text = fixArticle(text)
	text = fixPunc(text)
	text = fixQuotes(text)

	err = os.WriteFile(outputFile, []byte(text), 0644)
	if err != nil {
		fmt.Println("Error writing file:", outputFile)
		return
	}
	fmt.Println("File written successfully")
}

func ToCap(word string) string {
	if len(word) == 0 {
		return word
	}
	result := strings.ToUpper(word[0:1]) + strings.ToLower(word[1:])
	return result
}
func transformation(text string) string {
	tokens := strings.Fields(text)
	var result []string

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		if token == "(hex)" && len(result) > 0 {
			val := result[len(result)-1]
			num, _ := strconv.ParseInt(val, 16, 64)
			result[len(result)-1] = strconv.FormatInt(num, 10)
			continue
		}
		if token == "(bin)" && len(result) > 0 {
			val := result[len(result)-1]
			num, _ := strconv.ParseInt(val, 2, 64)
			result[len(result)-1] = strconv.FormatInt(num, 10)
			continue
		}
		if token == "(up)" && len(result) > 0 {
			result[len(result)-1] = strings.ToUpper(result[len(result)-1])
			continue
		}
		if token == "(low)" && len(result) > 0 {
			result[len(result)-1] = strings.ToLower(result[len(result)-1])
			continue
		}
		if token == "(cap)" && len(result) > 0 {
			result[len(result)-1] = ToCap(result[len(result)-1])
			continue
		}
		if token == "(up," && i+1 < len(tokens) {
			numToken := strings.TrimSuffix(tokens[i+1], ")")
			n, err := strconv.Atoi(numToken)
			if err == nil {
				for i := 0; i < n && i < len(result); i++ {
					result[len(result)-1-i] = strings.ToUpper(result[len(result)-1-i])
				}
			}
			i++
			continue
		}
		if token == "(low," && i+1 < len(tokens) {
			numToken := strings.TrimSuffix(tokens[i+1], ")")
			n, err := strconv.Atoi(numToken)
			if err == nil {
				for i := 0; i < n && i < len(result); i++ {
					result[len(result)-1-i] = strings.ToLower(result[len(result)-1-i])
				}
			}
			i++
			continue
		}
		if token == "(cap," && i+1 < len(tokens) {
			numToken := strings.TrimSuffix(tokens[i+1], ")")
			n, err := strconv.Atoi(numToken)
			if err == nil {
				for i := 0; i < n && i < len(tokens); i++ {
					result[len(result)-1-i] = ToCap(result[len(result)-1-i])
				}
			}
			i++
			continue
		}
		result = append(result, token)
	}
	return strings.Join(result, " ")
}
func fixPunc(text string) string {
	punc := []string{".", ",", ";", ":", "?", "!"}
	for _, p := range punc {
		text = strings.ReplaceAll(text, " "+p, p)
	}
	return text
}

func fixQuotes(text string) string {
	// text = strings.ReplaceAll(text, " '", "'")
	// text = strings.ReplaceAll(text, "' ", "'")
	// return text
	//word := strings.Join(text, " ")
	re := regexp.MustCompile(`'\s+(.*?)\s+'`)
	return re.ReplaceAllString(text, "'$1'")
}

func fixArticle(text string) string {
	words := strings.Fields(text)
	for i := 0; i < len(words)-1; i++ {
		if words[i] == "a" && strings.Contains("aeiouh", strings.ToLower(string(words[i+1][0]))) {
			words[i] = "an"
		}
		if words[i] == "A" && strings.Contains("aeiouh", strings.ToLower(string(words[i+1][0]))) {
			words[i] = "An"
		}
	}
	return strings.Join(words, " ")
}
