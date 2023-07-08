package errors

import "errors"

const (
	// NotFound error indicates a missing / not found record
	NotFound        = "NotFound"
	NotFoundMessage = "Record Not Found"

	// ValidationError indicates an error in input validation
	ValidationError        = "ValidationError"
	ValidationErrorMessage = "Validation error"

	// ResourceAlreadyExists indicates a duplicate / already existing record
	ResourceAlreadyExists     = "ResourceAlreadyExists"
	AlreadyExistsErrorMessage = "Resource already exists"

	// RepositoryError indicates an repository (e.g. database) error
	RepositoryError        = "RepositoryError"
	RepositoryErrorMessage = "Error in repository operation"

	// NotAuthenticated indicates an authentication error
	NotAuthenticated             = "NotAuthenticated"
	NotAuthenticatedErrorMessage = "Not Authenticated"

	// TokenGeneratorError indicates an token generation error
	TokenGeneratorError        = "TokenGeneratorError"
	TokenGeneratorErrorMessage = "Error in token generation"

	// NotAuthorized indicates an authorization error
	NotAuthorized             = "NotAuthorized"
	NotAuthorizedErrorMessage = "Not authorized"

	// UnknownError indicates an error that the app cannot find the cause for
	UnknownError        = "UnknownError"
	UnknownErrorMessage = "Something went wrong"
)

type AppError interface {
	Error() string
}

type AppErrorImpl struct {
	Err  error
	Type string
}

func NewAppErrorImpl(err error, errType string) AppError {
	return &AppErrorImpl{
		Err:  err,
		Type: errType,
	}
}

func (a *AppErrorImpl) Error() string {
	return a.Err.Error()
}

// NewAppErrorWithType initialize a new default error for given type.
func NewAppErrorWithType(errType string) AppError {
	var err error
	switch errType {
	case NotFound:
		err = errors.New(NotFoundMessage)
	case ValidationError:
		err = errors.New(ValidationErrorMessage)
	case ResourceAlreadyExists:
		err = errors.New(AlreadyExistsErrorMessage)
	case RepositoryError:
		err = errors.New(RepositoryErrorMessage)
	case NotAuthenticated:
		err = errors.New(NotAuthenticatedErrorMessage)
	case NotAuthorized:
		err = errors.New(NotAuthorizedErrorMessage)
	case TokenGeneratorError:
		err = errors.New(TokenGeneratorErrorMessage)
	default:
		err = errors.New(UnknownErrorMessage)
	}
	return NewAppErrorImpl(err, errType)
}
