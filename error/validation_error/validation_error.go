package validation_error

// ValidationError is a custom error that indicates there was an error validating a business logic aspect.
// The message container in this error is meant to be seen by the end user of the software.
type ValidationError struct {
	message string
}

func New(message string) ValidationError {
	return ValidationError{
		message: message,
	}
}

func (error ValidationError) Error() string {
	return error.message
}
