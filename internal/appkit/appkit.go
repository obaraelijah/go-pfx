package appkit

import (
	"errors"
	"sync/atomic"
)
import "C"

/*
#cgo darwin LDFLAGS: -framework AppKit

#include "appkit.h"
*/

var (
	runCounter        atomic.Uint32
	callbacks         Callbacks
	ErrAlreadyRunning = errors.New("already running")
	ErrNotMainThread  = errors.New("not on main thread")
)

type Callbacks struct {
	Init func()
}

func Run(cb Callbacks) error {
	if !runCounter.CompareAndSwap(0, 1) {
		return ErrAlreadyRunning
	}

	callbacks = cb

	r := C.pfx_ak_run()

	switch r {
	case C.PFX_SUCCESS:
		return nil

	case C.PFX_NOT_MAIN_THREAD:
		return ErrNotMainThread

	default:
		panic("unexpected response")
	}
}

func NewWindow(width int, height int) error {
	r := C.pfx_ak_new_window(C.int(width), C.int(height))

	switch r {
	case C.PFX_SUCCESS:
		return nil

	default:
		panic("unexpected response")
	}
}

//export pfx_ak_init_callback
func pfx_ak_init_callback() {
	callbacks.Init()
}
