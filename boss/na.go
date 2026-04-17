package main

import (
	"fmt"
	"strconv"
	"strings"
)

func applyCase(text string) string {
	words := strings.Fields(text)
	for i := 0; i < len(words); i++ {

		switch words[i] {

		case "(up,":
			nstr := strings.TrimSuffix(words[i+1], ")")
			num, _ := strconv.Atoi(nstr)
			for o := 1; o <= num; o++ {
				if i-o >= 0 {
					if words[i] == "(up," {
						words[i-o] = strings.ToUpper(words[i-o])
					}
				}
			}
			words = append(words[:i], words[i+2:]...)
			i++
		}
	}
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(applyCase("hey there, i am me and you are you (up, 4)"))
}
