package main

import (
	"fmt"
)

func main() {
	var queue []string
	searched := make(map[string]bool)
	res := 0

	graph := make(map[string][]string)
	graph["S"] = []string{"D", "M"}
	graph["D"] = []string{"L", "F"}
	graph["M"] = []string{"L", "T"}
	graph["L"] = []string{}
	graph["T"] = []string{"F"}
	graph["F"] = []string{}

	queue = append(queue, "S")
	for len(queue) > 0 {
		if _, found := searched[queue[0]]; found {
			queue = queue[1:]

			continue
		}

		searched[queue[0]] = true

		if queue[0] == "F" {
			fmt.Println("Searched F! res: ", res)

			return
		} else {
			for _, value := range graph[queue[0]] {
				queue = append(queue, value)
			}

			queue = queue[1:]
		}
		fmt.Println("Current queue", queue)
		res++
	}

	fmt.Println("Doesn't success search F! res: ", res)

	return
}
