// package implements the Loader of the Memory Array Redcode Simulator (MARS)
package loader

var LOADER_DESCRIPTION = `The loader loads warriors into core, converting all negative field
 values N to positive field values P such that 0 <= P < M and P = kM + N for some integer k.
 Subsequently, all field values G greater than or equal to M are converted to field values L = G
 modulo M. The loader also initializes each warrior's task queue with the appropriate task pointer.`

func WhatIsLoader() string {
	return LOADER_DESCRIPTION
}
