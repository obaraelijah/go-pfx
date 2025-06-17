package pfx

import "github.com/obaraelijah/go-pfx/hal"

type TextureViewable interface {
	TextureView() *TextureView
}

type TextureView struct {
	view hal.TextureView
}

func viewFromHal(view hal.TextureView) *TextureView {
	return &TextureView{
		view: view,
	}
}

func (t *TextureView) TextureView() *TextureView {
	return t
}
