// name: Landarby Vincent
//date: 06/30/2022
package main

import "fmt"

func main() {
	// declare a two-dimensional array with two sizes
	matrix := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	// add the int of the rows
	rowSum1 := matrix[0][0] + matrix[0][1] + matrix[0][2]
	rowSum2 := matrix[1][0] + matrix[1][1] + matrix[1][2]
	rowSum3 := matrix[2][0] + matrix[2][1] + matrix[2][2]

	totalRowSum := rowSum1 + rowSum2 + rowSum3 // get the total of the row

	// add the int of the columns
	columnSum1 := matrix[0][0] + matrix[1][0] + matrix[2][0]
	columnSum2 := matrix[0][1] + matrix[1][1] + matrix[2][1]
	columnSum3 := matrix[0][2] + matrix[1][2] + matrix[2][2]

	totalColumnSum := columnSum1 + columnSum2 + columnSum3 //get the total of the column

	// print the sum of the row
	fmt.Println(rowSum1)
	fmt.Println(rowSum2)
	fmt.Println(rowSum3)

	fmt.Println(totalRowSum) // print the total of the rows

	// print the sum of the column
	fmt.Println(columnSum1)
	fmt.Println(columnSum2)
	fmt.Println(columnSum3)

	fmt.Println(totalColumnSum) // print the total of the column

}
