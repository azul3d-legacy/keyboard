// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package keyboard implements various keyboard related data types.
package keyboard

import (
	"fmt"
)

// Key represents an single keyboard button.
//
// It should be noted that it does not represent an character that pressing an
// keyboard button would otherwise generate (hence you will find no capital
// keys defined).
type Key int

const (
	Invalid Key = iota

	// http://en.wikipedia.org/wiki/File:KB_United_States-NoAltGr.svg
	Tilde        // "~"
	Dash         // "-"
	Equals       // "="
	Semicolon    // ";"
	Apostrophe   // "'"
	Comma        // ","
	Period       // "."
	ForwardSlash // "/"
	BackSlash    // "\"
	Backspace
	Tab // "\t"
	CapsLock
	Space // " "
	Enter // "\r", "\n", "\r\n"
	Escape
	Insert
	PrintScreen
	Delete
	PageUp
	PageDown
	Home
	End
	Pause
	Sleep
	Clear
	Select
	Print
	Execute
	Help
	Applications
	ScrollLock
	Play
	Zoom

	// Arrow keys
	ArrowLeft
	ArrowRight
	ArrowDown
	ArrowUp

	// Lefties
	LeftBracket // [
	LeftShift
	LeftCtrl
	LeftSuper
	LeftAlt

	// Righties
	RightBracket // ]
	RightShift
	RightCtrl
	RightSuper
	RightAlt

	// Numbers
	Zero  // "0"
	One   // "1"
	Two   // "2"
	Three // "3"
	Four  // "4"
	Five  // "5"
	Six   // "6"
	Seven // "7"
	Eight // "8"
	Nine  // "9"

	// Functions
	F1
	F2
	F3
	F4
	F5
	F6
	F7
	F8
	F9
	F10
	F11
	F12
	F13
	F14
	F15
	F16
	F17
	F18
	F19
	F20
	F21
	F22
	F23
	F24

	// English characters
	A // "a"
	B // "b"
	C // "c"
	D // "d"
	E // "e"
	F // "f"
	G // "g"
	H // "h"
	I // "i"
	J // "j"
	K // "k"
	L // "l"
	M // "m"
	N // "n"
	O // "o"
	P // "p"
	Q // "q"
	R // "r"
	S // "s"
	T // "t"
	U // "u"
	V // "v"
	W // "w"
	X // "x"
	Y // "y"
	Z // "z"

	// Number pads
	NumLock
	NumMultiply // "*"
	NumDivide   // "/"
	NumAdd      // "+"
	NumSubtract // "-"
	NumZero     // "0"
	NumOne      // "1"
	NumTwo      // "2"
	NumThree    // "3"
	NumFour     // "4"
	NumFive     // "5"
	NumSix      // "6"
	NumSeven    // "7"
	NumEight    // "8"
	NumNine     // "9"
	NumDecimal  // "."
	NumComma    // ","
	NumEnter

	BrowserBack
	BrowserForward
	BrowserRefresh
	BrowserStop
	BrowserSearch
	BrowserFavorites
	BrowserHome

	MediaNext
	MediaPrevious
	MediaStop
	MediaPlayPause

	LaunchMail
	LaunchMedia
	LaunchAppOne
	LaunchAppTwo

	Kana
	Kanji
	Junja
	Attn
	CrSel
	ExSel
	EraseEOF
)

