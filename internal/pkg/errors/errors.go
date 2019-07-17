package errors

// Check will panic if an error is found.
func Check(e error) {
    if e != nil {
        panic(e)
    }
}
