package main

import "fmt"

/*
1 <= s.length <= 10^4
1 <= words.length <= 5000
1 <= words[i].length <= 30
s and words[i] consist of lowercase English letters.
*/

func main() {
	s := "abababab"
	words := []string{"ab", "ba"}
	dict := make(map[string]int)

	wordLen, totalWorldLen := len(words[0]), len(words[0])*len(words)
	result := make([]int, 0)

	for initalStart := 0; initalStart < wordLen; initalStart++ {
		start := initalStart
		end := totalWorldLen + initalStart
		wordsUnmatched, wordsMatched, wordsExcessUsed := 0, 0, 0

		for _, word := range words {
			dict[word] = 0
		}

		for _, word := range words {
			dict[word]++
		}

		for i := start; i < start+totalWorldLen && (i+wordLen) <= len(s); i += wordLen {
			cur := s[i : i+wordLen]
			val, ok := dict[cur]
			if ok {
				dict[cur]--
				if val <= 0 {
					wordsExcessUsed++
				} else {
					wordsMatched++
				}
			} else {
				wordsUnmatched++
			}
		}

		for ; end < len(s); end += wordLen {
			if wordsUnmatched == 0 && wordsMatched == len(words) && wordsExcessUsed == 0 {
				result = append(result, start)
			}

			prevCur := s[start : start+wordLen]

			fmt.Println(prevCur, wordsMatched, wordsUnmatched, wordsExcessUsed)

			val, ok := dict[prevCur]
			if ok {
				dict[prevCur]++
				if val < 0 {
					wordsExcessUsed--
				} else {
					wordsMatched--
				}
			} else {
				wordsUnmatched--
			}

			if end+wordLen <= len(s) {
				cur := s[end : end+wordLen]

				val, ok = dict[cur]
				if ok {
					dict[cur]--
					if val <= 0 {
						wordsExcessUsed++
					} else {
						wordsMatched++
					}
				} else {
					wordsUnmatched++
				}
				fmt.Println(cur, wordsMatched, wordsUnmatched, wordsExcessUsed)
			}
			start += wordLen

		}

		if wordsUnmatched == 0 && wordsMatched == len(words) && wordsExcessUsed == 0 {
			result = append(result, start)
		}
	}
	fmt.Println(result)
}
