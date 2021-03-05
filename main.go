package main

import "fmt"

/* Only the existing contents of the slice are modified. The new content added to slice has no effect to the slice
once the function returns. */
func incorrectmodifystring(slice []string) {
	for k, _ := range slice {
		slice[k] = "mod"
	}

	slice = append(slice, "c")
	slice = append(slice, "d")
	slice = append(slice, "e")
}

/* The modified slice is returned. The caller can use the returned slice. */
func correctmodifystring(slice []string) []string {
	for k, _ := range slice {
		slice[k] = "mod"
	}

	slice = append(slice, "c")
	slice = append(slice, "d")
	slice = append(slice, "e")

	return slice
}

func main() {
	var slice1 []string = []string{"a", "b"}

	/* this will print: before modification 2 [a b]. */
	fmt.Println("before modification", len(slice1), slice1)

	/* The content appended to slice has no effect. However the existing content does get changed. */
	incorrectmodifystring(slice1)

	/* this will print: after modification 2 [mod mod]. */
	fmt.Println("after incorrectmodifystring()", len(slice1), slice1)

	slice1 = correctmodifystring(slice1)

	/* this will print: after correctmodifystring 5 [mod mod c d e]. */
	fmt.Println("after correctmodifystring", len(slice1), slice1)
}
