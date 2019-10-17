package gosession

import (
	stderrors "errors"
	"github.com/pkg/errors"
)

// err
var (
	ErrTokenInvalid   = stderrors.New("invalid token")
	ErrTokenIncorrect = stderrors.New("incorrect token")
)

// IsTokenError invalid token
func IsTokenError(err error) bool {
	switch errors.Cause(err) {
	case ErrTokenInvalid:
		return true
	case ErrTokenIncorrect:
		return true
	default:
		return false
	}
}

// IsInvalidToken invalid token
func IsInvalidToken(err error) bool {
	return errors.Cause(err) == ErrTokenInvalid
}

// IsIncorrectToken incorrect token
func IsIncorrectToken(err error) bool {
	return errors.Cause(err) == ErrTokenIncorrect
}
