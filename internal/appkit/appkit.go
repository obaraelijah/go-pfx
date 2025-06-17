package appkit

import (
	"sync/atomic"

	"github.com/obaraelijah/go-pfx/hal"
)

/*
#cgo darwin LDFLAGS: -framework AppKit -framework QuartzCore
#include "appkit.h"
*/

import "C"

var (
	runCounter atomic.Uint32
	halCfg     hal.PlatformConfig
)

func NewPlatform() hal.Platform {
	return &Platform{}
}

type Platform struct {
}

func (p *Platform) Run(cfg hal.PlatformConfig) error {
	if !runCounter.CompareAndSwap(0, 1) {
		return hal.ErrAlreadyRunning
	}

	halCfg = cfg

	r := C.pfx_ak_run()

	switch r {
	case C.PFX_SUCCESS:
		return nil

	case C.PFX_NOT_MAIN_THREAD:
		return hal.ErrNotMainThread

	default:
		panic("unexpected response")
	}
}

//export pfx_ak_init_callback
func pfx_ak_init_callback() {
	if err := halCfg.Init(); err != nil {
		panic(err)
	}
}

func (p *Platform) Exit() {
	C.pfx_ak_stop()

	windows.Range(func(key, value any) bool {
		ptr := value.(C.id)
		C.pfx_ak_free_context(ptr)

		return true
	})

	windows.Clear()
}
