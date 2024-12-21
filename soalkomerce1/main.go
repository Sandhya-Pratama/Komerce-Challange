package main

import "fmt"

func separateCharacters(input string) (string, string) {
	vowels := "aeiou"
	var vowelChars, consonantChars string

	for i := 0; i < len(input); i++ {
		char := input[i]
		if char >= 'A' && char <= 'Z' {
			char += 'a' - 'A'
		}
		if char >= 'a' && char <= 'z' {
			isVowel := false
			for j := 0; j < len(vowels); j++ {
				if char == vowels[j] {
					isVowel = true
					break
				}
			}
			if isVowel {
				vowelChars += string(char)
			} else {
				consonantChars += string(char)
			}
		}
	}

	return vowelChars, consonantChars
}

func main() {
	var input string
	fmt.Print("Input one line of words (S): ")
	fmt.Scanln(&input)

	vowelChars, consonantChars := separateCharacters(input)

	fmt.Println("Vowel Characters:", vowelChars)
	fmt.Println("Consonant Characters:", consonantChars)
}
