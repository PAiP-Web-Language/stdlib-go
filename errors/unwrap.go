package errors

type UnwrapError struct {
}

func (UnwrapError) Error() string {
	return "UnwrapError: Tried unwrap null value"
}

