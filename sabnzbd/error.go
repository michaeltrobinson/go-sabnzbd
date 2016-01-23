package sabnzbd

import (
	"errors"
	"strings"
)

var errorIndicator string = "error:"

var (
	ErrApikeyIncorrect error = errors.New("API Key Incorrect")
	ErrApikeyRequired  error = errors.New("API Key Required")
)

func apiStringError(str string) error {
	switch {
	case str == "":
		return nil
	case strings.Contains(str, ErrApikeyIncorrect.Error()):
		return ErrApikeyIncorrect
	case strings.Contains(str, ErrApikeyRequired.Error()):
		return ErrApikeyRequired
	default:
		return errors.New(str)
	}
}

var ErrInvalidQueueCompleteAction error = errors.New("invalid queue complete action")
