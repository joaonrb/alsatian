package value

import (
	. "github.com/joaonrb/alsatian/result"
)

var _ Option[string] = (*bytesToStringOption)(nil)

type bytesToStringOption struct {
	Option[[]byte]
}

func (option *bytesToStringOption) validate(result Result[string]) Result[string] {
	return IfOK(option.Option.validate(IfOK(result, func(value string) Result[[]byte] {
		return OK[[]byte]{Value: []byte(value)}
	})), func(bytes []byte) Result[string] {
		return OK[string]{Value: string(bytes)}
	}).IfError(func(err error) Result[string] {
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
	})
}
