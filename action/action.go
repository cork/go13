package action

import (
	"log"
	"time"

	"../g13"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

// Actions contains the profile actions
type Actions map[g13.Button]*Action

// Profiles contains all the known actions grouped by profiles
type Profiles map[string]Actions

// Func used for Up and Down calls for actions
type Func func(state *g13.State)

// Action contacts the action and state of the specific action
type Action struct {
	Active bool
	Change Func
	Down   Func
	Up     Func
}

// BindLua provides helpers and binds Up and Down in a lua script
func (h *Handler) BindLua(script *string) *Action {
	L := lua.NewState()
	defer L.Close()

	L.SetGlobal("G13", L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"parseKey": func(L *lua.LState) int {
			key := L.ToString(2)

			L.Push(lua.LNumber(g13.ParseKey(key)))
			return 1
		},
		"test": func(L *lua.LState) int {
			buttons := g13.Button(L.ToInt64(2))
			target := g13.Button(L.ToInt64(3))

			L.Push(lua.LBool(buttons.Test(target)))
			return 1
		},
	}))

	L.SetGlobal("keyboard", L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"parseKey": func(L *lua.LState) int {
			key := L.ToString(2)
			result := ParseKeyboardKey(key)

			L.Push(lua.LNumber(result[0]))
			return 1
		},
		"parseKeys": func(L *lua.LState) int {
			key := L.ToString(2)
			result := ParseKeyboardKey(key)

			L.Push(luar.New(L, result))
			return 1
		},
	}))

	L.SetGlobal("Nanosecond", lua.LNumber(int64(time.Nanosecond)))
	L.SetGlobal("Millisecond", lua.LNumber(int64(time.Millisecond)))
	L.SetGlobal("Second", lua.LNumber(int64(time.Second)))
	L.SetGlobal("Minute", lua.LNumber(int64(time.Minute)))
	L.SetGlobal("Hour", lua.LNumber(int64(time.Hour)))
	L.SetGlobal("sleep", L.NewFunction(func(L *lua.LState) int {
		duration := L.ToInt64(1)
		time.Sleep(time.Duration(duration))

		return 0
	}))

	err := L.DoString(*script)
	if err != nil {
		log.Println(err)
	}

	luaUserHandler := luar.New(L, h.User)

	action := Action{
		Change: func(state *g13.State) {
			luaState := luar.New(L, state)

			L.CallByParam(lua.P{
				Fn:      L.GetGlobal("Change"),
				NRet:    0,
				Protect: false,
			}, luaState, luaUserHandler)
		},
		Down: func(state *g13.State) {
			luaState := luar.New(L, state)

			L.CallByParam(lua.P{
				Fn:      L.GetGlobal("Down"),
				NRet:    0,
				Protect: false,
			}, luaState, luaUserHandler)
		},
		Up: func(state *g13.State) {
			luaState := luar.New(L, state)

			L.CallByParam(lua.P{
				Fn:      L.GetGlobal("Up"),
				NRet:    0,
				Protect: false,
			}, luaState, luaUserHandler)
		},
	}

	if L.GetGlobal("Change").Type() == lua.LTNil {
		action.Change = func(state *g13.State) {}
	}

	if L.GetGlobal("Down").Type() == lua.LTNil {
		action.Down = func(state *g13.State) {}
	}

	if L.GetGlobal("Up").Type() == lua.LTNil {
		action.Up = func(state *g13.State) {}
	}

	return &action
}

// KeyDown Generate a down function for use with Action
func (h *Handler) KeyDown(key string) Func {
	keys := ParseKeyboardKey(key)

	return func(state *g13.State) {
		h.User.Keyboard.KeyDown(keys)
	}
}

// KeyUp Generate a up function for use with Action
func (h *Handler) KeyUp(key string) Func {
	keys := ParseKeyboardKey(key)

	return func(state *g13.State) {
		h.User.Keyboard.KeyUp(keys)
	}
}

// ReverseKeyUp Generate a up function for use with Action when the key is looped in reverse
func (h *Handler) ReverseKeyUp(key string) Func {
	keys := ParseKeyboardKey(key)

	return func(state *g13.State) {
		h.User.Keyboard.ReverseKeyUp(keys)
	}
}

// KeyPressAction is a shortcut for KeyDown + KeyUp with the same key
func (h *Handler) KeyPressAction(key string) *Action {
	return &Action{
		Down: h.KeyDown(key),
		Up:   h.ReverseKeyUp(key),
	}
}
