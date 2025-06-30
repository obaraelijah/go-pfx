//go:build windows

package pfx

import (
	"github.com/obaraelijah/go-pfx/hal"
	"github.com/obaraelijah/go-pfx/internal/windows"
)

func DefaultPlatform() hal.Platform {
	return windows.NewPlatform()
}
