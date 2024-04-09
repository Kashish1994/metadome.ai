package validators

import (
	"errors"
	"github.com/eduhub/requests"
	"regexp"
)

func ValidateRequest(req interface{}) error {
	switch r := req.(type) {
	case *requests.SignUpRequest:
		return ValidateSignUpRequest(r)
	case *requests.UpdateRequest:
		return ValidateUpdateRequest(r)
	default:
		return errors.New("unsupported request type")
	}
}

func ValidateSignUpRequest(req *requests.SignUpRequest) error {
	// Check for empty fields
	if req.FirstName == "" || req.LastName == "" || req.UserName == "" ||
		req.Phone == "" || req.Address == "" || req.Password == "" {
		return errors.New("all fields are required")
	}

	// Check if password contains at least one digit and one special character
	digitRegex := regexp.MustCompile(`[0-9]`)
	specialCharRegex := regexp.MustCompile(`[^a-zA-Z0-9]`)
	if !digitRegex.MatchString(req.Password) || !specialCharRegex.MatchString(req.Password) {
		return errors.New("password must contain at least one digit and one special character")
	}
	return nil
}

func ValidateUpdateRequest(req *requests.UpdateRequest) error {
	// Check for empty fields
	if req.FirstName == "" || req.LastName == "" ||
		req.Phone == "" || req.Address == "" {
		return errors.New("all fields are required")
	}
	return nil
}
