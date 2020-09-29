package main

import (
	"fmt"
	"extensibility/addons" // This is the module that will contain the mods
)

// Use major version to have multiple versions

func main() {
	fmt.Printf("Extensibility of Go use module init\n")

	fmt.Printf("Addons installed:\n\t%v\n\n", addons.AddonMap)	

	for {
		var input string

		fmt.Scanln(&input)	

		output := addons.AddonMap[input]
		if output != nil {
			fmt.Printf("%s\n", output("Hello"))
		} else {
			fmt.Printf("Addon \"%s\" is not installed\n", input)
		}
	}
}
