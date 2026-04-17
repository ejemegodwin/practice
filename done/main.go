package main

import (
	"fmt"
	"os"
	//"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Invalid number of arguments.")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input file")
		return
	}
	fmt.Println("input file read successfully.")
	text := string(data)

	text = transformation(text)
	text = fixArticle(text)
	text = fixPunc(text)
	text = fixQuotes(text)

	err = os.WriteFile(outputFile, []byte(text), 0644)
	if err != nil {
		fmt.Println("Error writing output file")
		return
	}

	fmt.Println("output file written successfully.")

}

func capitlize(text string) string {
	result := strings.ToUpper(text[0:1]) + strings.ToLower(text[1:])
	return result
}

func transformation(text string) string {
	tokens := strings.Fields(text)
	var result []string

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		if token == "(cap)" && len(result) > 0 {
			result[len(result)-1] = capitlize(result[len(result)-1])
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
		if token == "(hex)" && len(result) > 0 {
			val := result[len(result)-1]
			num, _ := strconv.ParseInt(val, 16, 64)
			result[len(result)-1] = strconv.FormatInt(num, 10)
			continue
		}
		// if token == "(hex)" && len(result) > 0 {
		// 	val := result[len(result)-1]
		// 	num, _ := strconv.ParseInt(val, 16, 64)
		// 	result[len(result)-1] = strconv.FormatInt(num, 10)
		// }
		if token == "(bin)" && len(result) > 0 {
			val := result[len(result)-1]
			num, _ := strconv.ParseInt(val, 2, 64)
			result[len(result)-1] = strconv.FormatInt(num, 10)
			continue
		}
		if token == "(cap," && i+1 < len(tokens) {
			numToken := strings.TrimSuffix(tokens[i+1], ")")
			n, err := strconv.Atoi(numToken)

			if err == nil {
				for i := 0; i < n && i < len(result); i++ {
					result[len(result)-1-i] = capitlize(result[len(result)-1-i])
				}
			}
			i++
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
		result = append(result, token)
	}
	return strings.Join(result, " ")
}

func fixArticle(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words)-1; i++ {
		if words[i] == "a" && strings.Contains("aeiouh", strings.ToLower(string(words[i+1][0]))) {
			words[i] = "an"
		}
	}
	return strings.Join(words, " ")
}

func fixPunc(s string) string {
	w := strings.Fields(s)
	text := []string{}

	for _, ch := range w {
      for len(ch) > 0 && strings.ContainsAny(ch[:1], ".,:;?!") {
		if len(text) > 0 {
			text[len(text)-1] += ch[:1]
		}
		ch = ch[1:]
	  }
	  if ch != "" {
		text = append(text, ch)
	  }
	}
	return strings.Join(text, " ")
}

// func fixPunc(text string) string {
// 	// text = strings.TrimSpace(text)
// 	// spaceRegex := regexp.MustCompile(`\s+`)
// 	// text = spaceRegex.ReplaceAllString(text, " ")
// 	// beforePunc := regexp.MustCompile(`\s+([,.:;?!])`)
// 	// text = beforePunc.ReplaceAllString(text, "$1")
// 	// afterPunc := regexp.MustCompile(`([,.:;?!])([^\s])`)
// 	// text = afterPunc.ReplaceAllString(text, "$1 $2")

// 	// return text
// 	wor := strings.Join(text, " ")
// 	word := regexp.MustCompile(`\s+([,.:;!?]\s)`)
// 	return word.ReplaceAllString(wor, "$1")
// }

func fixQuotes(text string) string {
	w := strings.Split(text, "'")

	for i := 1; i < len(w); i+=2 {
		w[i] = strings.TrimSpace(w[i])
	}
	return strings.Join(w, "'")
}

// func Quote(s string) string {
// 	w := strings.Split(s, `"`)

// 	for i := 1; i < len(w); i+=2 {
// 		w[i] = strings.TrimSpace(w[i])
// 	}
// 	return strings.Join(w, `"`)
// }

// func fixQuotes(text string) string {
// 	// openQuote := regexp.MustCompile(`'\s+`)
// 	// text = openQuote.ReplaceAllString(text, "'")
// 	// closeQuote := regexp.MustCompile(`\s+'`)
// 	// text = closeQuote.ReplaceAllString(text, "'")

// 	// return text



// 	// quote := regexp.MustCompile(`'\s+(.*?)\s+'`)
// 	// return quote.ReplaceAllString(text, "'$1'")
// }
