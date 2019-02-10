package action

import (
	uinput "github.com/sashko/go-uinput"
)

// UserMouse wrapper around uinput mouse
type UserMouse struct {
	mouse    uinput.Mice
	touchPad uinput.TouchPad
}

// MoveX wrapper around uinput Mouse MoveX
func (um *UserMouse) MoveX(x int32) {
	um.mouse.MoveX(x)
}

// MoveY wrapper around uinput Mouse MoveY
func (um *UserMouse) MoveY(y int32) {
	um.mouse.MoveY(y)
}

// MoveXYByStick moves by x,y minus centerX/Y if the abs value is bigger then deadZone
func (um *UserMouse) MoveXYByStick(x, y, centerX, centerY, deadZone int32) {
	x = x - centerX
	y = y - centerY
	if x > deadZone || x < -deadZone {
		um.MoveX(x)
	}

	if y > deadZone || y < -deadZone {
		um.MoveY(y)
	}
}

// MoveTo wrapper around uinput TouchPad MoveTo
func (um *UserMouse) MoveTo(x, y int32) {
	um.touchPad.MoveTo(x, y)
}

// LeftPress wrapper around uinput Mouse LeftPress
func (um *UserMouse) LeftPress() {
	um.mouse.LeftPress()
}

// LeftRelease wrapper around uinput Mouse LeftRelease
func (um *UserMouse) LeftRelease() {
	um.mouse.LeftRelease()
}

// LeftClick wrapper around uinput Mouse LeftClick
func (um *UserMouse) LeftClick() {
	um.mouse.LeftClick()
}

// RightPress wrapper around uinput Mouse RightPress
func (um *UserMouse) RightPress() {
	um.mouse.RightPress()
}

// RightRelease wrapper around uinput Mouse RightRelease
func (um *UserMouse) RightRelease() {
	um.mouse.RightRelease()
}

// RightClick wrapper around uinput Mouse RightClick
func (um *UserMouse) RightClick() {
	um.mouse.RightClick()
}

// MiddleClick wrapper around uinput Mouse MiddleClick
func (um *UserMouse) MiddleClick() {
	um.mouse.MiddleClick()
}

// SideClick wrapper around uinput Mouse SideClick
func (um *UserMouse) SideClick() {
	um.mouse.SideClick()
}

// ExtraClick wrapper around uinput Mouse ExtraClick
func (um *UserMouse) ExtraClick() {
	um.mouse.ExtraClick()
}

// ForwardClick wrapper around uinput Mouse ForwardClick
func (um *UserMouse) ForwardClick() {
	um.mouse.ForwardClick()
}

// BackClick wrapper around uinput Mouse BackClick
func (um *UserMouse) BackClick() {
	um.mouse.BackClick()
}
