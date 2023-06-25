package main

import "fmt"
import "strings"

const sentence = `
Good morning.
My lady.
`

func main() {
	fmt.Println(sentence)
	fmt.Println(strings.Join([]string {"Hyperledger", "Fabric", "."}, " "))
}
