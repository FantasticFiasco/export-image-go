package main

// Check will panic if an error is found.
func check(e error) {
	if e != nil {
		panic(e)
	}
}
