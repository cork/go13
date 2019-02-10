package g13

import (
	"encoding/binary"
	"log"
	"sort"
	"strings"
	"sync"

	"github.com/google/gousb"
)

// EventHandler interface for Event triggering
type EventHandler interface {
	Event(State)
	Close()
}

// Logitech usb Vendor constant
const Logitech = 0x046d

// LogitechG13 usb Product constant
const LogitechG13 = 0xc21c

// Button reprecents the button state of the bitmap
type Button uint64

// State contains the events current keyboard state
type State struct {
	X       byte
	Y       byte
	Buttons Button
}

// FindDevices finds all G13 devices and call ReadDevice for each device found
func FindDevices(action EventHandler) bool {
	ctx := gousb.NewContext()

	devs, err := ctx.OpenDevices(func(desc *gousb.DeviceDesc) bool {
		if desc.Vendor == Logitech && desc.Product == LogitechG13 {
			log.Printf("%03d.%03d %s:%s\n", desc.Bus, desc.Address, desc.Vendor, desc.Product)
			log.Printf("%+v\n", desc)
			return true
		}

		return false
	})

	if err != nil {
		log.Fatalf("list: %s", err)
	}

	var wg sync.WaitGroup

	for _, dev := range devs {
		wg.Add(1)
		go ReadDevice(dev, action, &wg)
	}

	if len(devs) == 0 {
		log.Fatal("No matching devices found.")
		ctx.Close()
		return false
	}

	go func() {
		wg.Wait()
		ctx.Close()
	}()

	return true
}

// ReadDevice reading putput from device
func ReadDevice(dev *gousb.Device, action EventHandler, wg *sync.WaitGroup) {
	defer dev.Close()
	defer wg.Done()

	dev.SetAutoDetach(true)

	buttons := make([]byte, 8)

	intf, done, err := dev.DefaultInterface()
	if err != nil {
		log.Fatalf("%s.DefaultInterface(): %v", dev, err)
	}
	defer done()

	endpoint, err := intf.InEndpoint(1)
	if err != nil {
		log.Fatalf("dev.InEndpoint(): %s", err)
	}

	buf := make([]byte, endpoint.Desc.MaxPacketSize)
	for {
		readBytes, err := endpoint.Read(buf)
		if err != nil {
			log.Fatalf("Reading from device failed: %v", err)
		}
		copy(buttons, buf[3:readBytes])
		state := State{
			X:       buf[1],
			Y:       buf[2],
			Buttons: Button(binary.LittleEndian.Uint64(buttons)),
		}
		log.Printf("%+v\n", state)
		action.Event(state)
	}
}

func (b Button) String() string {
	buttons := []string{}
	for name, value := range KEY {
		if b&value == value {
			buttons = append(buttons, name)
		}
	}

	sort.Sort(keys(buttons))

	return strings.Join(buttons, "+")
}

// Test tests if a button combo is pressent
func (b Button) Test(button Button) bool {
	return b&button == button
}
