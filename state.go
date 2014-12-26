// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package keyboard

// State represents a single keyboard key state.
type State uint8

// Keyboard key state constants, Down implies the key is currently pressed
// down, and up implies it is not. The InvalidState is declared to help users
// detect uninitialized variables.
const (
	InvalidState State = iota
	Down
	Up
)

// Aliases for lock keys (e.g. Caps Lock, Num Lock, etc). When the key is
// "locked" it is considered in the Down state, when it is "unlocked" it is in
// the Up state.
const (
	On  = Down
	Off = Up
)
