package errors

import "fmt"

// ErrorType represents different types of errors
type ErrorType int

const (
	// ErrorTypeValidation represents validation errors
	ErrorTypeValidation ErrorType = iota
	// ErrorTypeConfig represents configuration errors
	ErrorTypeConfig
	// ErrorTypeDocker represents Docker-related errors
	ErrorTypeDocker
	// ErrorTypeFileIO represents file I/O errors
	ErrorTypeFileIO
	// ErrorTypeGeneration represents generation errors
	ErrorTypeGeneration
	// ErrorTypeUnknown represents unknown errors
	ErrorTypeUnknown
)

// AppError represents an application error with context
type AppError struct {
	Type    ErrorType
	Message string
	Err     error
}

// Error implements the error interface
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// Unwrap returns the wrapped error
func (e *AppError) Unwrap() error {
	return e.Err
}

// NewValidationError creates a new validation error
func NewValidationError(message string, err error) *AppError {
	return &AppError{
		Type:    ErrorTypeValidation,
		Message: message,
		Err:     err,
	}
}

// NewConfigError creates a new configuration error
func NewConfigError(message string, err error) *AppError {
	return &AppError{
		Type:    ErrorTypeConfig,
		Message: message,
		Err:     err,
	}
}

// NewDockerError creates a new Docker error
func NewDockerError(message string, err error) *AppError {
	return &AppError{
		Type:    ErrorTypeDocker,
		Message: message,
		Err:     err,
	}
}

// NewFileIOError creates a new file I/O error
func NewFileIOError(message string, err error) *AppError {
	return &AppError{
		Type:    ErrorTypeFileIO,
		Message: message,
		Err:     err,
	}
}

// NewGenerationError creates a new generation error
func NewGenerationError(message string, err error) *AppError {
	return &AppError{
		Type:    ErrorTypeGeneration,
		Message: message,
		Err:     err,
	}
}

// GetUserFriendlyMessage returns a user-friendly error message
func GetUserFriendlyMessage(err error) string {
	if appErr, ok := err.(*AppError); ok {
		switch appErr.Type {
		case ErrorTypeValidation:
			return fmt.Sprintf("Validation error: %s", appErr.Message)
		case ErrorTypeConfig:
			return fmt.Sprintf("Configuration error: %s", appErr.Message)
		case ErrorTypeDocker:
			return fmt.Sprintf("Docker error: %s\nMake sure Docker is running and you have the necessary permissions.", appErr.Message)
		case ErrorTypeFileIO:
			return fmt.Sprintf("File error: %s", appErr.Message)
		case ErrorTypeGeneration:
			return fmt.Sprintf("Generation error: %s", appErr.Message)
		default:
			return fmt.Sprintf("Error: %s", appErr.Message)
		}
	}
	return err.Error()
}
