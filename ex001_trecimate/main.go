/*
First, you mash in a random large number to start with. Then, repeatedly do the following:

If the number is divisible by 3, divide it by 3.

If it's not, either add 1 or subtract 1 (to make it divisible by 3), then divide it by 3.

The game stops when you reach "1".

*/

package main

import (
	"fmt"
)

func main() {

	trecimate(33)
	// expected output of trecimate(33)
	// Have 33, adding 0
	// Have 11, adding 1
	// Have 4, adding -1
	// Have 1, stopping

}

func trecimate(n int) {
	fmt.Println("Replace this to get the required output")
}
