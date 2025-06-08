package errorpkg

import (
	"net/http"
)

// General
func ErrInternalServer() *ResponseError {
	return newError(http.StatusInternalServerError,
		"internal-server-error",
		"Something went wrong in our server. Please try again later.")
}

func ErrFailParseRequest() *ResponseError {
	return newError(http.StatusBadRequest,
		"fail-parse-request",
		"Failed to parse request. Please check your request format.")
}

func ErrForbiddenRole() *ResponseError {
	return newError(http.StatusForbidden,
		"forbidden-role",
		"You're not allowed to access this resource.")
}

func ErrForbiddenUser() *ResponseError {
	return newError(http.StatusForbidden,
		"forbidden-user",
		"You're not allowed to access this resource.")
}

func ErrNotFound() *ResponseError {
	return newError(http.StatusNotFound,
		"not-found",
		"Resource not found.")
}

func ErrValidation() *ResponseError {
	return newError(http.StatusUnprocessableEntity,
		"validation-error",
		"There are invalid fields in your request. Please check and try again")
}

func ErrRateLimitExceeded() *ResponseError {
	return newError(http.StatusTooManyRequests,
		"rate-limit-exceeded",
		"Rate limit exceeded. Please try again later.")
}

// Auth
func ErrCredentialsNotMatch() *ResponseError {
	return newError(http.StatusUnauthorized,
		"credentials-not-match",
		"Credentials do not match. Please try again.")
}

func ErrInvalidBearerToken() *ResponseError {
	return newError(http.StatusUnauthorized,
		"invalid-bearer-token",
		"Your auth session is invalid. Please renew your auth session.")
}

func ErrNoBearerToken() *ResponseError {
	return newError(http.StatusUnauthorized,
		"no-bearer-token",
		"You're not logged in. Please login first.")
}

func ErrEmailAlreadyRegistered() *ResponseError {
	return newError(http.StatusConflict,
		"email-already-registered",
		"Email already registered. Please login or use another email.")
}

// Proposal
func ErrReplyAlreadyExists() *ResponseError {
	return newError(http.StatusConflict,
		"reply-already-exists",
		"Reply already exists for this proposal. You cannot create another reply.")
}
