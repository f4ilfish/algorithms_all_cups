package closest_neighbors

import (
	"fmt"
	"math"
)

func main() {
	const min_neighbors_count = 2
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	if n < min_neighbors_count {
		fmt.Printf("Number of values less than %v", min_neighbors_count)
		return
	}

	neighbors := make([]int, n)
	for i := 0; i < n; i++ {
		_, err = fmt.Scan(&neighbors[i])
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
	}

	first_value := neighbors[0]
	second_value := neighbors[1]
	min_diff := math.Abs(float64(neighbors[0] - neighbors[1]))

	if n == min_neighbors_count {
		fmt.Printf("%v %v", first_value, second_value)
		return
	}

	for i := 1; i < (n - 1); i++ {
		diff := math.Abs(float64(neighbors[i] - neighbors[i+1]))
		if diff < min_diff {
			min_diff = diff
			first_value = neighbors[i]
			second_value = neighbors[i+1]
		}
	}

	fmt.Printf("%v %v", first_value, second_value)
}
