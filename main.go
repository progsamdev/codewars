package main

import (
	"codewars/digitalroot"
	"codewars/firstnonrepeating"
	"fmt"
)

func main() {
	dr := digitalroot.DigitalRoot(132189)
	fmt.Printf("Sum of digits: %d\n", dr)
	fr := firstnonrepeating.FirstNonRepeating("Go hang a salami, I'm a lasagna hog!")
	fmt.Printf("First non repeting letter: %s\n", fr)
}
