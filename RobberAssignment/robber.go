package RobberAssignment

import (
	"fmt"
)

func RobberAssignment() {
	arr := []int{100, 2, 15, 9, 48}

	max := 0
	for j := range arr {
		sum := 0
		for i := j; i < len(arr); {

			if j == 0 && i == len(arr)-1 {
				i = i + 2
				continue
			} else {

				sum = sum + arr[i]
				i = i + 2
			}
		}
		if sum > max {

			max = sum

		}
	}

	fmt.Println("Max money that the robber robbed= ", max)
}
