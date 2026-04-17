// None selected 

// Skip to content
// Using Gmail with screen readers
// Conversations
// 22% of 15 GB used
// Terms · Privacy · Programme Policies
// Last account activity: 13 minutes ago
// Details
package main

import (
	//"fmt"
	"strings"
	"strconv"
)
func cases(word string) string {
	words := strings.Fields(word)
	for i := 0; i < len(words); i++ {
		switch words[i] {
		// case "(up)":
		// 	words[i-1] = strings.ToUpper(words[i-1])
		// 	words = append(words[:i], words[i+1:]...)
		// 	i--
		case "(up,":
			nstr := strings.TrimSuffix(words[i+1], ")")
			num, _ := strconv.Atoi(nstr)
			for o := 1; o <= num; o++{
				if i-o >= 0 {
					if words[i] == "(up," {
						words[i-o] = strings.Title(words[i-o])
					}
				}
			}
			words = append(words[:i], words[i+2:]...)
			i--
		}
	}
	return strings.Join(words, " ")
}
//func main() {
//	fmt.Println(cases("hey there, i am me and you are you (up, 9674875750)"))
//}
// applyCases.go
// Displaying punc.go.