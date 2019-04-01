// Package model contains all the entities
package model

// The exception entity
type Exception struct {
	// exception code int
	CodeInt uint64

	// exception code string
	CodeString string

	// exception descrition string
	CodeDescription string
}

func (e *Exception) Error() string {
    return e.CodeString
}

// create a new blank exception
func NewException(codeString string, codeDescription string) *Exception {
	NewException := &Exception{
		CodeInt:         1,
		CodeString:      codeString,
		CodeDescription: codeDescription,
	}
	return NewException
}
