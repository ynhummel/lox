package interror

import "fmt"

type Err struct {
	Message  string
	Where    string
	Line     int
	HadError bool
}

func (e Err) Error() string {
	e.HadError = true
	return fmt.Sprintf("[line %d] Error%s: %s", e.Line, e.Where, e.Message)
}
