package internal_error

// InternalError is a custom error that indicates there was an error created when crossing software boundaries.
// The message container in this error IS NOT MEANT to be for end users.
// This message is intended for operators of the system.
type InternalError struct {
	message string
}

func New(message string) InternalError {
	return InternalError{
		message: message,
	}
}

func (error InternalError) Error() string {
	return error.message
}
