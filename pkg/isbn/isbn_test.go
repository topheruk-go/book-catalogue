package isbn

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestIsbnLength(t *testing.T) {
	tt := []struct {
		desc string
		data string
		err  string
	}{
		{desc: "invalid: format", data: "978071670344a", err: "invalid ISBN format"},
		{desc: "invalid: length 12 != 13", data: "978071670344", err: "invalid ISBN length 12"},
		{desc: "valid: isbn 13", data: "9780716703440"},
		{desc: "invalid: isbn 13", data: "9780716703410", err: "invalid ISBN value"},
		{desc: "invalid: length 14 != 13", data: "97807167034403", err: "invalid ISBN length 14"},
		{desc: "valid: isbn 10", data: "0716703440"},
		{desc: "valid: isbn 10 w/ dashes", data: "0-7167-0344-0"},
		{desc: "valid: isbn 13 w/ dashes", data: "978-0-7167-0344-0"},
	}
	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			isbn, err := Parse(tc.data)
			if err != nil {
				assert.Error(t, err, tc.err)
				return
			}

			assert.Equal(t, len(isbn.String()), 13)
		})
	}
}

func TestIsbn10(t *testing.T) {

}