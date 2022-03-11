package main

import (
	"TUTORIALYT/helper"

	"fmt"
)

func main() {
	helper.SayHello("via")
	// helper.sayGoodbye("via") // error
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // error
}
