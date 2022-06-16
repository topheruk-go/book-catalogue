package isbn

import (
	"fmt"
	"strings"
)

// type ISBN string //temp

type ISBN [13]byte

func (isbn ISBN) String() string {
	return string(isbn[:])
}

var (
	defaultPrefix = "978"
)

var (
	ErrValue  = fmt.Errorf("invalid ISBN value")
	ErrFormat = fmt.Errorf("invalid ISBN format")
)

type invalidLengthError struct{ len int }

func (err invalidLengthError) Error() string {
	return fmt.Sprintf("invalid ISBN length %d", err.len)
}

func Parse(s string) (ISBN, error) {
	var isbn ISBN
	switch len(s) {
	case 13 + 4: //XXX-X-XXXX-XXXX-X
	case 10: //XXXXXXXXXX
		s = defaultPrefix + s
		fallthrough
	case 13: //XXXXXXXXXXXXX or XXX-XX-XXX-XXXX-X
		if !strings.Contains(s, "-") {
			return check13(s)
		}
		s = defaultPrefix + "-" + s
	default:
		return isbn, invalidLengthError{len(s)}
	}
	// s is now in the format XXX-X-XXXX-XXXX-X
	return check13(strings.ReplaceAll(s, "-", ""))
}

func check13(s string) (isbn ISBN, err error) {
	var acc [2]int

	for i := 0; i < len(s); i++ {
		switch v := int(s[i] - '0'); {
		case v >= 10:
			return isbn, ErrFormat
		default:
			acc[i%2] += v
			isbn[i] = s[i]
		}
	}

	if (acc[0]+acc[1]*3)%10 != 0 {
		return isbn, ErrValue
	}
	return isbn, nil
}

// Support for in future implentations would be nice
// type ISBN struct {
// 	group
//  publisher
//  title
//  check digit
// }

// Useful
// wiki https://en.wikipedia.org/wiki/ISBN
