package appkit

import (
	"sync"
	"sync/atomic"
	"unsafe"

	"github.com/obaraelijah/go-pfx/hal"
)

/*
#include "appkit.h"
*/
import "C"

var (
	windowCounter atomic.Uint32
	windows       sync.Map
)

func (p *Platform) NewWindow(cfg hal.WindowConfig) (hal.Window, hal.WindowHandle, error) {
	var (
		res   C.id
		resWH C.id
	)

	id := hal.Window(windowCounter.Add(1))

	r := C.pfx_ak_new_window(C.uint64_t(id), C.int(cfg.Width), C.int(cfg.Height), &res, &resWH)

	switch r {
	case C.PFX_SUCCESS:
		windows.Store(id, res)
		return id, hal.MetalWindowHandle{Layer: unsafe.Pointer(resWH)}, nil

	default:
		panic("unexpected response")
	}
}

func (p *Platform) CloseWindow(id hal.Window) {
	raw, ok := windows.Load(id)
	if !ok {
		return
	}

	C.pfx_ak_close_window(raw.(C.id))
}

//export pfx_ak_close_requested_callback
func pfx_ak_close_requested_callback(id uint64) {
	halCfg.WindowCloseRequested(hal.Window(id))
}

//export pfx_ak_window_closed_callback
func pfx_ak_window_closed_callback(id uint64) {
	wid := hal.Window(id)

	halCfg.WindowClosed(wid)

	raw, ok := windows.LoadAndDelete(wid)
	if !ok {
		return
	}

	C.pfx_ak_free_context(raw.(C.id))
}

//export pfx_ak_draw_callback
func pfx_ak_draw_callback(id uint64) {
	halCfg.WindowRender(hal.Window(id))
}
