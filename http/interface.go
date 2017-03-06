package http

import (
	"github.com/lara-go/larago/http/errors"
	"github.com/lara-go/larago/http/responses"
)

// ErrorsHandlerContract for every handler to resolve errors during http calls..
type ErrorsHandlerContract interface {
	// Report error to logger.
	Report(err error)

	// Render error to return to the client.
	Render(request *Request, err error) responses.Response
}

// ValidationErrorsConverter interface for requests validators.
type ValidationErrorsConverter interface {
	// ConvertValidationErrors to field - message format.
	ConvertValidationErrors(err error, tagName string, validator interface{}) *errors.ValidationErrors
}

// SelfValidator interface for json requests.
type SelfValidator interface {
	Validate() error
}

// JSONRequestValidator interface for json requests.
type JSONRequestValidator interface {
	ValidateJSON() error
}

// FormRequestValidator interface for form requests.
type FormRequestValidator interface {
	ValidateForm() error
}

// QueryRequestValidator interface for query.
type QueryRequestValidator interface {
	ValidateQuery() error
}

// ParamsRequestValidator interface for url params.
type ParamsRequestValidator interface {
	ValidateParams() error
}
