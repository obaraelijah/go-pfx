//go:build !windows

package pfx

import (
	"github.com/obaraelijah/go-pfx/hal"
)

func WindowsPlatformEnabled() bool {
	return false
}

func WindowsPlatform() hal.Platform {
	panic("unsupported platform")
}
