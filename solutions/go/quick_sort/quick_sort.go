package quick_sort

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Error during reading:", err)
		return
	}

	if n <= 0 {
		fmt.Println("Count of numbers must be positive.")
		return
	}

	numbers := make([]int, n)
	for i := range numbers {
		if _, err := fmt.Scan(&numbers[i]); err != nil {
			fmt.Println("Error during reading:", err)
			return
		}
	}

	sorted := quickSort(numbers)
	strNumbers := make([]string, len(sorted))
	for i, num := range sorted {
		strNumbers[i] = strconv.Itoa(num)
	}

	fmt.Println(strings.Join(strNumbers, " "))
}

func quickSort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}

	pivot := numbers[len(numbers)/2]

	var left, middle, right []int

	for _, num := range numbers {
		switch {
		case num < pivot:
			left = append(left, num)
		case num == pivot:
			middle = append(middle, num)
		case num > pivot:
			right = append(right, num)
		}
	}

	left = quickSort(left)
	right = quickSort(right)

	return append(append(left, middle...), right...)
}
