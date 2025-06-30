//go:build !darwin

package pfx

import (
	"github.com/obaraelijah/go-pfx/hal"
)

func MetalGraphicsEnabled() bool {
	return false
}

func MetalGraphics() hal.Graphics {
	panic("unsupported platform")
}
