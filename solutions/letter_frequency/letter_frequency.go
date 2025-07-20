package letter_frequency

import "fmt"

func main() {
	var word string
	if _, err := fmt.Scanln(&word); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	if len(word) == 0 {
		fmt.Println("Empty string.")
		return
	}

	letterToFrequency := make(map[rune]int, len(word))
	for _, letter := range word {
		letterToFrequency[letter]++
	}

	maxFrequency := letterToFrequency[rune(word[0])]
	for _, frequency := range letterToFrequency {
		if frequency > maxFrequency {
			maxFrequency = frequency
		}
	}

	fmt.Printf("%d", maxFrequency)
}
