package last_even

import "fmt"

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	values := make([]int, n)
	for i := 0; i < n; i++ {
		_, err := fmt.Scan(&values[i])
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
	}

	last_even := -1
	for i := 0; i < n; i++ {
		if values[i]%2 == 0 {
			last_even = values[i]
		}
	}

	fmt.Printf("%v", last_even)
}
