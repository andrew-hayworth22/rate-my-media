package business

import "fmt"

var TRIM_SET = " \n\t"

type ValidationError struct {
	Errors map[string]string `json:"validation_errors"`
}

func (ve *ValidationError) Error() string {
	err := "validation errors:\n"
	for key, val := range ve.Errors {
		err += fmt.Sprintf("%s: %s\n", key, val)
	}
	return err
}

func validText(field string) string {
	return fmt.Sprintf("Enter a valid %s", field)
}
