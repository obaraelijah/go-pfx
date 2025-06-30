//go:build darwin || windows

package pfx

import (
	"github.com/obaraelijah/go-pfx/hal"
	"github.com/obaraelijah/go-pfx/internal/vulkan"
)

func VulkanGraphicsEnabled() bool {
	return true
}

func VulkanGraphics() hal.Graphics {
	return vulkan.NewGraphics()
}
