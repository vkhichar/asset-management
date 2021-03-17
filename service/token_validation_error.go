package service

// Return this error with appropriate message

type InvalidTokenError struct {
	message string
}

func NewInvalidTokenError(message string) InvalidTokenError {
	return InvalidTokenError{message: message}
}

func (error InvalidTokenError) Error() string {
	return error.message
}
