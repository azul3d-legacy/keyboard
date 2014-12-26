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
