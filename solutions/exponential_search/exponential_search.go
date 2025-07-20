package exponential_search

import "fmt"

func main() {
	const min_numbers_count = 2

	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Error during reading imput:", err)
		return
	}

	if n < min_numbers_count {
		fmt.Printf("Count of numbers must be bigger than or equal %d", min_numbers_count)
		return
	}

	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		_, err := fmt.Scan(&numbers[i])
		if err != nil {
			fmt.Println("Error during reading imput:", err)
			return
		}
	}

	var target int
	_, err = fmt.Scan(&target)
	if err != nil {
		fmt.Println("Error during reading imput:", err)
		return
	}

	if n == min_numbers_count {
		if target > numbers[1] {
			fmt.Printf("%d %d", 0, 2)
		} else {
			fmt.Printf("%d %d", 0, 1)
		}
		return
	}

	border := 1
	for {
		next_border := border * 2

		if next_border >= len(numbers) {
			fmt.Printf("%d %d", border, len(numbers))
			break
		}

		if numbers[next_border] > target {
			fmt.Printf("%d %d", border, next_border)
			break
		}

		border = next_border
	}
}
