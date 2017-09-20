package debugo

// PanicOnError panics if the input error is not nil
func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
