package lei

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// LEI types represent Legal Entity Identifiers as defined by ISO 17442-1:2020. The
// primary purpose of this type is to validate LEI checksums and string formatting.
// This library does not validate if an LEI is registered with GLEIF, nor does it
// convert an LEI to a registered entity name. From GLEIF:
//
// The Legal Entity Identifier (LEI) is a 20-character, alpha-numeric code based on the
// ISO 17442 standard developed by the International Organization for Standardization
// (ISO). It connects to key reference information that enables clear and unique
// identification of legal entities participating in financial transactions.
type LEI string

// Parse a string into an LEI, returning any errors if the LEI is not valid.
func Parse(s string) (LEI, error) {
	entity := LEI(s)
	if err := entity.Check(); err != nil {
		return LEI(""), err
	}
	return entity, nil
}

// Check if the LEI is valid. Returns ErrInvalidLength if the LEI is not 20 characters,
// returns ErrInvalidChar if any incorrect characters are in the LEI, returns
// ErrInvalidChecksum if the Mod97 checksum is incorrect. Returns nil if the LEI is a
// valid LEI with a valid checksum.
func (s LEI) Check() error {
	if len(s) != 20 {
		return InvalidLength(len(s))
	}

	checksum, err := Mod97(string(s))
	if err != nil {
		return err
	}

	if checksum != 1 {
		return ErrInvalidChecksum
	}

	return nil
}

// Generates a random LEI for testing purposes only.
func Random() LEI {
	prefix := randString(4)
	infix := randString(12)
	checksum, _ := Mod97(fmt.Sprintf("%s00%s00", prefix, infix))
	return LEI(fmt.Sprintf("%s00%s%02d", prefix, infix, 98-checksum))
}

const (
	char0 = '0'
	char9 = '9'
	charA = 'A'
	charZ = 'Z'
	code0 = uint32('0')
	codeA = uint32('A')
)

// Mod97 computes the Mod97 checksum as defined by ISO 7064.
func Mod97(s string) (uint32, error) {
	var buffer uint32
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		// Check for valid digits
		if !isDigit(runes[i]) {
			return 0, InvalidChar(i, runes[i])
		}

		// Compute the buffer
		charCode := uint32(runes[i])
		if charCode >= codeA {
			buffer = charCode + (buffer*100 - codeA + 10)
		} else {
			buffer = charCode + (buffer*10 - code0)
		}

		if buffer > 10000000 {
			buffer %= 97
		}
	}

	return buffer % 97, nil
}

// Returns true if the rune is a valid LEI character.
func isDigit(r rune) bool {
	return (r >= char0 && r <= char9) || (r >= charA && r <= charZ)
}

// The following constants are used for fast random string generation.
// See: https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
const letterBytes = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// A unique random source (not cryptographically secure).
var src = rand.NewSource(time.Now().UnixNano())

// Generate a random string of length n with characters 0-9A-Z.
func randString(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)

	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}
