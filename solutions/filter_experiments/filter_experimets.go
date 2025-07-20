package filter_experiments

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	experiments := make([]int, n)
	for i := 0; i < n; i++ {
		_, err := fmt.Scan(&experiments[i])
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
	}

	var element int
	_, err = fmt.Scan(&element)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	var filter_experiments []int
	for i := 0; i < n; i++ {
		if experiments[i] != element {
			filter_experiments = append(filter_experiments, experiments[i])
		}
	}

	fmt_filter_experiments := make([]string, n)
	for _, value := range filter_experiments {
		fmt_filter_experiments = append(fmt_filter_experiments, fmt.Sprintf("%d", value))
	}

	fmt.Println(strings.Join(fmt_filter_experiments, " "))
}
