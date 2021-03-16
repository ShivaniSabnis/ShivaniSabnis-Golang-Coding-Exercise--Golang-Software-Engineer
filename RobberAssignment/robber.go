package RobberAssignment

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RobberAssignment() {
	arr := make([]int, 0)
	fmt.Println("Enter space seperated house numbers:")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		scn := scanner.Text()
		str := strings.Split(scn, " ")
		for i := range str {
			s, _ := strconv.Atoi(str[i])
			arr = append(arr, s)
		}
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
		os.Exit(1)
	}
}
