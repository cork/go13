package action

import (
	"log"
	"strings"

	uinput "github.com/sashko/go-uinput"
)

// KEY possible keybiard keys
var KEY = map[string]uint16{
	"ESC":        uinput.Key1,
	"1":          uinput.Key1,
	"2":          uinput.Key2,
	"3":          uinput.Key3,
	"4":          uinput.Key4,
	"5":          uinput.Key5,
	"6":          uinput.Key6,
	"7":          uinput.Key7,
	"8":          uinput.Key8,
	"9":          uinput.Key9,
	"0":          uinput.Key0,
	"MINUS":      uinput.KeyMinus,
	"EQUAL":      uinput.KeyEqual,
	"BACKSPACE":  uinput.KeyBackspace,
	"TAB":        uinput.KeyTab,
	"Q":          uinput.KeyQ,
	"W":          uinput.KeyW,
	"E":          uinput.KeyE,
	"R":          uinput.KeyR,
	"T":          uinput.KeyT,
	"Y":          uinput.KeyY,
	"U":          uinput.KeyU,
	"I":          uinput.KeyI,
	"O":          uinput.KeyO,
	"P":          uinput.KeyP,
	"LEFTBRACE":  uinput.KeyLeftBrace,
	"RIGHTBRACE": uinput.KeyRightBrace,
	"ENTER":      uinput.KeyEnter,
	"LEFTCTRL":   uinput.KeyLeftCtrl,
	"A":          uinput.KeyA,
	"S":          uinput.KeyS,
	"D":          uinput.KeyD,
	"F":          uinput.KeyF,
	"G":          uinput.KeyG,
	"H":          uinput.KeyH,
	"J":          uinput.KeyJ,
	"K":          uinput.KeyK,
	"L":          uinput.KeyL,
	"SEMICOLON":  uinput.KeySemicolon,
	"APOSTROPHE": uinput.KeyApostrophe,
	"GRAVE":      uinput.KeyGrave,
	"LEFTSHIFT":  uinput.KeyLeftShift,
	"BACKSLASH":  uinput.KeyBackslash,
	"Z":          uinput.KeyZ,
	"X":          uinput.KeyX,
	"C":          uinput.KeyC,
	"V":          uinput.KeyV,
	"B":          uinput.KeyB,
	"N":          uinput.KeyN,
	"M":          uinput.KeyM,
	"COMMA":      uinput.KeyComma,
	"DOT":        uinput.KeyDot,
	"SLASH":      uinput.KeySlash,
	"RIGHTSHIFT": uinput.KeyRightShift,
	"KPASTERISK": uinput.KeyKpAsterisk,
	"LEFTALT":    uinput.KeyLeftAlt,
	"SPACE":      uinput.KeySpace,
	"CAPSLOCK":   uinput.KeyCapsLock,
	"F1":         uinput.KeyF1,
	"F2":         uinput.KeyF2,
	"F3":         uinput.KeyF3,
	"F4":         uinput.KeyF4,
	"F5":         uinput.KeyF5,
	"F6":         uinput.KeyF6,
	"F7":         uinput.KeyF7,
	"F8":         uinput.KeyF8,
	"F9":         uinput.KeyF9,
	"F10":        uinput.KeyF10,
	"NUMLOCK":    uinput.KeyNumLock,
	"SCROLLLOCK": uinput.KeyScrollLock,
	"KP7":        uinput.KeyKp7,
	"KP8":        uinput.KeyKp8,
	"KP9":        uinput.KeyKp9,
	"KPMINUS":    uinput.KeyKpMinus,
	"KP4":        uinput.KeyKp4,
	"KP5":        uinput.KeyKp5,
	"KP6":        uinput.KeyKp6,
	"KPPLUS":     uinput.KeyKpPlus,
	"KP1":        uinput.KeyKp1,
	"KP2":        uinput.KeyKp2,
	"KP3":        uinput.KeyKp3,
	"KP0":        uinput.KeyKp0,
	"KPDOT":      uinput.KeyKpDot,

	"ZENKAKUHANKAKU":   uinput.KeyZenkakuHankaku,
	"102ND":            uinput.Key102Nd,
	"F11":              uinput.KeyF11,
	"F12":              uinput.KeyF12,
	"RO":               uinput.KeyRo,
	"KATAKANA":         uinput.KeyKatakana,
	"HIRAGANA":         uinput.KeyHiragana,
	"HENKAN":           uinput.KeyHenkan,
	"KATAKANAHIRAGANA": uinput.KeyKatakanaHiragana,
	"MUHENKAN":         uinput.KeyMuhenkan,
	"KPJPCOMMA":        uinput.KeyKpJpComma,
	"KPENTER":          uinput.KeyKpEnter,
	"RIGHTCTRL":        uinput.KeyRightCtrl,
	"KPSLASH":          uinput.KeyKpSlash,
	"SYSRQ":            uinput.KeySysRq,
	"RIGHTALT":         uinput.KeyRightAlt,
	"LINEFEED":         uinput.KeyLineFeed,
	"HOME":             uinput.KeyHome,
	"UP":               uinput.KeyUp,
	"PAGEUP":           uinput.KeyPageUp,
	"LEFT":             uinput.KeyLeft,
	"RIGHT":            uinput.KeyRight,
	"END":              uinput.KeyEnd,
	"DOWN":             uinput.KeyDown,
	"PAGEDOWN":         uinput.KeyPageDown,
	"INSERT":           uinput.KeyInsert,
	"DELETE":           uinput.KeyDelete,
	"MACRO":            uinput.KeyMacro,
	"MUTE":             uinput.KeyMute,
	"VOLUMEDOWN":       uinput.KeyVolumeDown,
	"VOLUMEUP":         uinput.KeyVolumeUp,
	"POWER":            uinput.KeyPower, /* SC System Power Down */
	"KPEQUAL":          uinput.KeyKpEqual,
	"KPPLUSMINUS":      uinput.KeyKpPlusMinus,
	"PAUSE":            uinput.KeyPause,
	"SCALE":            uinput.KeyScale, /* AL Compiz Scale (Expose) */

	"KPCOMMA":   uinput.KeyKpComma,
	"HANGEUL":   uinput.KeyHangeul,
	"HANJA":     uinput.KeyHanja,
	"YEN":       uinput.KeyYen,
	"LEFTMETA":  uinput.KeyLeftMeta,
	"RIGHTMETA": uinput.KeyRightMeta,
	"COMPOSE":   uinput.KeyCompose,

	"STOP":           uinput.KeyStop, /* AC Stop */
	"AGAIN":          uinput.KeyAgain,
	"PROPS":          uinput.KeyProps, /* AC Properties */
	"UNDO":           uinput.KeyUndo,  /* AC Undo */
	"FRONT":          uinput.KeyFront,
	"COPY":           uinput.KeyCopy,  /* AC Copy */
	"OPEN":           uinput.KeyOpen,  /* AC Open */
	"PASTE":          uinput.KeyPaste, /* AC Paste */
	"FIND":           uinput.KeyFind,  /* AC Search */
	"CUT":            uinput.KeyCut,   /* AC Cut */
	"HELP":           uinput.KeyHelp,  /* AL Integrated Help Center */
	"MENU":           uinput.KeyMenu,  /* Menu (show menu) */
	"CALC":           uinput.KeyCalc,  /* AL Calculator */
	"SETUP":          uinput.KeySetup,
	"SLEEP":          uinput.KeySleep,  /* SC System Sleep */
	"WAKEUP":         uinput.KeyWakeUp, /* System Wake Up */
	"FILE":           uinput.KeyFile,   /* AL Local Machine Browser */
	"SENDFILE":       uinput.KeySendFile,
	"DELETEFILE":     uinput.KeyDeleteFile,
	"XFER":           uinput.KeyXfer,
	"PROG1":          uinput.KeyProg1,
	"PROG2":          uinput.KeyProg2,
	"WWW":            uinput.KeyWww, /* AL Internet Browser */
	"MSDOS":          uinput.KeyMsDos,
	"SCREENLOCK":     uinput.KeyScreenLock,
	"ROTATE_DISPLAY": uinput.KeyRotateDisplay, /* Display orientation for e.g. tablets */
	"CYCLEWINDOWS":   uinput.KeyCycleWindows,
	"MAIL":           uinput.KeyMail,
	"BOOKMARKS":      uinput.KeyBookmarks, /* AC Bookmarks */
	"COMPUTER":       uinput.KeyComputer,
	"BACK":           uinput.KeyBack,    /* AC Back */
	"FORWARD":        uinput.KeyForward, /* AC Forward */
	"CLOSECD":        uinput.KeyCloseCD,
	"EJECTCD":        uinput.KeyEjectCD,
	"EJECTCLOSECD":   uinput.KeyEjectCloseCd,
	"NEXTSONG":       uinput.KeyNextSong,
	"PLAYPAUSE":      uinput.KeyPlayPause,
	"PREVIOUSSONG":   uinput.KeyPreviousSong,
	"STOPCD":         uinput.KeyStopCD,
	"RECORD":         uinput.KeyRecord,
	"REWIND":         uinput.KeyRewind,
	"PHONE":          uinput.KeyPhone, /* Media Select Telephone */
	"ISO":            uinput.KeyISO,
	"CONFIG":         uinput.KeyConfig,   /* AL Consumer Control Configuration */
	"HOMEPAGE":       uinput.KeyHomePage, /* AC Home */
	"REFRESH":        uinput.KeyRefresh,  /* AC Refresh */
	"EXIT":           uinput.KeyExit,     /* AC Exit */
	"MOVE":           uinput.KeyMove,
	"EDIT":           uinput.KeyEdit,
	"SCROLLUP":       uinput.KeyScrollUp,
	"SCROLLDOWN":     uinput.KeyScrollDown,
	"KPLEFTPAREN":    uinput.KeyKpLeftParen,
	"KPRIGHTPAREN":   uinput.KeyKpRightParen,
	"NEW":            uinput.KeyNew,  /* AC New */
	"REDO":           uinput.KeyRedo, /* AC Redo/Repeat */

	"F13": uinput.KeyF13,
	"F14": uinput.KeyF14,
	"F15": uinput.KeyF15,
	"F16": uinput.KeyF16,
	"F17": uinput.KeyF17,
	"F18": uinput.KeyF18,
	"F19": uinput.KeyF19,
	"F20": uinput.KeyF20,
	"F21": uinput.KeyF21,
	"F22": uinput.KeyF22,
	"F23": uinput.KeyF23,
	"F24": uinput.KeyF24,

	"PLAYCD":         uinput.KeyPlayCd,
	"PAUSECD":        uinput.KeyPauseCd,
	"PROG3":          uinput.KeyProg3,
	"PROG4":          uinput.KeyProg4,
	"DASHBOARD":      uinput.KeyDashBoard, /* AL Dashboard */
	"SUSPEND":        uinput.KeySuspend,
	"CLOSE":          uinput.KeyClose, /* AC Close */
	"PLAY":           uinput.KeyPlay,
	"FASTFORWARD":    uinput.KeyFastForward,
	"BASSBOOST":      uinput.KeyBassBoost,
	"PRINT":          uinput.KeyPrint, /* AC Print */
	"HP":             uinput.KeyHp,
	"CAMERA":         uinput.KeyCamera,
	"SOUND":          uinput.KeySound,
	"QUESTION":       uinput.KeyQuestion,
	"EMAIL":          uinput.KeyEmail,
	"CHAT":           uinput.KeyChat,
	"SEARCH":         uinput.KeySearch,
	"CONNECT":        uinput.KeyConnect,
	"FINANCE":        uinput.KeyFinance, /* AL Checkbook/Finance */
	"SPORT":          uinput.KeySport,
	"SHOP":           uinput.KeyShop,
	"ALTERASE":       uinput.KeyAlterase,
	"CANCEL":         uinput.KeyCancel, /* AC Cancel */
	"BRIGHTNESSDOWN": uinput.KeyBrightnessDown,
	"BRIGHTNESSUP":   uinput.KeyBrightnessUp,
	"MEDIA":          uinput.KeyMedia,

	"SWITCHVIDEOMODE": uinput.KeySwitchVideoMode, /* Cycle between available video
	   outputs (Monitor/LCD/TV-out/etc) */
	"KBDILLUMTOGGLE": uinput.KeyKbdIllumToggle,
	"KBDILLUMDOWN":   uinput.KeyKbdIllumDown,
	"KBDILLUMUP":     uinput.KeyKbdIllumUp,

	"SEND":        uinput.KeySend,        /* AC Send */
	"REPLY":       uinput.KeyReply,       /* AC Reply */
	"FORWARDMAIL": uinput.KeyForwardMail, /* AC Forward Msg */
	"SAVE":        uinput.KeySave,        /* AC Save */
	"DOCUMENTS":   uinput.KeyDocuments,

	"BATTERY": uinput.KeyBattery,

	"BLUETOOTH": uinput.KeyBluetooth,
	"WLAN":      uinput.KeyWLan,
	"UWB":       uinput.KeyUwb,

	"VIDEO_NEXT":       uinput.KeyVideoNext,       /* drive next video source */
	"VIDEO_PREV":       uinput.KeyVideoPrev,       /* drive previous video source */
	"BRIGHTNESS_CYCLE": uinput.KeyBrightnessCycle, /* brightness up, after max is min */
	"BRIGHTNESS_AUTO":  uinput.KeyBrightnessAuto,  /* Set Auto Brightness: manual
	brightness control is off,
	rely on ambient */
	"DISPLAY_OFF": uinput.KeyDisplayOff, /* display device to off state */

	"WWAN":   uinput.KeyWwan,   /* Wireless WAN (LTE, UMTS, GSM, etc.) */
	"RFKILL": uinput.KeyRfKill, /* Key that controls all radios */

	"MICMUTE": uinput.KeyMicMute, /* Mute / unmute the microphone */
}

// ParseKeyboardKey converts a KEY+KEY+... to []int
func ParseKeyboardKey(key string) []uint16 {
	keys := strings.Split(key, "+")

	var b []uint16

	for _, key := range keys {
		if _, ok := KEY[key]; ok {
			b = append(b, KEY[key])
		} else {
			log.Fatalf("Unknown key: %s", key)
		}
	}

	return b
}
