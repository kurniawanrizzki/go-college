package errors

type AppError struct {
	sys error
	DebugError *string
	Message string
	Code Code
}