package main

import "fmt"

func searchData(item map[string]string) map[string]string {
	fmt.Println("Enter the item you want to search:")
	var itm string
	fmt.Scanln(&itm)

	var x = make(map[string]string)

	for k, v := range item {
		if v == itm {
			x[k] = x[v]
		}
	}

	return x
}

func main() {

	data := map[string]string{
		"apple":      "fruit",
		"beans":      "vegetable",
		"brocoli":    "vegetable",
		"banana":     "fruit",
		"cherries":   "fruit",
		"carot":      "vegetable",
		"cucumber":   "vegetable",
		"grapefruit": "fruit",
		"lettuce":    "vegetable",
		"lemon":      "fruit"}

	fmt.Println(data)

	newData := searchData(data)

	fmt.Println(newData)
}
