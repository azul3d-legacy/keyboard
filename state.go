// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package keyboard

import (
	"fmt"
)

type State uint8

const (
	InvalidState State  = iota
	Down                // Being held down currently
	Up                  // No longer being held down (released)
	On           = Down // the on/active state (for lock keys; Caps Lock; Num Lock; etc..)
	Off          = Up   // the off/inactive state (for lock keys; Caps Lock; Num Lock; etc..)
)

// String returns a string representation of this keyboard key state.
func (s State) String() string {
	switch s {
	case InvalidState:
		return "InvalidState"
	case Down:
		return "Down"
	case Up:
		return "Up"
	}
	return fmt.Sprintf("State(%d)", s)
}
