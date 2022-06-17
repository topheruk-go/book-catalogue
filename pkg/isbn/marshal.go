package isbn

func (isbn ISBN) MarshalJSON() ([]byte, error) {
	return nil, ErrTodo
}

func (isbn *ISBN) UnmarshalJSON(b []byte) (err error) {
	*isbn, err = Parse(string(b[1 : len(b)-1]))
	return err
}
