package g13

import (
	"log"
	"strings"
)

// Known buttons
const (
	G1         = Button(0x1)
	G2         = Button(0x2)
	G3         = Button(0x4)
	G4         = Button(0x8)
	G5         = Button(0x10)
	G6         = Button(0x20)
	G7         = Button(0x40)
	G8         = Button(0x80)
	G9         = Button(0x100)
	G10        = Button(0x200)
	G11        = Button(0x400)
	G12        = Button(0x800)
	G13        = Button(0x1000)
	G14        = Button(0x2000)
	G15        = Button(0x4000)
	G16        = Button(0x8000)
	G17        = Button(0x10000)
	G18        = Button(0x20000)
	G19        = Button(0x40000)
	G20        = Button(0x80000)
	G21        = Button(0x100000)
	G22        = Button(0x200000)
	Undef1     = Button(0x400000)
	LightState = Button(0x800000)
	BD         = Button(0x1000000)
	L1         = Button(0x2000000)
	L2         = Button(0x4000000)
	L3         = Button(0x8000000)
	L4         = Button(0x10000000)
	M1         = Button(0x20000000)
	M2         = Button(0x40000000)
	M3         = Button(0x80000000)
	MR         = Button(0x100000000)
	Left       = Button(0x200000000)
	Down       = Button(0x400000000)
	Top        = Button(0x800000000)
	Undef3     = Button(0x1000000000)
	Light      = Button(0x2000000000)
	Light2     = Button(0x4000000000)
	MiscToggle = Button(0x8000000000)
	Stick      = Button(0x10000000000) // Fake key added for event handling of stick movements
)

// KEY map of all the known G13 buttons
var KEY = map[string]Button{
	"BD":          BD,
	"DOWN":        Down,
	"G1":          G1,
	"G2":          G2,
	"G3":          G3,
	"G4":          G4,
	"G5":          G5,
	"G6":          G6,
	"G7":          G7,
	"G8":          G8,
	"G9":          G9,
	"G10":         G10,
	"G11":         G11,
	"G12":         G12,
	"G13":         G13,
	"G14":         G14,
	"G15":         G15,
	"G16":         G16,
	"G17":         G17,
	"G18":         G18,
	"G19":         G19,
	"G20":         G20,
	"G21":         G21,
	"G22":         G22,
	"L1":          L1,
	"L2":          L2,
	"L3":          L3,
	"L4":          L4,
	"LEFT":        Left,
	"LIGHT_STATE": LightState,
	"LIGHT":       Light,
	"LIGHT2":      Light2,
	"M1":          M1,
	"M2":          M2,
	"M3":          M3,
	"MR":          MR,
	"TOP":         Top,
	"UNDEF1":      Undef1,
	"UNDEF3":      Undef3,
	"STICK":       Stick,
}

type keys []string

func (a keys) Len() int {
	return len(a)
}
func (a keys) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a keys) Less(i, j int) bool {
	return a[i][0] < a[j][0] || (a[i][0] == a[j][0] && len(a[i]) < len(a[j])) || (a[i] < a[j] && len(a[i]) == len(a[j]))
}

// ParseKey generates a Button value from a KEY+KEY+... String
func ParseKey(key string) Button {
	keys := strings.Split(key, "+")

	var b Button

	for _, key := range keys {
		if _, ok := KEY[key]; ok {
			b |= KEY[key]
		} else {
			log.Fatalf("Unknown key: %s", key)
		}
	}

	return b
}
