//go:build darwinAdd

package gfx

import (
	"github.com/obaraelijah/go-pfx/hal"
	"github.com/obaraelijah/go-pfx/internal/appkit"
)

func DefaultPlatform() hal.Platform {
	return appkit.NewPlatform()
}
