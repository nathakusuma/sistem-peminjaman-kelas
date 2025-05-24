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

func ErrNotSubscribed() *ResponseError {
	return newError(http.StatusForbidden,
		"not-subscribed",
		"You're not subscribed to this feature. Please subscribe first.")
}

func ErrNotFound() *ResponseError {
	return newError(http.StatusNotFound,
		"not-found",
		"Resource not found.")
}

func ErrFileTooLarge() *ResponseError {
	return newError(http.StatusRequestEntityTooLarge,
		"file-too-large",
		"File size is too large. Please upload smaller file.")
}

func ErrInvalidFileFormat() *ResponseError {
	return newError(http.StatusUnprocessableEntity,
		"invalid-file-format",
		"Invalid file format. Please upload a valid file.")
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

func ErrInvalidOTP() *ResponseError {
	return newError(http.StatusUnauthorized,
		"invalid-otp",
		"Invalid OTP. Please try again or request a new OTP.")
}

func ErrInvalidRefreshToken() *ResponseError {
	return newError(http.StatusUnauthorized,
		"invalid-refresh-token",
		"Auth session is invalid. Please login again.")
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
