package pfx

import (
	"fmt"

	"github.com/obaraelijah/go-pfx/hal"
)

type ApplicationConfig struct {
	Init func(app *Application) error
}

type Application struct {
	platform hal.Platform
	graphics hal.Graphics
	cfg      ApplicationConfig
	windows  tmap[hal.Window, *Window]
}

func Run(cfg ApplicationConfig) error {
	return RunWith(cfg, DefaultPlatform(), DefaultGraphics())
}

func RunWith(cfg ApplicationConfig, platform hal.Platform, graphics hal.Graphics) error {
	app := &Application{
		platform: platform,
		graphics: graphics,
		cfg:      cfg,
	}

	if err := app.graphics.Init(hal.GPUConfig{}); err != nil {
		return fmt.Errorf("failed to init graphics: %w", err)
	}

	return app.run()
}

func (a *Application) run() error {
	return a.platform.Run(hal.PlatformConfig{
		Init:                 a.init,
		WindowCloseRequested: a.windowCloseRequested,
		WindowClosed:         a.windowClosed,
		WindowRender:         a.windowRender,
		WindowResized:        a.windowResized,
	})
}

func (a *Application) init() error {
	if a.cfg.Init == nil {
		return nil
	}

	return a.cfg.Init(a)
}

func (a *Application) Exit() {
	a.platform.Exit()

	// TODO: cleanup?
}
