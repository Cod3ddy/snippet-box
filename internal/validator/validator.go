package validator

import (
	"regexp"
	"slices"
	"strings"
	"unicode/utf8"
)

type Validator struct {
	FieldErrors map[string]string
	NonFieldErrors [] string
}

var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")


// Valid() returns true if the FieldErrors map doesn't contain any entries
func (v *Validator) Valid() bool {
	return len(v.FieldErrors) == 0 && len(v.NonFieldErrors) ==0
}

// Add an error messahe to the FieldErrors map (if an entry does not exist for a given key)
func (v *Validator) AddFieldError(key, message string) {
	if v.FieldErrors == nil {
		v.FieldErrors = make(map[string]string)
	}

	if _, exists := v.FieldErrors[key]; !exists {
		v.FieldErrors[key] = message
	}
}

// CheckField() adds an error message to the FieldErrors map only if a
// validation check is not 'ok'.
func (v *Validator) CheckField(ok bool, key, message string) {
	if !ok {
		v.AddFieldError(key, message)
	}
}

// Not Blank return true if a value is not empty string
func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

// MaxChars() returns true if a value contains no more than n characters
func MaxChars(value string, n int) bool{
	return utf8.RuneCountInString(value) <= n
}

// PermittedValue returns true if a value is in a list of specific permitted value
func PermittedValue[T comparable](value T, permittedValues ...T)bool{
	return slices.Contains(permittedValues, value)
}


// // MinChars() returns true if a value contains no more than n characters

func MinChars (value string, n int)  bool{
	return utf8.RuneCountInString(value) >= n
}

// Matches() returns true if a value matches a provided compiled regular expression
func Matches(value string, rx *regexp.Regexp) bool{
	return rx.MatchString(value)
}

// Add error messages to the new NonFieldErrors slice
func (v *Validator) AddNonFieldError(message string) {
	v.NonFieldErrors = append(v.NonFieldErrors, message)
}