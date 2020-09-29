// addons.go: Example addon manager Go file
// Cale Overstreet
// September 28, 2020
/*
This is an example of an addon manager for 
an extensible Go program. The idea is to 
Use this kind of design pattern to make 
structure Jablko. Users would have to just
edit the import section and the init function 
*/

package addons

// User can add packages they want to include
import (
	"extensibility/addons/math"
	"extensibility/addons/art"
)

// Create a global map
var AddonMap = make(map[string]func(string) string, 0)

func init() {
	// User can define the string mapping for their
	// installed packages. What I could do is use
	// A template Go file and replace in the config
	// set in a different file. This would give 
	// easier extensibility and would prevent users 
	// from having to directly edit this file

	AddonMap["math"] = math.Handler
	AddonMap["math2"] = math.Handler
	AddonMap["art"] = art.Handler
}
