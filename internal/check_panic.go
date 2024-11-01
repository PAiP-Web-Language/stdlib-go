package internal

// DoFPanics is a function that takes in a function and returns true if it panics
func DoFPanics(f func()) (ret bool) {
	defer func() {
		if r := recover(); r != nil {
			ret = true
		} else {
			ret = false
		}
	}()
	f()
	ret = false
	return
}

// DoFPanicsWithErr is a function that takes in a function and returns true if it panics with recovered error
func DoFPanicsWithErr(f func()) (ret bool, err any) {
	defer func() {
		if r := recover(); r != nil {
			ret = true
			err = r
		} else {
			ret = false
		}
	}()
	f()
	ret = false
	return
}
