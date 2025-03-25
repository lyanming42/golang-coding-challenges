package main

import (
	"fmt"
)

/*
1 <= s.length <= 10^4
1 <= words.length <= 5000
1 <= words[i].length <= 30
s and words[i] consist of lowercase English letters.
*/

func main() {
	s := "wordgoodgoodgoodbestword"
	words := []string{"word", "good", "best", "good"}
	dict := make(map[string]int)

	wordLen, totalWorldLen := len(words[0]), len(words[0])*len(words)

	result := make([]int, 0)

	for start := 0; start < len(s); start++ {
		for _, word := range words {
			dict[word] = 0
		}
		for _, word := range words {
			dict[word]++
		}

		wordsUnmatched := 0

		for i := start; i < start+totalWorldLen && (i+wordLen) <= len(s); i += wordLen {
			cur := s[i : i+wordLen]
			if dict[cur] > 0 {
				dict[cur]--
			}
		}

		for _, word := range words {
			wordsUnmatched += dict[word]
		}
		fmt.Println(dict)
		if wordsUnmatched == 0 {
			result = append(result, start)
		}
	}
	fmt.Println(result)
}
