// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package keyboard

import (
	"fmt"
	"time"
)

// StateEvent represents an event when an keyboard key changes state (i.e.
// being pushed down when it was previously up, or being toggled on when it was
// previously off, etc)
//
// If Key == Invalid then the key may not be known, but it can still be
// uniquely identified and it's state watched via the Raw member (e.g. for
// special or non-US keys).
//
// The Raw member must uniquely identify the keyboard button whose state is
// changing, and must always be present regardless of whether or not Key ==
// Invalid.
type StateEvent struct {
	T     time.Time
	Key   Key
	State State
	Raw   uint64
}

// Time returns the time at which this event occured.
func (e StateEvent) Time() time.Time {
	return e.T
}

// String returns an string representation of this event.
func (e StateEvent) String() string {
	return fmt.Sprintf("StateEvent(Key=%v, State=%v, Raw=%v, Time=%v)", e.Key, e.State, e.Raw, e.T)
}

// TypedEvent represents an event where some sort of user input has generated
// an input character which should be considered input.
type TypedEvent struct {
	T    time.Time
	Rune rune
}

// Time returns the time at which this event occured.
func (e TypedEvent) Time() time.Time {
	return e.T
}

// String returns an string representation of this event.
func (e TypedEvent) String() string {
	return fmt.Sprintf("TypedEvent(Rune=%U %q, Time=%v)", e.Rune, string(e.Rune), e.T)
}
