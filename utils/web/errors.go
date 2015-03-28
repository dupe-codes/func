// Defines helpful error types for interacting with HTTP requests/responses

package web

type GeneralError struct {
	Message string
}

func (err *GeneralError) Error() string { return err.Message }

type InvalidFieldsError struct {
	GeneralError
	Fields []string
}
