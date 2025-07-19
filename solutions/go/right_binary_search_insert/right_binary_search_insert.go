package right_binary_search_insert

import "fmt"

// Задача: Поиск элемента в отсортированном массиве,
// с условием, что если элемента нет, то определить его индекс,
// где он должен был бы располагаться.
// Решение: Правый бинарный поиск (последнего вхождения)
// с кастомной проверкой на добавление элемента.
func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Error due reading input:", err)
		return
	}

	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		_, err := fmt.Scan(&numbers[i])
		if err != nil {
			fmt.Println("Error due reading input:", err)
			return
		}
	}

	var target int
	_, err = fmt.Scan(&target)
	if err != nil {
		fmt.Println("Error due reading input:", err)
		return
	}

	switch n {
	case 0:
		fmt.Printf("%d", 0)
		return
	case 1:
		if numbers[0] >= target {
			fmt.Printf("%d", 0)
			return
		} else {
			fmt.Printf("%d", 1)
			return
		}
	}

	var left int = 0
	var right int = (n - 1)

	// left+1 сделан, чтобы в итоге осталось не 1, а 2 элемента
	for left+1 < right {
		var middle = left + (right-left)/2

		if target == numbers[middle] {
			fmt.Printf("%d", middle)
			return
		}

		// двигаем на middle, а не middle+1 или middle-1,
		// чтобы не потерять элемент во вхождении
		if target > numbers[middle] {
			left = middle
		} else {
			right = middle
		}
	}

	// для реализация условия, что если target не найден,
	// определить его индекс, если бы он был в массиве
	if target > numbers[right] {
		fmt.Printf("%d", (right + 1))
		return
	} else {
		fmt.Printf("%d", right)
		return
	}
}
