package windows

import (
	"sync/atomic"

	"github.com/obaraelijah/go-pfx/hal"
)

/*
#include "helper.h"
*/
import "C"

var (
	runCounter atomic.Uint32
	halCfg     hal.PlatformConfig
)

type Platform struct {
}

func NewPlatform() hal.Platform {
	return &Platform{}
}

func (p *Platform) Run(cfg hal.PlatformConfig) error {
	if !runCounter.CompareAndSwap(0, 1) {
		return hal.ErrAlreadyRunning
	}

	halCfg = cfg

	res := C.pfx_windows_init()

	switch res {
	case C.PFX_SUCCESS:
		return nil
	case C.PFX_MODULE_ERROR:
		return hal.ErrUnexpectedSystemResponse
	}

	return nil
}

//export pfx_windows_init_callback
func pfx_windows_init_callback() {
	if err := halCfg.Init(); err != nil {
		panic(err)
	}
}

func (p *Platform) Exit() {
	//TODO implement me
	panic("implement me")
}

func (p *Platform) NewWindow(cfg hal.WindowConfig) (hal.Window, hal.WindowHandle, error) {
	//TODO implement me
	panic("implement me")
}

func (p *Platform) CloseWindow(id hal.Window) {
	//TODO implement me
	panic("implement me")
}
