package main

import (
	"log"
	"time"

	"github.com/obaraelijah/go-pfx/pfx"
)

var lastPrint time.Time
var count int

type Example struct {
	app    *pfx.Application
	window *pfx.Window
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
	})
	if err != nil {
		return err
	}

	e.window = window

	log.Println("init complete")

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

func (e *Example) render() {
	count++

	if time.Since(lastPrint) > time.Second {
		lastPrint = time.Now()

		log.Println("FPS", count)
		count = 0
	}

	frame, err := e.window.BeginFrame()
	if err != nil {
		panic(err)
	}

	defer frame.Close()

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
