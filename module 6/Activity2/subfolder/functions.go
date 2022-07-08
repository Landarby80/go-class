package subfolder

import (
	"math/rand"
)

// func to generate the random number
func GetRandomNum(n int) []int {
	arrayNum := make([]int, n) // declare an array with size n

	for i := 0; i < n; i++ {
		arrayNum[i] = rand.Intn(100-(-100)) + (-100) // generate random number into the array
	}

	return arrayNum // return the array
}

//Compute the max of an array of int

func GetMax(arrayNum []int) int {

	maxVal := arrayNum[0]

	for _, v := range arrayNum {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

//Compute the index of the max of an array of int
func GetIndexMax(arrayNum []int) int {
	indexMax := 0

	for i := 1; i < len(arrayNum); i++ {
		if arrayNum[i] > arrayNum[indexMax] {
			indexMax = i
		}
	}
	return indexMax
}

//Compute the min of an array of int

func GetMin(arrayNum []int) int {

	minVal := arrayNum[0]

	for _, v := range arrayNum {
		if v < minVal {
			minVal = v
		}
	}
	return minVal
}

//Compute the index of the min of an array of int
func GetIndexMin(arrayNum []int) int {
	indexMin := 0

	for i := 1; i < len(arrayNum); i++ {
		if arrayNum[i] < arrayNum[indexMin] {
			indexMin = i
		}
	}
	return indexMin
}

//Sort an array of int in descending order and return the new sorted array in a separate array.
func GetDescOrder(arrayNum []int) []int {
	newArray := make([]int, len(arrayNum))

	for i := 0; i < len(arrayNum); i++ {
		for j := i + 1; j < len(arrayNum); j++ {
			if arrayNum[i] < arrayNum[j] {
				temp := arrayNum[i]
				arrayNum[i] = arrayNum[j]
				arrayNum[j] = temp
			}
		}
	}
	newArray = arrayNum
	return newArray
}

//Sort an array of int in ascending order and return the new sorted array in a separate array.

func GetAscOrder(arrayNum []int) []int {
	newArray := make([]int, len(arrayNum))

	for i := 0; i < len(arrayNum); i++ {
		for j := i + 1; j < len(arrayNum); j++ {
			if arrayNum[i] > arrayNum[j] {
				temp := arrayNum[i]
				arrayNum[i] = arrayNum[j]
				arrayNum[j] = temp
			}
		}
	}
	newArray = arrayNum
	return newArray
}

//Compute the mean of an array
func GetMean(arrayNum []int) int {
	var sum int
	var num int
	for i, v := range arrayNum {
		sum += v
		num = i + 1
	}

	mean := sum / num

	return mean
}

//Identify all positive numbers in the array and return the numbers in a slice
func GetPos(arrayNum []int) []int {
	slice := make([]int, len(arrayNum))

	for i, val := range arrayNum {
		if val >= 0 {
			slice[i] = val
		}
	}

	return slice
}

//Identify all negative numbers in the array and return the numbers in a slice
func GetNeg(arrayNum []int) []int {
	slice := make([]int, len(arrayNum))

	for i, val := range arrayNum {
		if val < 0 {
			slice[i] = val
		}
	}

	return slice
}

//Remove duplicates from an array of ints and return the unique elements in a slice
func RemoveDuplicates(arrayNum []int) []int {

	keys := make(map[int]bool)
	list := []int{}

	for _, entry := range arrayNum {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
