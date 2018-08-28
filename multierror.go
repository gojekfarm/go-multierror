package multierror

import (
	"errors"
	"strings"
	"sync"
)

// MultiError implements error interface.
// An instance of MultiError has zero or more errors.
type MultiError struct {
	mutex *sync.RWMutex
	errs  []error
}

// NewMultiError: returns a thread safe instance of multierror
func NewMultiError() *MultiError {
	return &MultiError{
		mutex: &sync.RWMutex{},
	}
}

// Push adds an error to MultiError.
func (m *MultiError) Push(errString string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.errs = append(m.errs, errors.New(errString))
}

// HasError checks if MultiError has any error.
func (m *MultiError) HasError() error {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	if len(m.errs) == 0 {
		return nil
	}

	return m
}

// Error implements error interface.
func (m *MultiError) Error() string {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	formattedError := make([]string, len(m.errs))
	for i, e := range m.errs {
		formattedError[i] = e.Error()
	}

	return strings.Join(formattedError, ", ")
}
