package config

import "fmt"

type UnprocessableError struct {
	ErrorCode int
	Message   string
}

/** 
 * UnprocessableError implements the error interface for unprocessable entity errors.
 */
func (e *UnprocessableError) Error() string {
	return fmt.Sprintf("UnprocessableError %d: %s", e.ErrorCode, e.Message)
}