// https://husobee.github.io/golang/database/2015/06/12/scanner-valuer.html

package isbn

import (
	"database/sql/driver"
	"reflect"
)

func (isbn ISBN) Value() (driver.Value, error) {
	return nil, ErrTodo
}

func (isbn *ISBN) Scan(v any) (err error) {
	switch u := v.(type) {
	// case []byte: //may not be required
	// *isbn, err = ParseBytes(v)
	case string:
		*isbn, err = Parse(u)
	default:
		return invalidTypeError{reflect.TypeOf(v)}
	}
	return err
}
