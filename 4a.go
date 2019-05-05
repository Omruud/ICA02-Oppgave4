package main

import (
	"fmt"
)

func Translate(expression string, language string) string {
	const string1 = "\xe5\x8c\x97\xe3\x81\xa8\xe5\x8d\x97"
	const string2 = "\x6e\x6f\x72\xc3\xb0\x75\x72 \x6f\x67 \x73\x75\xc3\xb0\x75\x72"

	r := expression
	f := language
	if r == "nord og sør" && f == "jp" {
		fmt.Printf("%s", string1)
		fmt.Println("\n")
	} else if r == "nord og sør" && f == "is" {
		fmt.Printf("%s", string2)
		fmt.Println("\n")
	} else {
		fmt.Println("error")
	}
	return ""
}
