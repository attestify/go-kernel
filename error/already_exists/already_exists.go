package already_exists

// AlreadyExists indicates that an action to create a new entity was unsuccessful
// because the entity being created already exists/.
// The message container IS MEANT to be for end users.
type AlreadyExists struct {
	message string
}

func New(message string) AlreadyExists {
	return AlreadyExists{
		message: message,
	}
}

func (error AlreadyExists) Error() string {
	return error.message
}
