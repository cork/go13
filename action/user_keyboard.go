package action

import (
	uinput "github.com/sashko/go-uinput"
)

// UserKeyboard wrapper around uinput keyboard
type UserKeyboard struct {
	kb uinput.Keyboard
}

// KeyDown wrapper around uinput Keyboard KeyDown
func (uk *UserKeyboard) KeyDown(keys []uint16) {
	for _, key := range keys {
		uk.kb.KeyDown(key)
	}
}

// KeyUp wrapper around uinput Keyboard KeyUp
func (uk *UserKeyboard) KeyUp(keys []uint16) {
	for _, key := range keys {
		uk.kb.KeyUp(key)
	}
}

// ReverseKeyUp wrapper around uinput Keyboard KeyUp
func (uk *UserKeyboard) ReverseKeyUp(keys []uint16) {
	for i := len(keys) - 1; i >= 0; i-- {
		uk.kb.KeyUp(keys[i])
	}
}
