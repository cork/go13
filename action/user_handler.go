package action

// UserHandler contains handler actions exposed to users
type UserHandler struct {
	h        *Handler
	Keyboard *UserKeyboard
	Mouse    *UserMouse
}

// ChangeProfile Switch profile in handler
func (uh *UserHandler) ChangeProfile(newProfile string) {
	if _, ok := uh.h.Actions[newProfile]; ok {
		uh.h.Profile = newProfile
	}
}

// ChangeConfig Switch configuration file used in handler
func (uh *UserHandler) ChangeConfig(newConfig string) {
	err := uh.h.Config.LoadTOMLConfig(newConfig)
	if err != nil {
		return
	}

	uh.h.Config.ToActions(uh.h, uh.h.Actions)
}
