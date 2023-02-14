package value

import (
	. "github.com/joaonrb/alsatian/result"
)

var _ Option[string] = (*bytesToStringOption)(nil)

type bytesToStringOption struct {
	Option[[]byte]
}

func (option *bytesToStringOption) validate(str Result[string]) Result[string] {
	strToBytes := IfOK(str, func(value string) Result[[]byte] {
		return OK[[]byte]{Value: []byte(value)}
	})
	validatedBytes := option.Option.validate(strToBytes)
	result := IfOK(validatedBytes, func(value []byte) Result[string] {
		return OK[string]{Value: string(value)}
	})
	result.IfError(func(err error) {
		result = convertBytesErrorToStringError(err)
	})
	return result
}

func convertBytesErrorToStringError(err error) Result[string] {
	switch e := err.(type) {
	case interface {
		Value() any
		setValue(any)
	}:
		if b, ok := e.Value().([]byte); ok {
			e.setValue(string(b))
		}
	}
	return Error[string]{Err: err}
}
