//go:build !darwin && !windows

package pfx

import "github.com/obaraelijah/go-pfx/hal"

func DefaultPlatform() hal.Platform {
	panic("unsupported platform")
}
