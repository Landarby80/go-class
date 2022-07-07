/* Create a map ex map[string]int{"orange": 5, "apple": 7,	"mango": 3, "strawberry": 9}
,sort the map based on key length in asecending order */

package main

import (
	"fmt"
	"sort"
)

func main() {
	data := map[string]int{"orange": 5, "apple": 7, "mango": 3, "strawberry": 9}

	// fmt.Println(data)
	newData := make([]string, 0, len(data))

	for k := range data {
		newData = append(newData, k)
	}

	sort.Strings(newData)
	fmt.Println(newData)
}
