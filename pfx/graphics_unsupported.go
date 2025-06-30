//go:build !darwin && !windows

package pfx

import "github.com/obaraelijah/go-pfx/hal"

func DefaultGraphics() hal.Graphics {
	panic("unsupported platform")
}
