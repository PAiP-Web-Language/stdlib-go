package errors

type NullableError struct {
}

func (NullableError) Error() string {
	return "NullableError: Nullable contains null"
}

