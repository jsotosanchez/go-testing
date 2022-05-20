package arrays

func Sum(numbers []int) (result int) {
	for _, value := range numbers {
		result += value
	}
	return
}
