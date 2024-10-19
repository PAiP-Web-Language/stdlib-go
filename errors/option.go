package errors

type OptionError struct {
}

func (OptionError) Error() string {
	return "OptionError: Option evaluates to None"
}

