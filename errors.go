package lei

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidLength   = errors.New("invalid length")
	ErrInvalidChecksum = errors.New("invalid checksum")
	ErrInvalidChar     = errors.New("invalid character")
	ErrUnknownRA       = errors.New("unknown registration authority")
)

func InvalidLength(len int) error {
	return fmt.Errorf("%w: %d, expected 20", ErrInvalidLength, len)
}

func InvalidChar(pos int, char rune) error {
	return fmt.Errorf("%w at position %d: %c", ErrInvalidChar, pos, char)
}

func UnknownRA(ra string) error {
	return fmt.Errorf("%w: %s", ErrUnknownRA, ra)
}
