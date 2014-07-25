// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package keyboard

import (
	"bytes"
	"fmt"
	"sync"
)

// Watcher watches the state of various keyboard keys.
type Watcher struct {
	access    sync.RWMutex
	states    map[Key]State
	rawStates map[uint64]State
}

// String returns a multi-line string representation of this keyboard watcher
// and it's associated states (but not raw ones).
func (w *Watcher) String() string {
	bb := new(bytes.Buffer)
	fmt.Fprintf(bb, "Watcher(\n")
	for k, s := range w.States() {
		fmt.Fprintf(bb, "    %v: %v\n", k, s)
	}
	fmt.Fprintf(bb, ")")
	return bb.String()
}

// SetState specifies the current state of the specified key.
func (w *Watcher) SetState(key Key, state State) {
	w.access.Lock()
	defer w.access.Unlock()

	w.states[key] = state
}

// States returns an copy of the internal key state map used by this watcher.
func (w *Watcher) States() map[Key]State {
	w.access.RLock()
	defer w.access.RUnlock()

	copy := make(map[Key]State)
	for key, state := range w.states {
		copy[key] = state
	}
	return copy
}

// State returns the current state of the specified key.
func (w *Watcher) State(key Key) State {
	w.access.Lock()
	defer w.access.Unlock()

	state, ok := w.states[key]
	if !ok {
		w.states[key] = Up
	}
	return state
}

// Down tells whether the specified key is currently in the down state.
func (w *Watcher) Down(key Key) bool {
	return w.State(key) == Down
}

// Up tells whether the specified key is currently in the up state.
func (w *Watcher) Up(key Key) bool {
	return w.State(key) == Up
}

// SetRawState specifies the current state of the specified raw key value.
func (w *Watcher) SetRawState(raw uint64, state State) {
	w.access.Lock()
	defer w.access.Unlock()

	w.rawStates[raw] = state
}

// RawStates returns an copy of the internal raw key state map used by this watcher.
func (w *Watcher) RawStates() map[uint64]State {
	w.access.RLock()
	defer w.access.RUnlock()

	copy := make(map[uint64]State)
	for raw, state := range w.rawStates {
		copy[raw] = state
	}
	return copy
}

// RawState returns the current state of the specified raw key value.
func (w *Watcher) RawState(raw uint64) State {
	w.access.Lock()
	defer w.access.Unlock()

	state, ok := w.rawStates[raw]
	if !ok {
		w.rawStates[raw] = Up
	}
	return state
}

// RawDown tells whether the specified raw key value is currently in the down state.
func (w *Watcher) RawDown(raw uint64) bool {
	return w.RawState(raw) == Down
}

// RawUp tells whether the specified raw key value is currently in the up state.
func (w *Watcher) RawUp(raw uint64) bool {
	return w.RawState(raw) == Up
}

// NewWatcher returns a new, initialized, watcher.
func NewWatcher() *Watcher {
	w := new(Watcher)
	w.states = make(map[Key]State)
	w.rawStates = make(map[uint64]State)
	return w
}
