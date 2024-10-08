package hello

import (
	"fmt"
	"strings"
)

/*
Hello
This function is exported, so it is accessible from outside the package.
*/
func Hello(who string) {
	hello(who)
}

/*
hello
This function is not exported, so it is not accessible from outside the package.
*/
func hello(who string) {
	if len(strings.TrimSpace(who)) == 0 {
		who = "World"
	}
	fmt.Printf("Hello, %s!"+"\n", who)
}
