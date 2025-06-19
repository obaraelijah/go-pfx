package pfx

import "github.com/obaraelijah/go-pfx/hal"

type Buffer struct {
	buffer hal.Buffer
}

func (a *Application) NewBuffer(data []byte) *Buffer {
	buffer := a.graphics.CreateBuffer(data)

	return &Buffer{
		buffer: buffer,
	}
}