func (k Key) String() string {
	switch k {
	case Invalid:
		return "Invalid"
	case Tilde:
		return "Tilde"
	case Dash:
		return "Dash"
	case Equals:
		return "Equals"
	case Semicolon:
		return "Semicolon"
	case Apostrophe:
		return "Apostrophe"
	case Comma:
		return "Comma"
	case Period:
		return "Period"
	case ForwardSlash:
		return "ForwardSlash"
	case BackSlash:
		return "BackSlash"
	case Backspace:
		return "Backspace"
	case Tab:
		return "Tab"
	case CapsLock:
		return "CapsLock"
	case Space:
		return "Space"
	case Enter:
		return "Enter"
	case Escape:
		return "Escape"
	case Insert:
		return "Insert"
	case PrintScreen:
		return "PrintScreen"
	case Delete:
		return "Delete"
	case PageUp:
		return "PageUp"
	case PageDown:
		return "PageDown"
	case Home:
		return "Home"
	case End:
		return "End"
	case Pause:
		return "Pause"
	case Sleep:
		return "Sleep"
	case Clear:
		return "Clear"
	case Select:
		return "Select"
	case Print:
		return "Print"
	case Execute:
		return "Execute"
	case Help:
		return "Help"
	case Applications:
		return "Applications"
	case ScrollLock:
		return "ScrollLock"
	case Play:
		return "Play"
	case Zoom:
		return "Zoom"
	case ArrowLeft:
		return "ArrowLeft"
	case ArrowRight:
		return "ArrowRight"
	case ArrowDown:
		return "ArrowDown"
	case ArrowUp:
		return "ArrowUp"
	case LeftBracket:
		return "LeftBracket"
	case LeftShift:
		return "LeftShift"
	case LeftCtrl:
		return "LeftCtrl"
	case LeftSuper:
		return "LeftSuper"
	case LeftAlt:
		return "LeftAlt"
	case RightBracket:
		return "RightBracket"
	case RightShift:
		return "RightShift"
	case RightCtrl:
		return "RightCtrl"
	case RightSuper:
		return "RightSuper"
	case RightAlt:
		return "RightAlt"
	case Zero:
		return "Zero"
	case One:
		return "One"
	case Two:
		return "Two"
	case Three:
		return "Three"
	case Four:
		return "Four"
	case Five:
		return "Five"
	case Six:
		return "Six"
	case Seven:
		return "Seven"
	case Eight:
		return "Eight"
	case Nine:
		return "Nine"
	case F1:
		return "F1"
	case F2:
		return "F2"
	case F3:
		return "F3"
	case F4:
		return "F4"
	case F5:
		return "F5"
	case F6:
		return "F6"
	case F7:
		return "F7"
	case F8:
		return "F8"
	case F9:
		return "F9"
	case F10:
		return "F10"
	case F11:
		return "F11"
	case F12:
		return "F12"
	case F13:
		return "F13"
	case F14:
		return "F14"
	case F15:
		return "F15"
	case F16:
		return "F16"
	case F17:
		return "F17"
	case F18:
		return "F18"
	case F19:
		return "F19"
	case F20:
		return "F20"
	case F21:
		return "F21"
	case F22:
		return "F22"
	case F23:
		return "F23"
	case F24:
		return "F24"
	case A:
		return "A"
	case B:
		return "B"
	case C:
		return "C"
	case D:
		return "D"
	case E:
		return "E"
	case F:
		return "F"
	case G:
		return "G"
	case H:
		return "H"
	case I:
		return "I"
	case J:
		return "J"
	case K:
		return "K"
	case L:
		return "L"
	case M:
		return "M"
	case N:
		return "N"
	case O:
		return "O"
	case P:
		return "P"
	case Q:
		return "Q"
	case R:
		return "R"
	case S:
		return "S"
	case T:
		return "T"
	case U:
		return "U"
	case V:
		return "V"
	case W:
		return "W"
	case X:
		return "X"
	case Y:
		return "Y"
	case Z:
		return "Z"
	case NumLock:
		return "NumLock"
	case NumMultiply:
		return "NumMultiply"
	case NumDivide:
		return "NumDivide"
	case NumAdd:
		return "NumAdd"
	case NumSubtract:
		return "NumSubtract"
	case NumZero:
		return "NumZero"
	case NumOne:
		return "NumOne"
	case NumTwo:
		return "NumTwo"
	case NumThree:
		return "NumThree"
	case NumFour:
		return "NumFour"
	case NumFive:
		return "NumFive"
	case NumSix:
		return "NumSix"
	case NumSeven:
		return "NumSeven"
	case NumEight:
		return "NumEight"
	case NumNine:
		return "NumNine"
	case NumDecimal:
		return "NumDecimal"
	case NumEnter:
		return "NumEnter"
	case BrowserBack:
		return "BrowserBack"
	case BrowserForward:
		return "BrowserForward"
	case BrowserRefresh:
		return "BrowserRefresh"
	case BrowserStop:
		return "BrowserStop"
	case BrowserSearch:
		return "BrowserSearch"
	case BrowserFavorites:
		return "BrowserFavorites"
	case BrowserHome:
		return "BrowserHome"
	case MediaNext:
		return "MediaNext"
	case MediaPrevious:
		return "MediaPrevious"
	case MediaStop:
		return "MediaStop"
	case MediaPlayPause:
		return "MediaPlayPause"
	case LaunchMail:
		return "LaunchMail"
	case LaunchMedia:
		return "LaunchMedia"
	case LaunchAppOne:
		return "LaunchAppOne"
	case LaunchAppTwo:
		return "LaunchAppTwo"
	case Kana:
		return "Kana"
	case Junja:
		return "Junja"
	case Kanji:
		return "Kanji"
	case Attn:
		return "Attn"
	case CrSel:
		return "CrSel"
	case ExSel:
		return "ExSel"
	case EraseEOF:
		return "EraseEOF"
	}
	return fmt.Sprintf("Key(%d)", k)
}
