package action

import (
	"time"

	uinput "github.com/sashko/go-uinput"
)

// UserMouse wrapper around uinput mouse
type UserMouse struct {
	mouse      uinput.Mice
	touchPad   uinput.TouchPad
	moving     bool
	stopMoveCh chan bool
	movedCh    chan bool
}

// MoveX wrapper around uinput Mouse MoveX
func (um *UserMouse) MoveX(x int32) {
	um.mouse.MoveX(x)
}

// MoveY wrapper around uinput Mouse MoveY
func (um *UserMouse) MoveY(y int32) {
	um.mouse.MoveY(y)
}

// MoveXYByStick moves by x,y minus centerX/Y if the abs value is bigger then deadZone repeated every step duration
func (um *UserMouse) MoveXYByStick(x, y, centerX, centerY, xdeadZone, ydeadZone int32, speed float32) {
	if um.moving {
		um.stopMoveCh <- true
		<-um.movedCh
		um.moving = false
	}

	x = x - centerX
	y = y - centerY

	if x < xdeadZone && x > -xdeadZone {
		x = 0
	} else {
		if x > xdeadZone {
			x -= xdeadZone
		} else {
			x += xdeadZone
		}
	}

	if y < ydeadZone && y > -ydeadZone {
		y = 0
	} else {
		if y > ydeadZone {
			y -= ydeadZone
		} else {
			y += ydeadZone
		}
	}

	x = int32(float32(x) * speed)
	y = int32(float32(y) * speed)

	if x != 0 || y != 0 {
		go func() {
			first := true
			for {
				select {
				case <-um.stopMoveCh:
					um.movedCh <- true
					return
				default:
					if x != 0 {
						um.MoveX(x)
					}
					if y != 0 {
						um.MoveY(y)
					}
				}
				if first {
					um.movedCh <- true
					first = false
				}
				time.Sleep(10 * time.Millisecond)
			}
		}()
		<-um.movedCh
		um.moving = true
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
