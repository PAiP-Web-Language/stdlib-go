package errors

// NullError is a type that is used to note that a value of error is null
// This is meant for situations where you need empty error value
type NullError struct {
}

func (NullError) Error() string {
	return "NullError"
}
