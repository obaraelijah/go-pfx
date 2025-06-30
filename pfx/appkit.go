//go:build darwinAdd

package pfx

import (
	"github.com/obaraelijah/go-pfx/hal"
	"github.com/obaraelijah/go-pfx/internal/appkit"
)

func AppKitPlatformEnabled() bool {
	return true
}

func AppKitPlatform() hal.Platform {
	return appkit.NewPlatform()
}
