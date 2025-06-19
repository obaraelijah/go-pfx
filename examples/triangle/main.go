package main

import (
	_ "embed"
	"unsafe"

	"log"
	"runtime"
	"time"

	"github.com/obaraelijah/go-pfx/pfx"
)

func init() {
	runtime.LockOSThread()
}

var lastPrint time.Time
var count int

//go:embed shader.metal
var metalShader string

type Example struct {
	app              *pfx.Application
	window           *pfx.Window
	shader           *pfx.Shader
	vertexFunction   *pfx.ShaderFunction
	fragmentFunction *pfx.ShaderFunction
	vertData         *pfx.Buffer
}

func (e *Example) init(app *pfx.Application) error {
	e.app = app

	log.Println("Creating main window")

	window, err := e.app.NewWindow(pfx.WindowConfig{
		Title:    "Triangle",
		Width:    800,
		Height:   600,
		OnClosed: e.closed,
		OnRender: e.render,
		OnResize: e.resize,
	})
	if err != nil {
		return err
	}

	e.window = window

	log.Println("init complete")

	e.shader, err = e.app.LoadShader(pfx.ShaderConfig{
		Source: metalShader,
	})
	if err != nil {
		return err
	}

	e.vertexFunction, err = e.shader.Function("vertexShader")
	if err != nil {
		return err
	}

	e.fragmentFunction, err = e.shader.Function("fragmentShader")
	if err != nil {
		return err
	}

	floatData := []float32{
		-0.5, -0.5, 0.0, 0,
		0.5, -0.5, 0.0, 0,
		0.0, 0.5, 0.0, 0,
	}
	byteData := unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(floatData))), len(floatData)*4)
	e.vertData = e.app.NewBuffer(byteData)

	return nil
}

func (e *Example) Run() error {
	return pfx.Run(pfx.ApplicationConfig{
		Init: e.init,
	})
}

func (e *Example) closed() {
	log.Println("Main window closed")

	e.app.Exit()
}

func (e *Example) resize(width float64, height float64) {
	log.Println("Main window resized", width, height)
}

func (e *Example) render(frame *pfx.Frame) {
	defer frame.Close()

	if time.Since(lastPrint) > time.Second {
		lastPrint = time.Now()

		log.Println("FPS", count)
		count = 0
	}

	buf := frame.NewCommandBuffer()

	rp := buf.BeginRenderPass(pfx.RenderPassDescriptor{
		ColorAttachments: []pfx.ColorAttachment{
			{
				Target:     frame,
				Load:       false,
				ClearColor: pfx.NewColor(1, 0, 1, 1),
				Discard:    false,
			},
		},
	})

	rp.End()

	buf.Submit()

	if err := frame.Present(); err != nil {
		panic(err)
	}
}

func main() {
	log.Println("Triangle Example")

	ex := &Example{}

	if err := ex.Run(); err != nil {
		panic(err)
	}

	log.Println("App exited")
}
