package utix

import (
	"errors"
	"fmt"
)

type JError struct {
	Error string `json:"error"`
}

func NewJError(err error) JError {
	jerr := JError{"error"}
	if err != nil {
		jerr.Error = err.Error()
	}

	return jerr
}

func CheckErorr(e error) {

	if e != nil {
		fmt.Println(e, "ERROR FUNC IN ACTION")
	}

}

var ErrInvalidEmail = errors.New("INVALID EMAIL SAD")
var ErrEmailAlreadyExists = errors.New("ALREADY EXISTING EMAIL")
var ErrEmptyPassword = errors.New("PASSWORD CANT BE EMPTY")
var ErrShortPassword = errors.New("PASSWORD TOO SHORT ")
var ErrInvalidCredentials = errors.New("INVALID CREDENTIALS")
var ErrIncorrectPassword = errors.New("INCORRECT PASSWORD")
var ErrIncorrectEmail = errors.New("EMAIL NOT REGISTERED")
var ErrUnknown = errors.New("idk where the tf")
var ErrLogout = errors.New("LOGOUT")
