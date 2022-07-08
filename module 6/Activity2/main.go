package main

import (
	"Activity2/subfolder"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	num := 100

	arrayNum := subfolder.GetRandomNum(num)

	fmt.Println("the array is:", arrayNum)

	maxVal := subfolder.GetMax(arrayNum)
	fmt.Println("the max value is :", maxVal)

	indexMax := subfolder.GetIndexMax(arrayNum)
	fmt.Println("the max index is :", indexMax)

	minVal := subfolder.GetMin(arrayNum)
	fmt.Println("the min value is :", minVal)

	indexMin := subfolder.GetIndexMin(arrayNum)
	fmt.Println("the min index is :", indexMin)

	descOrder := subfolder.GetDescOrder(arrayNum)
	fmt.Println("Descending Order:", descOrder)

	ascOrder := subfolder.GetAscOrder(arrayNum)
	fmt.Println("Ascending Order:", ascOrder)

	mean := subfolder.GetMean(arrayNum)
	fmt.Println("Mean = ", mean)

	posSlice := subfolder.GetPos(arrayNum)
	fmt.Println("Positive numbers : ", posSlice)

	negSlice := subfolder.GetNeg(arrayNum)
	fmt.Println("Negative numbers : ", negSlice)

	removeDupl := subfolder.RemoveDuplicates(arrayNum)
	fmt.Println("duplicate remove : ", removeDupl)

}
