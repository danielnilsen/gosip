package sip

import (
	"fmt"
)

// ResponseError encapsulates an unhandled >=400 SIP error response.
type ResponseError struct {
	Msg *Msg
}

func (err *ResponseError) Error() string {
	return fmt.Sprintf("%s: %d", err.Msg.Status, err.Msg.Phrase)
}