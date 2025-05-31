package client

import "errors"

var (
	ErrErrorResponseStatus = errors.New("received error HTTP response status")
)
