package half_arr_len_frequency

import "fmt"

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	numbers := make([]int, n)
	for i := range numbers {
		if _, err := fmt.Scan(&numbers[i]); err != nil {
			fmt.Println("Error reading imput:", err)
			return
		}
	}

	numberToFrequency := make(map[int]int, n)
	for _, value := range numbers {
		numberToFrequency[value]++
	}

	half_length := n / 2
	for key, value := range numberToFrequency {
		if value > half_length {
			fmt.Printf("%d", key)
			return
		}
	}

	fmt.Printf("%d", -1)
}
