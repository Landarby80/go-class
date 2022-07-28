/*
1)Create Text file(Documants containing summary of context like movie review ,either in possitive or negetive )
2)Read the text file and depending upon the words available in the Document
give a conclusion about review (Possitive/negetive) in terms of % of possitive/negetive
3)Give high priority to Awesome than Good
*/

package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// read the the text file
	data, err := ioutil.ReadFile("/Users/vince/Documents/goClass/go-class/module 9/review.txt")

	//check for errors
	if err != nil {
		panic(err)
	}
	// convert the file in string
	strData := string(data)

	goodWords := []string{}

	fmt.Println(strData)
}
