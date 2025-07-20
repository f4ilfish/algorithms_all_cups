package main

import "fmt"

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	if n < 1 {
		fmt.Println("Count of numbers must be positive.")
		return
	}

	numbers := make([]int, n)
	for i := range numbers {
		if _, err := fmt.Scan(&numbers[i]); err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
	}

	var sequenceLength int
	if _, err := fmt.Scan(&sequenceLength); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	if sequenceLength > len(numbers) || sequenceLength < 1 {
		fmt.Println("Sequence length must be less or equal numbers length.")
		return
	}

	dp := make([][]int, sequenceLength)
	for i := range dp {
		dp[i] = make([]int, len(numbers))
	}
	copy(dp[0], numbers)

	for i := 1; i < sequenceLength; i++ {
		for j := 0; j < len(numbers)-i; j++ {
			dp[i][j] = dp[i-1][j] * numbers[i+j]
		}
	}

	maxSequenceProduct := dp[sequenceLength-1][0]
	for j := 1; j <= (len(numbers) - sequenceLength); j++ {
		if dp[sequenceLength-1][j] > maxSequenceProduct {
			maxSequenceProduct = dp[sequenceLength-1][j]
		}
	}

	fmt.Printf("%d", maxSequenceProduct)
}
