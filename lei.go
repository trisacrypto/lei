package lei

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type LEI string

func Parse(s string) (LEI, error) {
	entity := LEI(s)
	if err := entity.Check(); err != nil {
		return LEI(""), err
	}
	return entity, nil
}

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

func isDigit(r rune) bool {
	return (r >= char0 && r <= char9) || (r >= charA && r <= charZ)
}

const letterBytes = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func randString(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
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
