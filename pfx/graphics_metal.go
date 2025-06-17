//go:build darwin

package pfx

import (
	"github.com/obaraelijah/go-pfx/hal"
	"github.com/obaraelijah/go-pfx/internal/metal"
)

func DefaultGraphics() hal.Graphics {
	return metal.NewGraphics()
}
