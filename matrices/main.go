//create 3 X 3 Matrix , take input from user

package main

import "fmt"

func main() {

	s := [2][3]int{{0, 0, 0}, {0, 0, 0}}

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("s[%d][%d]=", i, j)
			fmt.Scanf("%d", &s[i][j])
		}
	}
	fmt.Println(s)
}
