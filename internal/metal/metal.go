package metal

/*
#cgo darwin LDFLAGS: -framework Metal

#include "metal.h"
*/
import "C"

import (
	"unsafe"

	"github.com/obaraelijah/go-pfx/hal"
)

type Graphics struct {
	device C.id
	queue  C.id
}

func NewGraphics() hal.Graphics {
	return &Graphics{}
}

func (g *Graphics) Init(cfg hal.GPUConfig) error {
	r := C.pfx_mtl_open(&g.device, &g.queue)

	switch r {
	case C.PFX_SUCCESS:
		return nil

	default:
		panic("unexpected response")
	}
}

func (g *Graphics) CreateSurface(rawWH hal.WindowHandle) (hal.Surface, error) {
	wh, ok := rawWH.(hal.MetalWindowHandle)
	if !ok {
		return nil, hal.ErrUnsupportedWindowHandle
	}

	layer := C.id(wh.Layer)

	r := C.pfx_mtl_configure_surface(g.device, layer)

	switch r {
	case C.PFX_SUCCESS:
		return &Surface{
			graphics: g,
			layer:    layer,
		}, nil

	default:
		panic("unexpected response")
	}
}

type Surface struct {
	graphics *Graphics
	layer    C.id
}

func (s *Surface) AcquireTexture() (hal.SurfaceTexture, error) {
	var (
		drawable C.id
		texture  C.id
	)

	C.pfx_mtl_acquire_texture(s.layer, &drawable, &texture)

	return &SurfaceTexture{
		graphics: s.graphics,
		drawable: drawable,
		texture:  texture,
	}, nil
}

type SurfaceTexture struct {
	graphics *Graphics
	drawable C.id
	texture  C.id
}

func (s *SurfaceTexture) Present() error {
	C.pfx_mtl_present_texture(s.graphics.queue, s.drawable)

	return nil
}

func (s *SurfaceTexture) Discard() {
	C.pfx_mtl_discard_surface_texture(s.drawable)
}

func (s *SurfaceTexture) View() hal.TextureView {
	// TODO: ownership

	return &TextureView{
		texture: s.texture,
	}
}

type TextureView struct {
	texture C.id
}

type CommandBuffer struct {
	buffer        C.id
	renderEncoder C.id
}

func (g *Graphics) CreateCommandBuffer() hal.CommandBuffer {
	var buf C.id

	// TODO: synchronise
	C.pfx_mtl_create_command_buf(g.queue, &buf)

	return &CommandBuffer{
		buffer: buf,
	}
}

func (c *CommandBuffer) BeginRenderPass(description hal.RenderPassDescriptor) {
	cAttachs := make([]C.ColorAttachment, len(description.ColorAttachments))

	for i, c := range description.ColorAttachments {
		tv, ok := c.View.(*TextureView)
		if !ok {
			panic("unexpected view type")
		}

		cAttachs[i] = C.ColorAttachment{
			view:  tv.texture,
			load:  C.bool(c.Load),
			store: C.bool(!c.Discard),
			r:     C.double(c.ClearColor.R),
			g:     C.double(c.ClearColor.G),
			b:     C.double(c.ClearColor.B),
			a:     C.double(c.ClearColor.A),
		}
	}

	cAttachsPtr := unsafe.SliceData(cAttachs)

	C.pfx_mtl_begin_rpass(
		c.buffer,
		cAttachsPtr,
		C.uint64_t(len(cAttachs)),
		&c.renderEncoder,
	)
}

func (c *CommandBuffer) EndRenderPass() {
	C.pfx_mtl_end_rpass(c.renderEncoder)
}

func (c *CommandBuffer) Submit() {
	C.pfx_mtl_cbuf_submit(c.buffer)
}
