package pfx

import (
	"github.com/obaraelijah/go-pfx/hal"
	"github.com/obaraelijah/go-pfx/internal/vulkan"
)

func DefaultGraphics() hal.Graphics {
	return vulkan.NewGraphics()
}
