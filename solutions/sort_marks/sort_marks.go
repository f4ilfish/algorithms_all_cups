package sort_marks

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

	marks := make([]int, n)
	for i := 0; i < n; i++ {
		_, err := fmt.Scan(&marks[i])
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
	}

	filtered_marks := make([]int, n)
	fm_i := 0
	for i := 0; i < n; i++ {
		if marks[i] == 0 {
			continue
		}
		filtered_marks[fm_i] = marks[i]
		fm_i++
	}
	for ; fm_i < n; fm_i++ {
		filtered_marks[fm_i] = 0
	}

	fmt_filtered_marks := make([]string, n)
	for i, value := range filtered_marks {
		fmt_filtered_marks[i] = fmt.Sprintf("%d", value)
	}

	fmt.Println(strings.Join(fmt_filtered_marks, " "))
}
