package max_score

import "fmt"

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	scores := make([]int, n)
	for i := 0; i < n; i++ {
		_, err = fmt.Scan(&scores[i])
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
	}

	best_run_score := 0
	for i := 0; i < n; i++ {
		run_score := scores[i]
		if run_score > best_run_score {
			best_run_score = run_score
		}
	}

	fmt.Println(best_run_score)
}
