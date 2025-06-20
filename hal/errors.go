package hal

import (
	"errors"
)

var (
	ErrUnexpectedStatus        = errors.New("unexpected status")
	ErrAlreadyRunning          = errors.New("already running")
	ErrNotMainThread           = errors.New("not on main thread")
	ErrUnsupportedWindowHandle = errors.New("unsupported window handle")
	ErrIncompatibleSurface     = errors.New("incompatible surface")
	ErrFunctionNotFound        = errors.New("function not found")
	ErrIncompatibleDriver      = errors.New("incompatible driver")
	ErrNoSuitableDevice        = errors.New("could not find a suitable device")
	ErrMissingFeature          = errors.New("feature missing")
)
