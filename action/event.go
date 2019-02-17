package action

import (
	"log"
	"sync"

	"../g13"
	uinput "github.com/sashko/go-uinput"
)

// Handler Keyboard implementation of EventHandler
type Handler struct {
	Actions Profiles
	Profile string
	event   chan *g13.State
	wg      sync.WaitGroup
	User    UserHandler
}

// EVENTS list of special g13 event types
var EVENTS = map[string]g13.Button{
	"EVENT": g13.Button(0x10000000000),
	"STICK": g13.Button(0x20000000000),
}

// Event process events from g13
func (h *Handler) Event(state g13.State) {
	h.event <- &state
}

// NewHandler initiate a new Handler struct with a bound keyboard
func NewHandler() *Handler {
	kb, err := uinput.CreateKeyboard()
	if err != nil {
		log.Fatalf("CreateKeyboard: %s", err)
	}
	mouse, err := uinput.CreateMice(0, 2000, 0, 2000)
	if err != nil {
		log.Fatalf("CreateMouse: %s", err)
	}
	touchPad, err := uinput.CreateTouchPad(0, 2000, 0, 2000)
	if err != nil {
		log.Fatalf("CreateTouchPad: %s", err)
	}
	h := Handler{Profile: "main", Actions: Profiles{}, event: make(chan *g13.State)}
	h.User = UserHandler{
		Keyboard: &UserKeyboard{kb: kb},
		Mouse:    &UserMouse{mouse: mouse, touchPad: touchPad, stopMoveCh: make(chan bool), movedCh: make(chan bool)},
	}

	h.wg.Add(1)
	go h.eventHandler()

	return &h
}

func (h *Handler) eventHandler() {
	defer h.wg.Done()

	for {
		select {
		case state, ok := <-h.event:
			if !ok {
				log.Println("event handler shutting down")
				return
			}

			stick := (state.X != state.Old.X || state.Y != state.Old.Y)

			for buttons, action := range h.Actions[h.Profile] {
				if (stick && buttons == EVENTS["STICK"]) || buttons == EVENTS["EVENT"] {
					action.Change(state)
				} else if active := state.Buttons.Test(buttons); active != action.Active {
					action.Active = active
					if action.Active {
						action.Down(state)
					} else {
						action.Up(state)
					}
				}
			}
		}
	}
}

// Close the keyboard when done
func (h *Handler) Close() {
	close(h.event)
	h.User.Keyboard.kb.Close()
	h.User.Mouse.mouse.Close()
	h.User.Mouse.touchPad.Close()
	h.wg.Wait()
}
