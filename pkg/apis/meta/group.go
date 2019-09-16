// Package meta contains meta API versions
package meta

// StatusInterface allows you to return your stack info without having know the object
type StatusInterface interface {
	GetStackID() string
}
