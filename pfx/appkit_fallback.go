//go:build !darwin

package pfx

import (
	"github.com/obaraelijah/go-pfx/hal"
)

func AppKitPlatformEnabled() bool {
	return false
}

func AppKitPlatform() hal.Platform {
	panic("unsupported platform")
}
