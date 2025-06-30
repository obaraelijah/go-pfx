//go:build !darwin && !windows

package pfx

import (
	"github.com/obaraelijah/go-pfx/hal"
)

func VulkanGraphicsEnabled() bool {
	return false
}

func VulkanGraphics() hal.Graphics {
	panic("unsupported platform")
}
