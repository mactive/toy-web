package main

func main() {

}

func LargestRange(array []int) []int {
	visited := make(map[int]int)
	for i := 0; i < len(array); i++ {
		visited[array[i]] = 1
	}
	largestRange := []int{array[0], array[1]}
	for i := 0; i < len(array); i++ {
		currentNum := array[i]
		if visited[currentNum] == 2 {
			continue
		}
		leftNum := currentNum - 1
		rightNum := currentNum + 1
		for ; visited[leftNum] != 0; {
			visited[leftNum] = 2
			leftNum -= 1
		}

		for ; visited[rightNum] != 0; {
			visited[rightNum] = 2
			rightNum += 1
		}


		if rightNum - 1 - (leftNum + 1) > largestRange[1] - largestRange[0] {
			largestRange[0] = leftNum + 1
			largestRange[1] = rightNum - 1
		}
	}

	return largestRange
}
