package config

import "fmt"

type BussinessError struct {
	ErrorCode int
	Message   string
}

/** 
 * BussinessError implements the error interface for user-related errors.
 */
func (e *BussinessError) Error() string {
	return fmt.Sprintf("BussinessError %d: %s", e.ErrorCode, e.Message)
}
