package pfx

import "sync"

type tmap[K comparable, V any] struct {
	m sync.Map
}

func (t *tmap[K, V]) Set(k K, v V) {
	t.m.Store(k, v)
}

func (t *tmap[K, V]) Get(k K) (V, bool) {
	v, ok := t.m.Load(k)
	if !ok {
		var def V

		return def, false
	}

	return v.(V), true
}

func (t *tmap[K, V]) Remove(k K) (V, bool) {
	v, ok := t.m.LoadAndDelete(k)

	return v.(V), ok
}
